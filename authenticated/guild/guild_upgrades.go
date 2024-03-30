package authenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type GuildUpgrades struct {
	Gw2sdk Connection.GW2sdk
}

type GuildUpgradesResponseDetails []int

func (a *GuildUpgrades) GetDetails(guild_id string) GuildUpgradesResponseDetails {

	path := fmt.Sprintf("guild/%s/upgrades", guild_id)

	response := GuildUpgradesResponseDetails{}
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
