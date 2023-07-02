package main

import (
	"strings"
	"time"

	"github.com/1-aquila-1/sdk-go/webinfra"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	switch strings.ToUpper(request.RequestContext.HTTP.Method) {
	case "POST":
		return handlePost(request)
	default:
		return webinfra.MethodNotAllowed()
	}
}

func main() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	lambda.Start(Handler)
}
