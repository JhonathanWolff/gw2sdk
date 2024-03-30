package authenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type GuildStorage struct {
	Gw2sdk Connection.GW2sdk
}

type GuildStorageResponseDetails []struct {
	Id   int `json:"id"`
	Rank int `json:"count"`
}

func (a *GuildStorage) GetDetails(guild_id string) GuildStorageResponseDetails {

	path := fmt.Sprintf("guild/%s/storage", guild_id)

	response := GuildStorageResponseDetails{}
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
