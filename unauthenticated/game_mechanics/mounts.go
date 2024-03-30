package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Mounts struct {
	Gw2sdk Connection.GW2sdk
}

type MountsResponse []string

func (a *Mounts) Get() MountsResponse {

	response := MountsResponse{}
	a.Gw2sdk.Retrieve("mounts", nil, &response)

	return response

}
