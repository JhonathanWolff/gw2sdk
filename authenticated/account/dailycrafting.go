package authenticated

import (
	Connection "gw2sdk/connection"
)

type DailyCraftingResponse []string

type DailyCrafting struct {
	Gw2sdk Connection.GW2sdk
}

func (a *DailyCrafting) Get() DailyCraftingResponse {

	response := DailyCraftingResponse{}
	a.Gw2sdk.Retrieve("account/dailycrafting", nil, &response)

	return response

}
