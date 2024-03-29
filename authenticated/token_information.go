package authenticated

import (
	Connection "gw2sdk/connection"
)

type TokenInformation struct {
	Gw2sdk Connection.GW2sdk
}

type TokenInformationResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func (a *TokenInformation) Get() TokenInformationResponse {

	response := TokenInformationResponse{}
	a.Gw2sdk.Retrieve("tokeninfo", nil, &response)

	return response

}
