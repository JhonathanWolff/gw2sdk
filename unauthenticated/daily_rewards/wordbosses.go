package unauthenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type Worldbosses struct {
	Gw2sdk Connection.GW2sdk
}

type WorldbossesResponse []string

type WorldbossesResponseDetails []struct {
	Id string `json:"id"`
}

func (a *Worldbosses) Get() WorldbossesResponse {

	response := WorldbossesResponse{}
	a.Gw2sdk.Retrieve("worldbosses", nil, &response)

	return response

}

func (a *Worldbosses) GetDetails(id string) WorldbossesResponseDetails {

	response := WorldbossesResponseDetails{}

	path := fmt.Sprintf("worldbosses/%s", id)
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
