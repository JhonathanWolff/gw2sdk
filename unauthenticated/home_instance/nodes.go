package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type HomeNodes struct {
	Gw2sdk Connection.GW2sdk
}

type HomeNodesResponse []string

type HomeNodesResponseDetails []struct {
	Id string `json:"id"`
}

func (a *HomeNodes) Get() HomeNodesResponse {

	response := HomeNodesResponse{}
	a.Gw2sdk.Retrieve("home/nodes", nil, &response)

	return response

}

func (a *HomeNodes) GetDetails(parameters map[string]string) HomeNodesResponseDetails {

	response := HomeNodesResponseDetails{}

	a.Gw2sdk.Retrieve("home/nodes", parameters, &response)

	return response

}
