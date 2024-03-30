package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type GuildPermissions struct {
	Gw2sdk Connection.GW2sdk
}

type GuildPermissionsResponse []string

type GuildPermissionsResponseDetails []struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a *GuildPermissions) Get() GuildPermissionsResponse {

	response := GuildPermissionsResponse{}
	a.Gw2sdk.Retrieve("guild/permissions", nil, &response)

	return response

}

func (a *GuildPermissions) GetDetails(
	parameters map[string]string,
) GuildPermissionsResponseDetails {

	response := GuildPermissionsResponseDetails{}
	a.Gw2sdk.Retrieve("guild/permissions", parameters, &response)

	return response

}
