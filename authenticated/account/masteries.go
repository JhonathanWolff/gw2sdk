package authenticated

import (
	Connection "gw2sdk/connection"
)

type Masteries struct {
	Gw2sdk Connection.GW2sdk
}

type MasteriesResponse []struct {
	ID    int `json:"id"`
	Level int `json:"level"`
}

func (a *Masteries) Get() MasteriesResponse {

	response := MasteriesResponse{}
	a.Gw2sdk.Retrieve("account/masteries", nil, &response)

	return response

}
