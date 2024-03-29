package authenticated

import (
	Connection "gw2sdk/connection"
)

type Achivements struct {
	Gw2sdk Connection.GW2sdk
}

type AchivementsResponse []struct {
	ID      int   `json:"id"`
	Current int   `json:"current,omitempty"`
	Max     int   `json:"max,omitempty"`
	Done    bool  `json:"done"`
	Bits    []int `json:"bits,omitempty"`
}

func (a *Achivements) Get() AchivementsResponse {

	response := AchivementsResponse{}
	a.Gw2sdk.Retrieve("account/achievements", nil, &response)

	return response

}
