package unauthenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type Dailycrafting struct {
	Gw2sdk Connection.GW2sdk
}

type DailycraftingResponse []string

type DailycraftingResponseDetails []struct {
	Id string `json:"id"`
}

func (a *Dailycrafting) Get() DailycraftingResponse {

	response := DailycraftingResponse{}
	a.Gw2sdk.Retrieve("dailycrafting", nil, &response)

	return response

}

func (a *Dailycrafting) GetDetails(id string) DailycraftingResponseDetails {

	response := DailycraftingResponseDetails{}

	path := fmt.Sprintf("dailycrafting/%s", id)
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
