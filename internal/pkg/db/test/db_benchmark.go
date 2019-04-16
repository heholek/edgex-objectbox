/*
 * Copyright 2018 ObjectBox Ltd. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strconv"
	"time"

	"github.com/edgexfoundry/edgex-go/internal/core/data/interfaces"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
)

type BenchmarkContext struct {
	db interfaces.DBClient

	start time.Time
	stop  time.Time

	// Iteration, 0 based
	I int
}

// (Re)starts the benchmark clock; do any initialization before this call
func (b *BenchmarkContext) StartClock() {
	b.start = time.Now()
}

// Stops the benchmark clock; do any clean up after this call
func (b *BenchmarkContext) StopClock() {
	b.stop = time.Now()
}

func RunBenchmarkN(db interfaces.DBClient, name string, n int, f func(ctx *BenchmarkContext) error) {
	if n == 0 {
		panic("Zero count")
	}
	var durationLo, durationHi, durationSum time.Duration

	ctx := BenchmarkContext{db: db}

	runtime.GC() // Do GC before to avoid unrelated GC affecting results
	runtime.GC() // Run twice to catch objects with finalizers too
	for i := 0; i < n; i++ {
		ctx.I = i
		ctx.start = time.Now()
		ctx.stop = time.Time{}
		err := f(&ctx)
		duration := time.Since(ctx.start)
		if !ctx.stop.IsZero() {
			duration = ctx.stop.Sub(ctx.start)
		}

		if err != nil {
			panic("Benchmark " + name + " failed after " + duration.String() + " in " + strconv.Itoa(i) +
				". iteration: " + err.Error())
		}
		durationSum += duration
		if duration > durationHi {
			durationHi = duration
		}
		if duration < durationLo || durationLo.Nanoseconds() == 0 {
			durationLo = duration
		}
	}

	nsSum := durationSum.Nanoseconds()
	iterationsPerSec := math.NaN()
	if nsSum != 0 {
		iterationsPerSec = float64(n) * float64(time.Second) / float64(nsSum)
	}
	precision := 2
	if iterationsPerSec < 100 {
		precision = 10
	}
	ips := strconv.FormatFloat(iterationsPerSec, 'f', precision, 64)
	fmt.Printf("%v: %v iterations in %v (%v iterations per second)\n", name, n, durationSum, ips)
	durationAvg := time.Duration(uint64(durationSum) / uint64(n))
	fmt.Printf("%v iterations: avg: %v, lo: %v, hi: %v\n", name, durationAvg, durationLo, durationHi)
	println()

}

func BenchmarkDBFixedN(db interfaces.DBClient, verify bool) {
	defer db.CloseSession()
	durable := true
	benchmarkReadingsN(db, verify, durable)
}

func benchmarkReadingsN(db interfaces.DBClient, verify bool, durable bool) {
	// Plain IDs do not require .hex(); must use reflect to avoid import cycle to identify DB
	dbType := reflect.TypeOf(db).String()
	println("\nBenchmarking " + dbType)
	println("---------------------------------------------")

	// Remove any events and readings before and after test
	_ = db.ScrubAllEvents()
	defer db.ScrubAllEvents()

	count := 10000
	countPostfix := "[" + strconv.Itoa(count) + "]"
	readings := make([]string, count)
	RunBenchmarkN(db, "AddReading", count, func(ctx *BenchmarkContext) error {
		reading := contract.Reading{}
		reading.Name = "test" + strconv.Itoa(ctx.I)
		reading.Device = "device" + strconv.Itoa(ctx.I/100)
		ctx.StartClock()
		id, err := db.AddReading(reading)
		if durable && ctx.I == count-1 {
			// Last one; ensure DBs actually made data durable
			durableStart := time.Now()
			db.EnsureAllDurable(false)
			ctx.StopClock() // Stop asap before logging
			durableDuration := time.Since(durableStart)
			println("Making changes durable: " + durableDuration.String())
		} else {
			ctx.StopClock()
		}
		readings[ctx.I] = id
		return err
	})

	RunBenchmarkN(db, "Readings"+countPostfix, 10, func(ctx *BenchmarkContext) error {
		readings, err := db.Readings()
		ctx.StopClock()
		size := len(readings)
		if verify && size != count {
			panic("Unexpected size: " + strconv.Itoa(size))
		}
		return err
	})

	RunBenchmarkN(db, "ReadingCount"+countPostfix, 100, func(ctx *BenchmarkContext) error {
		size, err := db.ReadingCount()
		ctx.StopClock()
		if verify && size != count {
			panic("Unexpected size: " + strconv.Itoa(size))
		}
		return err
	})

	RunBenchmarkN(db, "ReadingById", count, func(ctx *BenchmarkContext) error {
		id := readings[ctx.I]
		ctx.StartClock()
		reading, err := db.ReadingById(id)
		ctx.StopClock()

		if verify && reading.Id != id {
			println(reading.String())
			panic("Expected ID " + id + " but got " + reading.Id)
		}

		return err
	})

	RunBenchmarkN(db, "ReadingsByDevice", 100, func(ctx *BenchmarkContext) error {
		device := "device" + strconv.Itoa(ctx.I)
		ctx.StartClock()
		slice, err := db.ReadingsByDevice(device, 100)
		ctx.StopClock()

		if verify {
			if len(slice) != 100 {
				panic("Unexpected slice size: " + strconv.Itoa(len(slice)))
			}

			for idx, reading := range slice {
				if reading.Device != device {
					println("[" + strconv.Itoa(idx) + "] " + reading.String())
					panic("Expected device " + device + " but got " + reading.Device)
				}
			}
		}
		return err
	})

}
