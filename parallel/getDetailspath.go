package parallel

import (
	"encoding/json"
	"time"
)

type GetRequestDetails interface {
	GetDetails(value string) interface{}
}

func RunRequestDetailsPath(value string, channel chan []byte, request GetRequestDetails) {
	response := request.GetDetails(value)
	res2B, _ := json.Marshal(response)
	channel <- res2B
}

func ParallelGetDetailsPath(list_requests []GetRequestDetails, value string) [][]byte {

	var routines int = 0
	channel := make(chan []byte, len(list_requests))

	var array_response [][]byte = make([][]byte, 0)
	for request := range list_requests {
		routines += 1
		go RunRequestDetailsPath(value, channel, list_requests[request])

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
