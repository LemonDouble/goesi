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
type GetCharactersCharacterIdChatChannels200Ok struct {

	/* allowed array */
	Allowed []GetCharactersCharacterIdChatChannelsAllowed `json:"allowed,omitempty"`

	/* blocked array */
	Blocked []GetCharactersCharacterIdChatChannelsBlocked `json:"blocked,omitempty"`

	/* Unique channel ID. Always negative for player-created channels. Permanent (CCP created) channels have a positive ID, but don't appear in the API */
	ChannelId int32 `json:"channel_id,omitempty"`

	/* Normalized, unique string used to compare channel names */
	ComparisonKey string `json:"comparison_key,omitempty"`

	/* Whether this is a password protected channel */
	HasPassword bool `json:"has_password,omitempty"`

	/* Message of the day for this channel */
	Motd string `json:"motd,omitempty"`

	/* muted array */
	Muted []GetCharactersCharacterIdChatChannelsMuted `json:"muted,omitempty"`

	/* Displayed name of channel */
	Name string `json:"name,omitempty"`

	/* operators array */
	Operators []GetCharactersCharacterIdChatChannelsOperator `json:"operators,omitempty"`

	/* owner_id integer */
	OwnerId int32 `json:"owner_id,omitempty"`
}
