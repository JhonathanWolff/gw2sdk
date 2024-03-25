package achivements

import (
	Connection "gw2sdk/connection"
)

type Category struct {
	Gw2sdk Connection.GW2sdk
}

type CategoryResponse []int

type CategoryResponseDetails []struct {
	ID           int64                 `json:"id"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Order        int64                 `json:"order"`
	Icon         string                `json:"icon"`
	Achievements []CategoryAchievement `json:"achievements"`
	Tomorrow     []CategoryAchievement `json:"tomorrow"`
}

type CategoryAchievement struct {
	ID             int64          `json:"id"`
	RequiredAccess RequiredAccess `json:"required_access"`
	Flags          []string       `json:"flags"`
	Level          []int64        `json:"level"`
}

type RequiredAccess struct {
	Product   string `json:"product"`
	Condition string `json:"condition"`
}

// https://wiki.guildwars2.com/wiki/API:2/achievements/categories
func (a *Category) Get(parameters map[string]string) interface{} {

	if parameters == nil && len(parameters) == 0 {
		response := CategoryResponse{}
		a.Gw2sdk.Retrieve("achievements/categories", nil, &response)
		return response
	}

	response := CategoryResponseDetails{}
	a.Gw2sdk.Retrieve("achievements/categories?v=2022-03-23T19:00:00.000Z", parameters, &response)
	return response

}
