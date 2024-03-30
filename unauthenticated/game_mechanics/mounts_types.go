package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type MountsTypes struct {
	Gw2sdk Connection.GW2sdk
}

type MountsTypesResponse []string
type MountsTypesResponseDetails []struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DefaultSkin int    `json:"default_skin"`
	Skins       []int  `json:"skins"`
	Skills      []struct {
		ID   int    `json:"id"`
		Slot string `json:"slot"`
	} `json:"skills"`
}

func (a *MountsTypes) Get() MountsTypesResponse {

	response := MountsTypesResponse{}
	a.Gw2sdk.Retrieve("mounts/types", nil, &response)

	return response

}

func (a *MountsTypes) GetDetails(parameters map[string]string) MountsTypesResponseDetails {

	response := MountsTypesResponseDetails{}
	a.Gw2sdk.Retrieve("mounts/types", parameters, &response)

	return response

}
