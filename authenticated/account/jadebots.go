package authenticated

import (
	Connection "gw2sdk/connection"
)

type JadeBots struct {
	Gw2sdk Connection.GW2sdk
}

type JadeBotsResponse []int

func (a *JadeBots) Get() JadeBotsResponse {

	response := JadeBotsResponse{}
	a.Gw2sdk.Retrieve("account/jadebots", nil, &response)

	return response

}
