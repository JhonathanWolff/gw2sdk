package connection

import (
	"encoding/json"
	"io"
	"net/http"
)

const API_URL string = "https://api.guildwars2.com/v2/"

type GW2sdk struct {
	Token       string
	ApiResponse ApiResponse
}

func (gw2 *GW2sdk) Retrieve(target interface{}) {

	endpoint := "account"

	request_url := API_URL + endpoint

	client := &http.Client{}

	baerer_token := "Bearer " + gw2.Token

	req, err := http.NewRequest("GET", request_url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", baerer_token)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	bytes, _ := io.ReadAll(res.Body)
	json.Unmarshal(bytes, &target)

}
