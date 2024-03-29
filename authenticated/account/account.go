package authenticated

import (
	Connection "gw2sdk/connection"
)

type Account struct {
	Gw2sdk Connection.GW2sdk
}

type AccountResponse struct {
	Id   string `json:id`
	Name string `json:name`
}

func (a *Account) Get() AccountResponse {

	response := AccountResponse{}
	a.Gw2sdk.Retrieve("account", nil, &response)

	return response

}
