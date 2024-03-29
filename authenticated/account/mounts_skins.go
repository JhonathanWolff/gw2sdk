package authenticated

import (
	Connection "gw2sdk/connection"
)

type MountsSkins struct {
	Gw2sdk Connection.GW2sdk
}

type MountsSkinsResponse []int

func (a *MountsSkins) Get() MountsSkinsResponse {

	response := MountsSkinsResponse{}
	a.Gw2sdk.Retrieve("account/mounts/skins", nil, &response)

	return response

}
