/*******************************************************************************
 * Copyright 2017 Dell Inc.
 * Copyright (c) 2019 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package command

import (
	"context"
	"sync"

	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/container"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/startup"
	"github.com/objectbox/edgex-objectbox/internal/pkg/di"
	"github.com/objectbox/edgex-objectbox/internal/pkg/endpoint"
	"github.com/objectbox/edgex-objectbox/internal/pkg/errorconcept"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/metadata"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

// Global variables
var Configuration = &ConfigurationStruct{}
var LoggingClient logger.LoggingClient

var mdc metadata.DeviceClient
// OBX - forwarding calls to metadata as it's database is the one currently storing all commands data
// var dbClient interfaces.DBClient
var mdcc metadata.CommandClient

// Global ErrorConcept variables
var httpErrorHandler errorconcept.ErrorHandler

// BootstrapHandler fulfills the BootstrapHandler contract and performs initialization needed by the command service.
func BootstrapHandler(wg *sync.WaitGroup, ctx context.Context, startupTimer startup.Timer, dic *di.Container) bool {
	// update global variables.
	LoggingClient = container.LoggingClientFrom(dic.Get)
	httpErrorHandler = errorconcept.NewErrorHandler(LoggingClient)
	// OBX
	// dbClient = container.DBClientFrom(dic.Get)

	// initialize clients required by service.
	registryClient := container.RegistryFrom(dic.Get)
	mdc = metadata.NewDeviceClient(
		types.EndpointParams{
			ServiceKey:  clients.CoreMetaDataServiceKey,
			Path:        clients.ApiDeviceRoute,
			UseRegistry: registryClient != nil,
			Url:         Configuration.Clients["Metadata"].Url() + clients.ApiDeviceRoute,
			Interval:    Configuration.Service.ClientMonitor,
		},
		endpoint.Endpoint{RegistryClient: &registryClient})

	// OBX
	mdcc = metadata.NewCommandClient(
		types.EndpointParams{
			ServiceKey:  clients.CoreMetaDataServiceKey,
			Path:        clients.ApiCommandRoute,
			UseRegistry: registryClient != nil,
			Url:         Configuration.Clients["Metadata"].Url() + clients.ApiCommandRoute,
			Interval:    Configuration.Service.ClientMonitor,
		},
		endpoint.Endpoint{RegistryClient: &registryClient})

	return true
}
