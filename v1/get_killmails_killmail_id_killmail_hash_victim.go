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

/* victim object */
type GetKillmailsKillmailIdKillmailHashVictim struct {

	/* alliance_id integer */
	AllianceId int32 `json:"alliance_id,omitempty"`

	/* character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`

	/* corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`

	/* How much total damage was taken by the victim  */
	DamageTaken int32 `json:"damage_taken,omitempty"`

	/* faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`

	/* items array */
	Items []GetKillmailsKillmailIdKillmailHashItem1 `json:"items,omitempty"`

	Position GetKillmailsKillmailIdKillmailHashPosition `json:"position,omitempty"`

	/* The ship that the victim was piloting and was destroyed  */
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
}
