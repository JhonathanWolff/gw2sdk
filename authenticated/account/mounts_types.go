package authenticated

import (
	Connection "gw2sdk/connection"
)

type MountsTypes struct {
	Gw2sdk Connection.GW2sdk
}

type MountsTypesResponse []string

func (a *MountsTypes) Get() MountsTypesResponse {

	response := MountsTypesResponse{}
	a.Gw2sdk.Retrieve("account/mounts/types", nil, &response)

	return response

}
