package authenticated

import (
	Connection "gw2sdk/connection"
)

type MailCarriers struct {
	Gw2sdk Connection.GW2sdk
}

type MailCarriersResponse []int

func (a *MailCarriers) Get() MailCarriersResponse {

	response := MailCarriersResponse{}
	a.Gw2sdk.Retrieve("account/mailcarriers", nil, &response)

	return response

}
