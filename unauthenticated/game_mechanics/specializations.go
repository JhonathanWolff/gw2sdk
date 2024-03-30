package unauthenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
)

type Specializations struct {
	Gw2sdk Connection.GW2sdk
}

type SpecializationsResponse []int

type SpecializationsResponseDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Profession  string `json:"profession"`
	Elite       bool   `json:"elite"`
	MinorTraits []int  `json:"minor_traits"`
	MajorTraits []int  `json:"major_traits"`
	Icon        string `json:"icon"`
	Background  string `json:"background"`
}

func (a *Specializations) Get() SpecializationsResponse {

	response := SpecializationsResponse{}
	a.Gw2sdk.Retrieve("specializations", nil, &response)

	return response

}

func (a *Specializations) GetDetails(id int) SpecializationsResponseDetails {

	response := SpecializationsResponseDetails{}
	path := fmt.Sprintf("specializations/%d", id)
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
