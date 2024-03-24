package main

import (
	"fmt"
	"gw2sdk/account"
	"gw2sdk/connection"
	"os"
)

func main() {
	var gw2sdk connection.GW2sdk = connection.GW2sdk{Token: os.Getenv("GW2_TOKEN")}
	var account = account.Account{Gw2sdk: gw2sdk}
	result := account.Get()
	fmt.Println(result)

}
