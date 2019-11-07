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

package defs

import (
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
)

// Command represents a command on a device and is used by core-command service.
// In addition to model.Command, it adds DeviceId - a reference to device

type Command struct {
	models.Timestamps `objectbox:"inline"`
	Id                string
	Name              string
	Get               Get
	Put               Put

	// Additional
	DeviceId uint64 `objectbox:"link:Device"`
}

func (c *Command) ToContract() contract.Command {
	return contract.Command{
		Timestamps: c.Timestamps,
		Id:         c.Id,
		Name:       c.Name,
		Get: contract.Get{
			contract.Action{
				c.Get.Action.Path,
				c.Get.Action.Responses,
				c.Get.Action.URL,
			},
		},
		Put: contract.Put{
			contract.Action{
				c.Put.Action.Path,
				c.Put.Action.Responses,
				c.Put.Action.URL,
			},
			c.Put.ParameterNames,
		},
	}
}

func CommandSliceToContract(commands []Command) []contract.Command {
	var result = make([]contract.Command, 0, len(commands))
	for i := range commands {
		result = append(result, commands[i].ToContract())
	}
	return result
}
