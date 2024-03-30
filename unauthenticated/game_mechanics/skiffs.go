package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Skiffs struct {
	Gw2sdk Connection.GW2sdk
}

type SkiffsResponse []int

type SkiffsResponseDetails []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	DyeSlots []struct {
		ColorID  int    `json:"color_id"`
		Material string `json:"material"`
	} `json:"dye_slots"`
}

func (a *Skiffs) Get() SkiffsResponse {

	response := SkiffsResponse{}
	a.Gw2sdk.Retrieve("skiffs", nil, &response)

	return response

}

func (a *Skiffs) GetDetails(parameters map[string]string) SkiffsResponseDetails {

	response := SkiffsResponseDetails{}
	a.Gw2sdk.Retrieve("skiffs", parameters, &response)

	return response

}
