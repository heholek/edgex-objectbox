/*******************************************************************************
 * Copyright 2017 Dell Inc.
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

package models

import (
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// this is not an entity, (that's why the file is named .skip.go)

type Service struct {
	models.DescribedObject `objectbox:"inline"`
	Id                     string
	Name                   string `objectbox:"unique"`
	LastConnected          int64
	LastReported           int64
	OperatingState         models.OperatingState
	Labels                 []string
	Addressable            Addressable `objectbox:"link"`
}
