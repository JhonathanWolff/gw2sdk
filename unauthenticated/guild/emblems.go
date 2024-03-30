package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Emblem struct {
	Gw2sdk Connection.GW2sdk
}

type EmblemResponse []string

func (a *Emblem) Get() EmblemResponse {

	response := EmblemResponse{}
	a.Gw2sdk.Retrieve("emblem", nil, &response)

	return response

}
