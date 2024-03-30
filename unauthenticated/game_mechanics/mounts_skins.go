package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type MountSkins struct {
	Gw2sdk Connection.GW2sdk
}

type MountSkinsResponse []string
type MountSkinsResponseDetails []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Mount    string `json:"mount"`
	DyeSlots []struct {
		ColorID  int    `json:"color_id"`
		Material string `json:"material"`
	} `json:"dye_slots"`
}

func (a *MountSkins) Get() MountSkinsResponse {

	response := MountSkinsResponse{}
	a.Gw2sdk.Retrieve("mounts/skins", nil, &response)

	return response

}

func (a *MountSkins) GetDetails(parameters map[string]string) MountSkinsResponseDetails {

	response := MountSkinsResponseDetails{}
	a.Gw2sdk.Retrieve("mounts/skins", parameters, &response)

	return response

}
