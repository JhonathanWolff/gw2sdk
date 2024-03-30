package unauthenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type Mapchests struct {
	Gw2sdk Connection.GW2sdk
}

type MapchestsResponse []string

type MapchestsResponseDetails []struct {
	Id string `json:"id"`
}

func (a *Mapchests) Get() MapchestsResponse {

	response := MapchestsResponse{}
	a.Gw2sdk.Retrieve("mapchests", nil, &response)

	return response

}

func (a *Mapchests) GetDetails(id string) MapchestsResponseDetails {

	response := MapchestsResponseDetails{}

	path := fmt.Sprintf("mapchests/%s", id)
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
