package authenticated

import (
	Connection "gw2sdk/connection"
)

type PvpHeroes struct {
	Gw2sdk Connection.GW2sdk
}

type PvpHeroesResponse []int

func (a *PvpHeroes) Get() PvpHeroesResponse {

	response := PvpHeroesResponse{}
	a.Gw2sdk.Retrieve("account/pvp/heroes", nil, &response)

	return response

}
