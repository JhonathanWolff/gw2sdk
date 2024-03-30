package unauthenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type Guild struct {
	Gw2sdk Connection.GW2sdk
}

type GuildResponse []int

type GuildResponseDetails struct {
	Level     int    `json:"level"`
	Motd      string `json:"motd"`
	Influence int    `json:"influence"`
	Aetherium int    `json:"aetherium"`
	Resonance int    `json:"resonance"`
	Favor     int    `json:"favor"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	Emblem    struct {
		Background struct {
			ID     int   `json:"id"`
			Colors []int `json:"colors"`
		} `json:"background"`
		Foreground struct {
			ID     int   `json:"id"`
			Colors []int `json:"colors"`
		} `json:"foreground"`
		Flags []string `json:"flags"`
	} `json:"emblem"`
}

func (a *Guild) GetDetails(id int) GuildResponseDetails {

	response := GuildResponseDetails{}
	path := fmt.Sprintf("guild/%d", id)
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
