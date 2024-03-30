package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Masteries struct {
	Gw2sdk Connection.GW2sdk
}

type MasteriesResponse []int
type MasteriesResponseDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Requirement string `json:"requirement"`
	Order       int    `json:"order"`
	Background  string `json:"background"`
	Region      string `json:"region"`
	Levels      []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Instruction string `json:"instruction"`
		Icon        string `json:"icon"`
		PointCost   int    `json:"point_cost"`
		ExpCost     int    `json:"exp_cost"`
	} `json:"levels"`
}

func (a *Masteries) Get() MasteriesResponse {

	response := MasteriesResponse{}
	a.Gw2sdk.Retrieve("masteries", nil, &response)

	return response

}

func (a *Masteries) GetDetails(parameters map[string]string) MasteriesResponseDetails {

	response := MasteriesResponseDetails{}
	a.Gw2sdk.Retrieve("masteries", parameters, &response)

	return response

}
