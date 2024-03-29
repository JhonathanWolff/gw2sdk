package authenticated

import (
	Connection "gw2sdk/connection"
)

type Progression struct {
	Gw2sdk Connection.GW2sdk
}

type ProgressionResponse []struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

func (a *Progression) Get() ProgressionResponse {

	response := ProgressionResponse{}
	a.Gw2sdk.Retrieve("account/jadebots", nil, &response)

	return response

}
