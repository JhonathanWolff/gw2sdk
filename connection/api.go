package connection

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_URL string = "https://api.guildwars2.com/v2/"

type GW2sdk struct {
	Token       string
	ApiResponse ApiResponse
}

func ParseParameters(parameters map[string]string) string {

	if parameters == nil {
		return ""
	}

	first := true
	var parsed string = ""
	operator := "&"
	for k, v := range parameters {

		if first {
			first = false
			operator = "?"
		}
		parsed += operator + k + "=" + v

	}

	return parsed

}

func (gw2 *GW2sdk) Retrieve(
	endpoint_path string,
	parameters map[string]string,
	target interface{},
) {

	request_url := API_URL + endpoint_path + ParseParameters(parameters)
	fmt.Println(request_url)

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

	bytes, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		panic(string(bytes))
	}

	//	fmt.Println(string(bytes))

	defer res.Body.Close()

	json.Unmarshal(bytes, &target)

}
