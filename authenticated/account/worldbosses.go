package authenticated

import (
	Connection "gw2sdk/connection"
)

type WorldBosses struct {
	Gw2sdk Connection.GW2sdk
}

type WorldBossesResponse []string

func (a *WorldBosses) Get() WorldBossesResponse {

	response := WorldBossesResponse{}
	a.Gw2sdk.Retrieve("account/worldbosses", nil, &response)

	return response

}
