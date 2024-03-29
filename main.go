package main

import (
	"fmt"
	authenticated "gw2sdk/authenticated/account"
	"gw2sdk/connection"
	"os"
)

func main() {
	var gw2sdk connection.GW2sdk = connection.GW2sdk{Token: os.Getenv("GW2_TOKEN")}
	account := &authenticated.Account{Gw2sdk: gw2sdk}
	// var parameter map[string]string = make(map[string]string)
	// parameter["ids"] = "1,2"

	fmt.Println(account.Get())

}
