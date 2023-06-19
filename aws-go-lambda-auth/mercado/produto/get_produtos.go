package main

import (
	"github.com/aws/aws-lambda-go/events"
)

func handleGet(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	token := request.Headers["authorization"]
	return StatusOk(Produto{
		Codigo: "w23e -|" + token,
		Nome:   "Feijão",
	})
}
