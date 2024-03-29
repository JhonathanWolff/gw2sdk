package main

import (
	"encoding/json"
	"fmt"
	other "gw2sdk/authenticated"
	"gw2sdk/connection"
	"os"
)

func main() {
	var gw2sdk connection.GW2sdk = connection.GW2sdk{Token: os.Getenv("GW2_TOKEN")}
	characters := &other.TokenInformation{Gw2sdk: gw2sdk}
	// account := &authenticated.Account{Gw2sdk: gw2sdk}

	response := characters.Get()
	res2B, _ := json.Marshal(response)
	fmt.Println(string(res2B))

}
