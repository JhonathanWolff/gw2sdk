package authenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type GuildStash struct {
	Gw2sdk Connection.GW2sdk
}

type GuildStashResponseDetails []struct {
	UpgradeID int    `json:"upgrade_id"`
	Size      int    `json:"size"`
	Coins     int    `json:"coins"`
	Note      string `json:"note,omitempty"`
	Inventory []struct {
		ID    int `json:"id"`
		Count int `json:"count"`
	} `json:"inventory"`
}

func (a *GuildStash) GetDetails(guild_id string) GuildStashResponseDetails {

	path := fmt.Sprintf("guild/%s/stash", guild_id)

	response := GuildStashResponseDetails{}
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
