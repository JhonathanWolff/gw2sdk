package authenticated

import (
	Connection "gw2sdk/connection"
)

type PvPStanding struct {
	Gw2sdk Connection.GW2sdk
}

type PvPStandingResponse []struct {
	Current struct {
		TotalPoints int `json:"total_points"`
		Division    int `json:"division"`
		Tier        int `json:"tier"`
		Points      int `json:"points"`
		Repeats     int `json:"repeats"`
		Rating      int `json:"rating"`
		Decay       int `json:"decay"`
	} `json:"current"`
	Best struct {
		TotalPoints int `json:"total_points"`
		Division    int `json:"division"`
		Tier        int `json:"tier"`
		Points      int `json:"points"`
		Repeats     int `json:"repeats"`
	} `json:"best"`
	SeasonID string `json:"season_id"`
}

func (a *PvPStanding) Get() PvPStandingResponse {

	response := PvPStandingResponse{}
	a.Gw2sdk.Retrieve("pvp/standings", nil, &response)

	return response

}
