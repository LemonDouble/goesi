/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.7
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
type GetSovereigntyStructures200Ok struct {
	AllianceId                  int32     `json:"alliance_id,omitempty"`                   /* The alliance that owns the structure.  */
	SolarSystemId               int32     `json:"solar_system_id,omitempty"`               /* Solar system in which the structure is located.  */
	StructureId                 int64     `json:"structure_id,omitempty"`                  /* Unique item ID for this structure. */
	StructureTypeId             int32     `json:"structure_type_id,omitempty"`             /* A reference to the type of structure this is.  */
	VulnerabilityOccupancyLevel float32   `json:"vulnerability_occupancy_level,omitempty"` /* The occupancy level for the next or current vulnerability window. This takes into account all development indexes and capital system bonuses. Also known as Activity Defense Multiplier from in the client. It increases the time that attackers must spend using their entosis links on the structure.  */
	VulnerableEndTime           time.Time `json:"vulnerable_end_time,omitempty"`           /* The time at which the next or current vulnerability window ends. At the end of a vulnerability window the next window is recalculated and locked in along with the vulnerabilityOccupancyLevel. If the structure is not in 100% entosis control of the defender, it will go in to 'overtime' and stay vulnerable for as long as that situation persists. Only once the defenders have 100% entosis control and has the vulnerableEndTime passed does the vulnerability interval expire and a new one is calculated.  */
	VulnerableStartTime         time.Time `json:"vulnerable_start_time,omitempty"`         /* The next time at which the structure will become vulnerable. Or the start time of the current window if current time is between this and vulnerableEndTime.  */
}
