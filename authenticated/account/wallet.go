package authenticated

import (
	Connection "gw2sdk/connection"
)

type Wallet struct {
	Gw2sdk Connection.GW2sdk
}

type WalletResponse []struct {
	Id    int `json:"id"`
	Value int `json:"value"`
}

func (a *Wallet) Get() WalletResponse {

	response := WalletResponse{}
	a.Gw2sdk.Retrieve("account/wallet", nil, &response)

	return response

}
