package unauthenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type Jadebots struct {
	Gw2sdk Connection.GW2sdk
}

type JadebotsResponse []int

type JadebotsResponseDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UnlockItem  int    `json:"unlock_item"`
}

func (a *Jadebots) Get() JadebotsResponse {

	response := JadebotsResponse{}
	a.Gw2sdk.Retrieve("jadebots", nil, &response)

	return response

}

func (a *Jadebots) GetDetails(id int) JadebotsResponseDetails {

	response := JadebotsResponseDetails{}
	path := fmt.Sprintf("jadebots/%d", id)
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
