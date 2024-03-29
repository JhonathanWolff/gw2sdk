package authenticated

import (
	Connection "gw2sdk/connection"
)

type Outfits struct {
	Gw2sdk Connection.GW2sdk
}

type OutfitsResponse []int

func (a *Outfits) Get() OutfitsResponse {

	response := OutfitsResponse{}
	a.Gw2sdk.Retrieve("account/outfits", nil, &response)

	return response

}
