package main

import (
	"github.com/objectbox/edgex-objectbox/internal/pkg/config"
	"github.com/objectbox/edgex-objectbox/internal/support/scheduler"
	"testing"
)

func TestToml(t *testing.T) {
	configuration := &scheduler.ConfigurationStruct{}
	if err := config.VerifyTomlFiles(configuration); err != nil {
		t.Fatalf("%v", err)
	}
	if configuration.Service.StartupMsg == "" {
		t.Errorf("configuration.StartupMsg is zero length.")
	}
}
