package authenticated

import (
	Connection "gw2sdk/connection"
	"time"
)

type PvPGames struct {
	Gw2sdk Connection.GW2sdk
}
type PvPGamesResponse []struct {
	ID         string    `json:"id"`
	MapID      int       `json:"map_id"`
	Started    time.Time `json:"started"`
	Ended      time.Time `json:"ended"`
	Result     string    `json:"result"`
	Team       string    `json:"team"`
	Profession string    `json:"profession"`
	Scores     struct {
		Red  int `json:"red"`
		Blue int `json:"blue"`
	} `json:"scores"`
	RatingType   string `json:"rating_type"`
	RatingChange int    `json:"rating_change"`
	Season       string `json:"season"`
}
type PvpGamesIds []string

func (a *PvPGames) Get() PvpGamesIds {

	response := PvpGamesIds{}
	a.Gw2sdk.Retrieve("pvp/games", nil, &response)
	return response

}

func (a *PvPGames) GetDetails(game_id string) PvPGamesResponse {

	response := PvPGamesResponse{}
	var parameter map[string]string = make(map[string]string)
	parameter["ids"] = game_id
	a.Gw2sdk.Retrieve("pvp/games", parameter, &response)

	return response

}
