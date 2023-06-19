package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	switch strings.ToLower(request.RequestContext.HTTP.Method) {
	case "get":
		return handleGet(request)
	default:
		return MethodNotAllowed()
	}
}

func main() {
	lambda.Start(Handler)
}
