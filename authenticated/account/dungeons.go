package authenticated

import (
	Connection "gw2sdk/connection"
)

type Dungeons struct {
	Gw2sdk Connection.GW2sdk
}

type DungeonsResponse []string

func (a *Dungeons) Get() DungeonsResponse {

	response := DungeonsResponse{}
	a.Gw2sdk.Retrieve("account/dungeons", nil, &response)

	return response

}
