package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Outfits struct {
	Gw2sdk Connection.GW2sdk
}

type OutfitsResponse []int

type OutfitsResponseDetails []struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	UnlockItems []int  `json:"unlock_items"`
}

func (a *Outfits) Get() OutfitsResponse {

	response := OutfitsResponse{}
	a.Gw2sdk.Retrieve("outfits", nil, &response)

	return response

}

func (a *Outfits) GetDetails(parameters map[string]string) OutfitsResponseDetails {

	response := OutfitsResponseDetails{}
	a.Gw2sdk.Retrieve("outfits", parameters, &response)

	return response

}
