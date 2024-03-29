package authenticated

import (
	Connection "gw2sdk/connection"
)

type Raids struct {
	Gw2sdk Connection.GW2sdk
}

type RaidsResponse []string

func (a *Raids) Get() RaidsResponse {

	response := RaidsResponse{}
	a.Gw2sdk.Retrieve("account/raids", nil, &response)

	return response

}
