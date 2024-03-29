package authenticated

import (
	Connection "gw2sdk/connection"
)

type Finishers struct {
	Gw2sdk Connection.GW2sdk
}

type FinishersResponse []struct {
	Id        int  `json:"id"`
	Permanent bool `json:"permanent"`
}

func (a *Finishers) Get() FinishersResponse {

	response := FinishersResponse{}
	a.Gw2sdk.Retrieve("account/finishers", nil, &response)

	return response

}
