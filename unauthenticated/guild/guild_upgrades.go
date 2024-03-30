package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type GuildUpgrades struct {
	Gw2sdk Connection.GW2sdk
}

type GuildUpgradesResponse []int

type GuildUpgradesResponseDetails []struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	BuildTime     int    `json:"build_time"`
	Icon          string `json:"icon"`
	Type          string `json:"type"`
	RequiredLevel int    `json:"required_level"`
	Experience    int    `json:"experience"`
	Prerequisites []int  `json:"prerequisites"`
	BagMaxItems   int    `json:"bag_max_items"`
	BagMaxCoins   int    `json:"bag_max_coins"`
	Costs         []struct {
		Type   string `json:"type"`
		Count  int    `json:"count"`
		Name   string `json:"name,omitempty"`
		ItemID int    `json:"item_id,omitempty"`
	} `json:"costs"`
}

func (a *GuildUpgrades) Get() GuildUpgradesResponse {

	response := GuildUpgradesResponse{}
	a.Gw2sdk.Retrieve("guild/upgrades", nil, &response)

	return response

}

func (a *GuildUpgrades) GetDetails(parameters map[string]string) GuildUpgradesResponseDetails {

	response := GuildUpgradesResponseDetails{}
	a.Gw2sdk.Retrieve("guild/upgrades", parameters, &response)

	return response

}
