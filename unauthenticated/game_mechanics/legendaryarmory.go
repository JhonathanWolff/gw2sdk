package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Legendaryarmory struct {
	Gw2sdk Connection.GW2sdk
}

type LegendaryarmoryResponse []struct {
	Id       int `json:"id"`
	MaxCount int `json:"max_count"`
}

func (a *Legendaryarmory) Get() LegendaryarmoryResponse {

	response := LegendaryarmoryResponse{}
	a.Gw2sdk.Retrieve("legendaryarmory?ids=all", nil, &response)

	return response

}
