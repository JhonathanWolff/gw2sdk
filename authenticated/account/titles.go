package authenticated

import (
	Connection "gw2sdk/connection"
)

type Titles struct {
	Gw2sdk Connection.GW2sdk
}

type TitlesResponse []int

func (a *Titles) Get() TitlesResponse {

	response := TitlesResponse{}
	a.Gw2sdk.Retrieve("account/titles", nil, &response)

	return response

}
