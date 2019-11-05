/*******************************************************************************
 * Copyright 2019 Dell Inc.
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

package main

import (
	"flag"

	"github.com/objectbox/edgex-objectbox"
	"github.com/objectbox/edgex-objectbox/internal"
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
	"github.com/objectbox/edgex-objectbox/internal/support/scheduler"

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

	httpServer := httpserver.NewBootstrap(scheduler.LoadRestRoutes())
	bootstrap.Run(
		configDir,
		profileDir,
		internal.ConfigFileName,
		useRegistry,
		clients.SupportSchedulerServiceKey,
		scheduler.Configuration,
		startupTimer,
		di.NewContainer(di.ServiceConstructorMap{}),
		[]interfaces.BootstrapHandler{
			secret.NewSecret().BootstrapHandler,
			database.NewDatabase(&httpServer, scheduler.Configuration).BootstrapHandler,
			scheduler.BootstrapHandler,
			telemetry.BootstrapHandler,
			httpServer.BootstrapHandler,
			message.NewBootstrap(clients.SupportSchedulerServiceKey, edgex.Version).BootstrapHandler,
		})
}
