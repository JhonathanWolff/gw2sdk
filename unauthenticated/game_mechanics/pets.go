package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Pets struct {
	Gw2sdk Connection.GW2sdk
}

type PetsResponse []int

type PetsResponseDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Skills      []struct {
		ID int `json:"id"`
	} `json:"skills"`
}

func (a *Pets) Get() PetsResponse {

	response := PetsResponse{}
	a.Gw2sdk.Retrieve("pets", nil, &response)

	return response

}

func (a *Pets) GetDetails(parameters map[string]string) PetsResponseDetails {

	response := PetsResponseDetails{}
	a.Gw2sdk.Retrieve("pets", parameters, &response)

	return response

}
