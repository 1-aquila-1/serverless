package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	IsAuthorized bool    `json:"isAuthorized"`
	Context      Context `json:"context"`
}

type Context struct {
	Token string `json:"token"`
}

func Handler(request events.APIGatewayV2HTTPRequest) (Response, error) {
	token := request.Headers["authorization"]
	autorizado := token == "1q2w3e4r5t6y7u8i9o0"

	resp := Response{
		IsAuthorized: autorizado,
		Context: Context{
			Token: "Token aluno",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
