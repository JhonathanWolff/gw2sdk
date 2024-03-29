package authenticated

import (
	Connection "gw2sdk/connection"
)

type Recipes struct {
	Gw2sdk Connection.GW2sdk
}

type RecipesResponse []int

func (a *Recipes) Get() RecipesResponse {

	response := RecipesResponse{}
	a.Gw2sdk.Retrieve("account/recipes", nil, &response)

	return response

}
