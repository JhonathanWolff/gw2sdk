package authenticated

import (
	Connection "gw2sdk/connection"
)

type Minis struct {
	Gw2sdk Connection.GW2sdk
}

type MinisResponse []int

func (a *Minis) Get() MinisResponse {

	response := MinisResponse{}
	a.Gw2sdk.Retrieve("account/minis", nil, &response)

	return response

}
