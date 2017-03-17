/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev8
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

import (
	"time"
)

/* 200 ok object */
type GetMarketsStructuresStructureId200Ok struct {

	/* duration integer */
	Duration int32 `json:"duration,omitempty"`

	/* is_buy_order boolean */
	IsBuyOrder bool `json:"is_buy_order,omitempty"`

	/* issued string */
	Issued time.Time `json:"issued,omitempty"`

	/* location_id integer */
	LocationId int64 `json:"location_id,omitempty"`

	/* min_volume integer */
	MinVolume int32 `json:"min_volume,omitempty"`

	/* order_id integer */
	OrderId int64 `json:"order_id,omitempty"`

	/* price number */
	Price float32 `json:"price,omitempty"`

	/* range string */
	Range_ string `json:"range,omitempty"`

	/* type_id integer */
	TypeId int32 `json:"type_id,omitempty"`

	/* volume_remain integer */
	VolumeRemain int32 `json:"volume_remain,omitempty"`

	/* volume_total integer */
	VolumeTotal int32 `json:"volume_total,omitempty"`
}
