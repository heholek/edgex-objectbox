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
	"github.com/edgexfoundry/edgex-go/internal/core/data/interfaces"
	"math"
	"strconv"
	"time"
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
