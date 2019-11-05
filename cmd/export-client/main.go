//
// Copyright (c) 2017
// Mainflux
// Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"flag"

	"github.com/objectbox/edgex-objectbox"
	"github.com/objectbox/edgex-objectbox/internal"
	"github.com/objectbox/edgex-objectbox/internal/export/client"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/handlers/database"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/handlers/httpserver"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/handlers/message"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/handlers/secret"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/interfaces"
	"github.com/objectbox/edgex-objectbox/internal/pkg/bootstrap/startup"
	"github.com/objectbox/edgex-objectbox/internal/pkg/di"
	"github.com/objectbox/edgex-objectbox/internal/pkg/telemetry"
	"github.com/objectbox/edgex-objectbox/internal/pkg/usage"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
)

func main() {
	startupTimer := startup.NewStartUpTimer(internal.BootRetrySecondsDefault, internal.BootTimeoutSecondsDefault)

	var useRegistry bool
	var configDir, profileDir string

	flag.BoolVar(&useRegistry, "registry", false, "Indicates the service should use Registry.")
	flag.BoolVar(&useRegistry, "r", false, "Indicates the service should use Registry.")
	flag.StringVar(&profileDir, "profile", "", "Specify a profile other than default.")
	flag.StringVar(&profileDir, "p", "", "Specify a profile other than default.")
	flag.StringVar(&configDir, "confdir", "", "Specify local configuration directory")

	flag.Usage = usage.HelpCallback
	flag.Parse()

	httpServer := httpserver.NewBootstrap(client.LoadRestRoutes())
	bootstrap.Run(
		configDir,
		profileDir,
		internal.ConfigFileName,
		useRegistry,
		clients.ExportClientServiceKey,
		client.Configuration,
		startupTimer,
		di.NewContainer(di.ServiceConstructorMap{}),
		[]interfaces.BootstrapHandler{
			secret.NewSecret().BootstrapHandler,
			database.NewDatabase(&httpServer, client.Configuration).BootstrapHandler,
			client.BootstrapHandler,
			telemetry.BootstrapHandler,
			httpServer.BootstrapHandler,
			message.NewBootstrap(clients.ExportClientServiceKey, edgex.Version).BootstrapHandler,
		})
}
