package main

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type Response struct {
	Body       map[string]interface{}
	StatusCode int
}

func Client(Uri string) (Response, error) {
	client := resty.New()
	request, err := client.R().EnableTrace().Get(Uri)

	response := Response{}

	if err != nil {
		return response, err
	}
	var res map[string]interface{}
	if err := json.Unmarshal(request.Body(), &res); err != nil {
		return response, nil
	}

	response.Body = res
	response.StatusCode = request.StatusCode()
	return response, nil
}
