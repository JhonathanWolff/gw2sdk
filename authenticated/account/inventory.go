package authenticated

import (
	Connection "gw2sdk/connection"
)

type Inventory struct {
	Gw2sdk Connection.GW2sdk
}

type InventoryResponse []struct {
	ID      int    `json:"id"`
	Count   int    `json:"count"`
	Binding string `json:"binding"`
}

func (a *Inventory) Get() InventoryResponse {

	response := InventoryResponse{}
	a.Gw2sdk.Retrieve("account/inventory", nil, &response)

	return response

}
