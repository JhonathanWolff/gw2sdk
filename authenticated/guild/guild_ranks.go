package authenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type GuildRanks struct {
	Gw2sdk Connection.GW2sdk
}

type GuildRanksResponseDetails []struct {
	ID          string   `json:"id"`
	Order       int      `json:"order"`
	Permissions []string `json:"permissions"`
	Icon        string   `json:"icon"`
}

func (a *GuildRanks) GetDetails(guild_id string) GuildRanksResponseDetails {

	path := fmt.Sprintf("guild/%s/ranks", guild_id)

	response := GuildRanksResponseDetails{}
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
