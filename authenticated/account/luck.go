package authenticated

import (
	Connection "gw2sdk/connection"
)

type Luck struct {
	Gw2sdk Connection.GW2sdk
}

type LuckResponse []struct {
	Id    string `json:"id"`
	Value int    `json:"value"`
}

func (a *Luck) Get() LuckResponse {

	response := LuckResponse{}
	a.Gw2sdk.Retrieve("account/luck", nil, &response)

	return response

}
