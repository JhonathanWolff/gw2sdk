package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Traits struct {
	Gw2sdk Connection.GW2sdk
}

type TraitsResponse []int

type TraitsResponseDetails struct {
	ID          int    `json:"id"`
	Tier        int    `json:"tier"`
	Order       int    `json:"order"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slot        string `json:"slot"`
	Facts       []struct {
		Text        string `json:"text,omitempty"`
		Type        string `json:"type"`
		Icon        string `json:"icon"`
		Value       int    `json:"value,omitempty"`
		Target      string `json:"target,omitempty"`
		Duration    int    `json:"duration,omitempty"`
		Status      string `json:"status,omitempty"`
		Description string `json:"description,omitempty"`
		ApplyCount  int    `json:"apply_count,omitempty"`
	} `json:"facts"`
	Specialization int    `json:"specialization"`
	Icon           string `json:"icon"`
}

func (a *Traits) Get() TraitsResponse {

	response := TraitsResponse{}
	a.Gw2sdk.Retrieve("traits", nil, &response)

	return response

}

func (a *Traits) GetDetails(parameters map[string]string) TraitsResponseDetails {

	response := TraitsResponseDetails{}
	a.Gw2sdk.Retrieve("traits", parameters, &response)

	return response

}
