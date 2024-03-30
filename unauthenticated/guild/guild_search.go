package unauthenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type GuildSearch struct {
	Gw2sdk Connection.GW2sdk
}

type GuildSearchResponse []string

func (a *GuildSearch) GetDetails(guild_name string) GuildSearchResponse {

	response := GuildSearchResponse{}

	path := fmt.Sprintf("guild/search/%s", guild_name)
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
