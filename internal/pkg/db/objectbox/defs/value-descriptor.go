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

type ValueDescriptor struct {
	Id           string
	Created      int64
	Description  string
	Modified     int64
	Origin       int64
	Name         string      `objectbox:"unique"`
	Min          interface{} `objectbox:"type:[]byte converter:interfaceJson"`
	Max          interface{} `objectbox:"type:[]byte converter:interfaceJson"`
	DefaultValue interface{} `objectbox:"type:[]byte converter:interfaceJson"`
	Type         string
	UomLabel     string
	Formatting   string
	Labels       []string
}
