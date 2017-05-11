/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.5.dev2
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goesiv1

/* 200 ok object */
type GetUniverseStargatesStargateIdOk struct {
	Destination GetUniverseStargatesStargateIdDestination `json:"destination,omitempty"`
	Name        string                                    `json:"name,omitempty"` /* name string */
	Position    GetUniverseStargatesStargateIdPosition    `json:"position,omitempty"`
	StargateId  int32                                     `json:"stargate_id,omitempty"` /* stargate_id integer */
	SystemId    int32                                     `json:"system_id,omitempty"`   /* The solar system this stargate is in */
	TypeId      int32                                     `json:"type_id,omitempty"`     /* type_id integer */
}
