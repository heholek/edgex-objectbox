//
// Copyright (c) 2018 Tencent
// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package client

import (
	"context"
	"sync"

	"github.com/objectbox/edgex-objectbox/internal/export"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/container"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/startup"
	"github.com/objectbox/edgex-objectbox/internal/pkg/di"
	"github.com/objectbox/edgex-objectbox/internal/pkg/endpoint"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/export/distro"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

// Global variables
var dbClient export.DBClient
var LoggingClient logger.LoggingClient
var Configuration = &ConfigurationStruct{}
var dc distro.DistroClient

// BootstrapHandler fulfills the BootstrapHandler contract and performs initialization needed by the export-client service.
func BootstrapHandler(wg *sync.WaitGroup, ctx context.Context, startupTimer startup.Timer, dic *di.Container) bool {
	// update global variables.
	LoggingClient = container.LoggingClientFrom(dic.Get)
	dbClient = container.DBClientFrom(dic.Get)

	// initialize clients required by service.
	registryClient := container.RegistryFrom(dic.Get)
	dc = distro.NewDistroClient(
		types.EndpointParams{
			ServiceKey:  clients.ExportDistroServiceKey,
			UseRegistry: registryClient != nil,
			Url:         Configuration.Clients["Distro"].Url(),
			Interval:    Configuration.Service.ClientMonitor,
		},
		endpoint.Endpoint{RegistryClient: &registryClient})

	return true
}
