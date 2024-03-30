package authenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
	"time"
)

type GuildTeams struct {
	Gw2sdk Connection.GW2sdk
}

type GuildTeamsResponseDetails []struct {
	ID      int `json:"id"`
	Members []struct {
		Name string `json:"name"`
		Role string `json:"role"`
	} `json:"members"`
	Name      string `json:"name"`
	State     string `json:"state"`
	Aggregate struct {
		Wins       int `json:"wins"`
		Losses     int `json:"losses"`
		Desertions int `json:"desertions"`
		Byes       int `json:"byes"`
		Forfeits   int `json:"forfeits"`
	} `json:"aggregate"`
	Ladders struct {
		Unranked struct {
			Wins       int `json:"wins"`
			Losses     int `json:"losses"`
			Desertions int `json:"desertions"`
			Byes       int `json:"byes"`
			Forfeits   int `json:"forfeits"`
		} `json:"unranked"`
	} `json:"ladders"`
	Games []struct {
		ID      string    `json:"id"`
		MapID   int       `json:"map_id"`
		Started time.Time `json:"started"`
		Ended   time.Time `json:"ended"`
		Result  string    `json:"result"`
		Team    string    `json:"team"`
		Scores  struct {
			Red  int `json:"red"`
			Blue int `json:"blue"`
		} `json:"scores"`
		RatingType string `json:"rating_type"`
	} `json:"games"`
	Seasons []struct {
		ID     string `json:"id"`
		Wins   int    `json:"wins"`
		Losses int    `json:"losses"`
		Rating int    `json:"rating"`
	} `json:"seasons"`
}

func (a *GuildTeams) GetDetails(guild_id string) GuildTeamsResponseDetails {

	path := fmt.Sprintf("guild/%s/teams", guild_id)

	response := GuildTeamsResponseDetails{}
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
