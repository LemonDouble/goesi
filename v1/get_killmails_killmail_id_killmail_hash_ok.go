/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.3.dev1
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
type GetKillmailsKillmailIdKillmailHashOk struct {
	Attackers     []GetKillmailsKillmailIdKillmailHashAttacker `json:"attackers,omitempty"`       /* attackers array */
	KillmailId    int32                                        `json:"killmail_id,omitempty"`     /* ID of the killmail */
	KillmailTime  time.Time                                    `json:"killmail_time,omitempty"`   /* Time that the victim was killed and the killmail generated  */
	MoonId        int32                                        `json:"moon_id,omitempty"`         /* Moon if the kill took place at one */
	SolarSystemId int32                                        `json:"solar_system_id,omitempty"` /* Solar system that the kill took place in  */
	Victim        GetKillmailsKillmailIdKillmailHashVictim     `json:"victim,omitempty"`
	WarId         int32                                        `json:"war_id,omitempty"` /* War if the killmail is generated in relation to an official war  */
}
