package authenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
	"time"
)

type CommerceTransactions struct {
	Gw2sdk Connection.GW2sdk
}

type CommerceTransactionsResponse []struct {
	ID        int64     `json:"id"`
	ItemID    int       `json:"item_id"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	Created   time.Time `json:"created"`
	Purchased time.Time `json:"purchased"`
}

func (a *CommerceTransactions) Get(
	history_or_current string,
	buys_or_sells string,
) CommerceTransactionsResponse {

	response := CommerceTransactionsResponse{}
	path := fmt.Sprintf("commerce/transactions/%s/%s", history_or_current, buys_or_sells)
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
