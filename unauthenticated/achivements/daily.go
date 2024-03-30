package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type AchivementDaily struct {
	Gw2sdk Connection.GW2sdk
}

type DailyDetails struct {
	ID    int `json:"id"`
	Level struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"level"`
	RequiredAccessObject struct {
		Product   string `json:"product"`
		Condition string `json:"condition"`
	} `json:"required_access"`
	//RequiredAccessArray []string `json:"required_access"`
}

type AchivementDailyResponse struct{}

// https://wiki.guildwars2.com/wiki/API:2/achievements/daily
func (a *AchivementDaily) Get(parameters map[string]string) interface{} {

	response := AchivementResponseDetails{}
	a.Gw2sdk.Retrieve("achievements/daily", parameters, &response)
	return response

}
