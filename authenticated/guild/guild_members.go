package authenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
	"time"
)

type GuildMembers struct {
	Gw2sdk Connection.GW2sdk
}

type GuildMembersResponseDetails []struct {
	Name   string    `json:"name"`
	Rank   string    `json:"rank"`
	Joined time.Time `json:"joined"`
}

func (a *GuildMembers) GetDetails(guild_id string) GuildMembersResponseDetails {

	path := fmt.Sprintf("guild/%s/members", guild_id)

	response := GuildMembersResponseDetails{}
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
