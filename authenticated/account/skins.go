package authenticated

import (
	Connection "gw2sdk/connection"
)

type Skins struct {
	Gw2sdk Connection.GW2sdk
}

type SkinsResponse []int

func (a *Skins) Get() SkinsResponse {

	response := SkinsResponse{}
	a.Gw2sdk.Retrieve("account/skins", nil, &response)

	return response

}
