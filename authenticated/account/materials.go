package authenticated

import (
	Connection "gw2sdk/connection"
)

type Materials struct {
	Gw2sdk Connection.GW2sdk
}

type MaterialsResponse []struct {
	ID       int `json:"id"`
	Category int `json:"category"`
	Count    int `json:"count"`
}

func (a *Materials) Get() MaterialsResponse {

	response := MaterialsResponse{}
	a.Gw2sdk.Retrieve("account/materials", nil, &response)

	return response

}
