package authenticated

import (
	Connection "gw2sdk/connection"
)

type Bank struct {
	Gw2sdk Connection.GW2sdk
}

type BankResponse []struct {
	ID        int    `json:"id"`
	Count     int    `json:"count"`
	Upgrades  []int  `json:"upgrades,omitempty"`
	Infusions []int  `json:"infusions,omitempty"`
	Skin      int    `json:"skin,omitempty"`
	Dyes      []int  `json:"dyes,omitempty"`
	Binding   string `json:"binding,omitempty"`
}

func (a *Bank) Get() BankResponse {

	response := BankResponse{}
	a.Gw2sdk.Retrieve("account/bank", nil, &response)

	return response

}
