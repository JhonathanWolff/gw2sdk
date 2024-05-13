package main

import (
	"encoding/json"
	"fmt"
	"gw2sdk/connection"
	unauthenticated "gw2sdk/unauthenticated/home_instance"
	"os"
)

func main() {

	var gw2sdk connection.GW2sdk = connection.GW2sdk{Token: os.Getenv("GW2_TOKEN")}
	info_not_logon := &unauthenticated.HomeNodes{Gw2sdk: gw2sdk}

	response := info_not_logon.Get()
	res2B, _ := json.Marshal(response)
	fmt.Println(string(res2B))

}
