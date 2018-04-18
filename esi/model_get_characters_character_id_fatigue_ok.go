/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.0
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

package esi

import (
	"time"
)

/* A list of GetCharactersCharacterIdFatigueOk. */
//easyjson:json
type GetCharactersCharacterIdFatigueOkList []GetCharactersCharacterIdFatigueOk

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdFatigueOk struct {
	LastJumpDate          time.Time `json:"last_jump_date,omitempty"`           /* Character's last jump activation */
	JumpFatigueExpireDate time.Time `json:"jump_fatigue_expire_date,omitempty"` /* Character's jump fatigue expiry */
	LastUpdateDate        time.Time `json:"last_update_date,omitempty"`         /* Character's last jump update */
}
