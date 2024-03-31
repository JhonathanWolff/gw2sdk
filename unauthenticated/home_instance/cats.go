package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type HomeCats struct {
	Gw2sdk Connection.GW2sdk
}

type HomeCatsResponse []string

type HomeCatsResponseDetails []struct {
	Id   int    `json:"id"`
	Hint string `json:"hint"`
}

func (a *HomeCats) GetDetails(parameters map[string]string) HomeCatsResponseDetails {

	response := HomeCatsResponseDetails{}

	a.Gw2sdk.Retrieve("home/cats", parameters, &response)

	return response

}
