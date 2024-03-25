package main

import (
	"fmt"
	"gw2sdk/achivements"
	"gw2sdk/connection"
	"os"
)

func main() {
	var gw2sdk connection.GW2sdk = connection.GW2sdk{Token: os.Getenv("GW2_TOKEN")}
	var achivement = achivements.Achivement{Gw2sdk: gw2sdk}
	var parameter map[string]string = make(map[string]string)

	parameter["ids"] = "2481,5010"

	fmt.Println(achivement.Get(parameter))

}
