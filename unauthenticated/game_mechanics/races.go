package unauthenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type Races struct {
	Gw2sdk Connection.GW2sdk
}

type RacesResponse []int

type RacesResponseDetails []struct {
	ID     string `json:"id"`
	Skills []int  `json:"skills"`
}

func (a *Races) Get() RacesResponse {

	response := RacesResponse{}
	a.Gw2sdk.Retrieve("races", nil, &response)

	return response

}

func (a *Races) GetDetails(race string) RacesResponseDetails {

	response := RacesResponseDetails{}
	path := fmt.Sprintf("races/%s", race)
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
