//
// Copyright (c) 2018
// Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"testing"

	"github.com/objectbox/edgex-objectbox/internal/pkg/config"
	"github.com/objectbox/edgex-objectbox/internal/support/notifications"
)

func TestToml(t *testing.T) {
	configuration := &notifications.ConfigurationStruct{}
	if err := config.VerifyTomlFiles(configuration); err != nil {
		t.Fatalf("%v", err)
	}
	if configuration.Service.StartupMsg == "" {
		t.Errorf("configuration.StartupMsg is zero length.")
	}
}
