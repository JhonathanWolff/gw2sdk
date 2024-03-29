package authenticated

import (
	Connection "gw2sdk/connection"
)

type MapChests struct {
	Gw2sdk Connection.GW2sdk
}

type MapChestsResponse []string

func (a *MapChests) Get() MapChestsResponse {

	response := MapChestsResponse{}
	a.Gw2sdk.Retrieve("account/mapchests", nil, &response)

	return response

}
