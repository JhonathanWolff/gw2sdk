package authenticated

import (
	Connection "gw2sdk/connection"
)

type WizzardsVaultListings struct {
	Gw2sdk Connection.GW2sdk
}

type WizzardsVaultListingsResponse []struct {
	ID            int    `json:"id"`
	ItemID        int    `json:"item_id"`
	ItemCount     int    `json:"item_count"`
	Type          string `json:"type"`
	Cost          int    `json:"cost"`
	Purchased     int    `json:"purchased,omitempty"`
	PurchaseLimit int    `json:"purchase_limit,omitempty"`
}

func (a *WizzardsVaultListings) Get() WizzardsVaultListingsResponse {

	response := WizzardsVaultListingsResponse{}
	a.Gw2sdk.Retrieve("account/wizardsvault/listings", nil, &response)

	return response

}
