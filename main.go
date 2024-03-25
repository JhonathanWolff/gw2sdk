package main

import (
	"fmt"
	"gw2sdk/achivements"
	"gw2sdk/connection"
	"os"
)

func main() {
	var gw2sdk connection.GW2sdk = connection.GW2sdk{Token: os.Getenv("GW2_TOKEN")}
	var achivement = achivements.Category{Gw2sdk: gw2sdk}
	var parameter map[string]string = make(map[string]string)
	parameter["ids"] = "1,2"

	//parameter := "18DB115A-8637-4290-A636-821362A3C4A8"

	fmt.Println(achivement.Get(parameter))

}
