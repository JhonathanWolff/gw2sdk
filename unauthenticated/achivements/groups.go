package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type AchivementGroups struct {
	Gw2sdk Connection.GW2sdk
}

type AchivementGroupsResponse []string

type AchivementGroupsResponseDetails struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	Categories  []int  `json:"categories"`
}

// https://wiki.guildwars2.com/wiki/API:2/achievements/groups
func (a *AchivementGroups) Get(guid *string) interface{} {

	if guid == nil {
		response := AchivementGroupsResponse{}
		a.Gw2sdk.Retrieve("achievements/groups", nil, &response)
		return response
	}

	response := AchivementGroupsResponseDetails{}
	a.Gw2sdk.Retrieve("achievements/groups/"+*guid, nil, &response)
	return response

}
