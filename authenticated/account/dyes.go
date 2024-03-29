package authenticated

import (
	Connection "gw2sdk/connection"
)

type Dyes struct {
	Gw2sdk Connection.GW2sdk
}

type DyesResponse []int

func (a *Dyes) Get() DyesResponse {

	response := DyesResponse{}
	a.Gw2sdk.Retrieve("account/dyes", nil, &response)

	return response

}
