package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func MethodNotAllowed() (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode:      http.StatusInternalServerError,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func StatusOk(obj interface{}) (events.APIGatewayV2HTTPResponse, error) {
	body, err := json.Marshal(obj)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode:      http.StatusInternalServerError,
			IsBase64Encoded: false,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, err
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
