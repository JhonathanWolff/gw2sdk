package achivements

import (
	Connection "gw2sdk/connection"
)

type Achivement struct {
	Gw2sdk Connection.GW2sdk
}

func (a *Achivement) Get() interface{} {

	var map_response map[string]interface{} = make(map[string]interface{})
	result := a.Gw2sdk.Retrieve(map_response)

	return result

}
