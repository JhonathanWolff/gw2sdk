package authenticated

import (
	Connection "gw2sdk/connection"
)

type Skiffs struct {
	Gw2sdk Connection.GW2sdk
}

type SkiffsResponse []int

func (a *Skiffs) Get() SkiffsResponse {

	response := SkiffsResponse{}
	a.Gw2sdk.Retrieve("account/skiffs", nil, &response)

	return response

}
