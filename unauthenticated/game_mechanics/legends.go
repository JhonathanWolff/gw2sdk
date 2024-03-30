package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Legends struct {
	Gw2sdk Connection.GW2sdk
}

type LegendsResponse []string

type LegendsResponseDetails []struct {
	ID        string `json:"id"`
	Swap      int    `json:"swap"`
	Heal      int    `json:"heal"`
	Elite     int    `json:"elite"`
	Utilities []int  `json:"utilities"`
}

func (a *Legends) Get() LegendsResponse {

	response := LegendsResponse{}
	a.Gw2sdk.Retrieve("legends", nil, &response)

	return response

}

func (a *Legends) GetDetails(parameters map[string]string) LegendsResponseDetails {

	response := LegendsResponseDetails{}
	a.Gw2sdk.Retrieve("legends", parameters, &response)

	return response

}
