package authenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type GuildTreasury struct {
	Gw2sdk Connection.GW2sdk
}

type GuildTreasuryResponseDetails []struct {
	ItemID   int `json:"item_id"`
	Count    int `json:"count"`
	NeededBy []struct {
		UpgradeID int `json:"upgrade_id"`
		Count     int `json:"count"`
	} `json:"needed_by"`
}

func (a *GuildTreasury) GetDetails(guild_id string) GuildTreasuryResponseDetails {

	path := fmt.Sprintf("guild/%s/treasury", guild_id)

	response := GuildTreasuryResponseDetails{}
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
