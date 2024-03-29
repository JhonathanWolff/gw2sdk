package authenticated

import (
	Connection "gw2sdk/connection"
)

type CreateSubToken struct {
	Gw2sdk Connection.GW2sdk
}

type CreateSubTokenResponse []struct {
	SubToken string `json:"subtoken"`
}

func (a *CreateSubToken) Get(paramters map[string]string) CreateSubTokenResponse {

	response := CreateSubTokenResponse{}
	a.Gw2sdk.Retrieve("createsubtoken", paramters, &response)

	return response

}
