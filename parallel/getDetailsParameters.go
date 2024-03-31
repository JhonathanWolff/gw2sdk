package parallel

import (
	"encoding/json"
	"time"
)

type GetRequestDetailsParameters interface {
	GetDetails(parameters map[string]string) interface{}
}

func RunRequestDetailsParameters(
	parameters map[string]string, channel chan []byte,
	request GetRequestDetailsParameters,
) {
	response := request.GetDetails(parameters)
	res2B, _ := json.Marshal(response)
	channel <- res2B
}

func ParallelGetDetailsParameters(
	list_requests []GetRequestDetailsParameters,
	parameters map[string]string,
) [][]byte {

	var routines int = 0
	channel := make(chan []byte, len(list_requests))

	var array_response [][]byte = make([][]byte, 0)
	for request := range list_requests {
		routines += 1
		go RunRequestDetailsParameters(parameters, channel, list_requests[request])

		if routines%5 == 0 {
			time.Sleep(time.Second * 1)
		}

	}

	print("Waiting for the channels")
	for len(list_requests) != len(channel) {
		time.Sleep(time.Second * 1)
	}
	close(channel)

	print("parsing response")
	for value := range channel {
		array_response = append(array_response, value)
	}
	return array_response

}
