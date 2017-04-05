/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev22
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

	/* name string */
	Name string `json:"name,omitempty"`

	Position GetUniverseStargatesStargateIdPosition `json:"position,omitempty"`

	/* stargate_id integer */
	StargateId int32 `json:"stargate_id,omitempty"`

	/* The solar system this stargate is in */
	SystemId int32 `json:"system_id,omitempty"`

	/* type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
