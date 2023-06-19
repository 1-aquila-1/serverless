package main

import (
	"github.com/aws/aws-lambda-go/events"
)

func handleGet(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return StatusOk(Produto{
		Codigo: "w23e",
		Nome:   "Feijão",
	})
}
