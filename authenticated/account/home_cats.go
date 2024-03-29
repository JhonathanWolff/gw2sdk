package authenticated

import (
	Connection "gw2sdk/connection"
)

type HomeCats struct {
	Gw2sdk Connection.GW2sdk
}

type HomeCatsResponse []struct {
	ID   int    `json:"id"`
	Hint string `json:"hint"`
}

func (a *HomeCats) Get() HomeCatsResponse {

	response := HomeCatsResponse{}
	a.Gw2sdk.Retrieve("account/home/cats", nil, &response)

	return response

}
