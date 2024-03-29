package authenticated

import (
	Connection "gw2sdk/connection"
)

type HomeNodes struct {
	Gw2sdk Connection.GW2sdk
}

type HomeNodesResponse []string

func (a *HomeNodes) Get() HomeNodesResponse {

	response := HomeNodesResponse{}
	a.Gw2sdk.Retrieve("account/home/nodes", nil, &response)

	return response

}
