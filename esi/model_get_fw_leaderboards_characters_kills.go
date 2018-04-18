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

/* A list of GetFwLeaderboardsCharactersKills. */
//easyjson:json
type GetFwLeaderboardsCharactersKillsList []GetFwLeaderboardsCharactersKills

/* Top 100 rankings of pilots by number of kills from yesterday, last week and in total. */
//easyjson:json
type GetFwLeaderboardsCharactersKills struct {
	Yesterday   []GetFwLeaderboardsCharactersYesterday   `json:"yesterday,omitempty"`    /* Top 100 ranking of pilots by kills in the past day */
	LastWeek    []GetFwLeaderboardsCharactersLastWeek    `json:"last_week,omitempty"`    /* Top 100 ranking of pilots by kills in the past week */
	ActiveTotal []GetFwLeaderboardsCharactersActiveTotal `json:"active_total,omitempty"` /* Top 100 ranking of pilots active in faction warfare by total kills. A pilot is considered \"active\" if they have participated in faction warfare in the past 14 days. */
}
