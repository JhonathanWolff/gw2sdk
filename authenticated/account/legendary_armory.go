package authenticated

import (
	Connection "gw2sdk/connection"
)

type LegendaryArmory struct {
	Gw2sdk Connection.GW2sdk
}

type LegendaryArmoryResponse []struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}

func (a *LegendaryArmory) Get() LegendaryArmoryResponse {

	response := LegendaryArmoryResponse{}
	a.Gw2sdk.Retrieve("account/legendaryarmory", nil, &response)

	return response

}
