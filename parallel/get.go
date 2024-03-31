package parallel

import (
	"encoding/json"
	"time"
)

type GetRequest interface {
	Get() interface{}
}

func RunRequest(channel chan []byte, request GetRequest) {
	response := request.Get()
	res2B, _ := json.Marshal(response)
	channel <- res2B
}

func ParallelGet(list_requests []GetRequest) [][]byte {

	var routines int = 0
	channel := make(chan []byte, len(list_requests))

	var array_response [][]byte = make([][]byte, 0)
	for request := range list_requests {
		routines += 1
		go RunRequest(channel, list_requests[request])

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
