package authenticated

import (
	Connection "gw2sdk/connection"
)

type Novelties struct {
	Gw2sdk Connection.GW2sdk
}

type NoveltiesResponse []int

func (a *Novelties) Get() NoveltiesResponse {

	response := NoveltiesResponse{}
	a.Gw2sdk.Retrieve("account/novelties", nil, &response)

	return response

}
