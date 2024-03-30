package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Achivement struct {
	Gw2sdk Connection.GW2sdk
}

type AchivementResponse []int

type AchivementResponseDetails []struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	LockedText  string   `json:"locked_text"`
	Type        string   `json:"type"`
	Flags       []string `json:"flags"`
	Tiers       []struct {
		Count  int `json:"count"`
		Points int `json:"points"`
	} `json:"tiers"`
	Rewards []struct {
		Type  string `json:"type"`
		ID    int    `json:"id,omitempty"`
		Count int    `json:"count"`
	} `json:"rewards"`
}

func (a *Achivement) GetDtails(parameters map[string]string) interface{} {
	response := AchivementResponseDetails{}
	a.Gw2sdk.Retrieve("achievements", parameters, &response)
	return response

}

func (a *Achivement) Get() interface{} {

	response := AchivementResponse{}
	a.Gw2sdk.Retrieve("achievements", nil, &response)
	return response

}
