package authenticated

import (
	Connection "gw2sdk/connection"
)

type MasteriesPoints struct {
	Gw2sdk Connection.GW2sdk
}

type Totals struct {
	Region string `json:"region"`
	Spent  int    `json:"spent"`
	Earned int    `json:"earned"`
}

type MasteriesPointsResponse struct {
	Totals   []Totals `json:"totals"`
	Unlocked []int    `json:"unlocked"`
}

func (a *MasteriesPoints) Get() MasteriesPointsResponse {

	response := MasteriesPointsResponse{}
	a.Gw2sdk.Retrieve("account/mastery/points", nil, &response)

	return response

}
