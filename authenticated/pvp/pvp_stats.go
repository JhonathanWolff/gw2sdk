package authenticated

import (
	Connection "gw2sdk/connection"
)

type PvPStats struct {
	Gw2sdk Connection.GW2sdk
}

type PvpAggregate struct {
	Wins       int `json:"wins"`
	Losses     int `json:"losses"`
	Desertions int `json:"desertions"`
	Byes       int `json:"byes"`
	Forfeits   int `json:"forfeits"`
}
type PvpProfessionsStats struct {
	Wins       int `json:"wins"`
	Losses     int `json:"losses"`
	Desertions int `json:"desertions"`
	Byes       int `json:"byes"`
	Forfeits   int `json:"forfeits"`
}

type LadderStats struct {
	Wins       int `json:"wins"`
	Losses     int `json:"losses"`
	Desertions int `json:"desertions"`
	Byes       int `json:"byes"`
	Forfeits   int `json:"forfeits"`
}

type PvPStatsResponse struct {
	PvpRank          int          `json:"pvp_rank"`
	PvpRankPoints    int          `json:"pvp_rank_points"`
	PvpRankRollovers int          `json:"pvp_rank_rollovers"`
	Aggregate        PvpAggregate `json:"aggregate"`
	Professions      struct {
		Elementalist PvpProfessionsStats `json:"elementalist"`
		Guardian     PvpProfessionsStats `json:"guardian"`
		Mesmer       PvpProfessionsStats `json:"mesmer"`
		Thief        PvpProfessionsStats `json:"thief"`
		Necromancer  PvpProfessionsStats `json:"necromancer"`
		Ranger       PvpProfessionsStats `json:"ranger"`
		Engineer     PvpProfessionsStats `json:"engineer"`
		Revenant     PvpProfessionsStats `json:"revenant"`
	} `json:"professions"`
	Ladders struct {
		Ranked   LadderStats `json:"ranked"`
		Unranked LadderStats `json:"unranked"`
	} `json:"ladders"`
}

func (a *PvPStats) Get() PvPStatsResponse {

	response := PvPStatsResponse{}
	a.Gw2sdk.Retrieve("pvp/stats", nil, &response)

	return response

}
