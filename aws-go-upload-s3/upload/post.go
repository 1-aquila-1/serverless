package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/1-aquila-1/sdk-go/aws-sdk-go-v2/dynamodb/dynamodbinfra"
	"github.com/1-aquila-1/sdk-go/common"
	"github.com/1-aquila-1/sdk-go/webinfra"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func handlePost(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	bucket := os.Getenv("BUCKET")
	if len(bucket) == 0 {
		return webinfra.BandRequest(errors.New("Nome do bucket não localizado"))
	}

	req := InputDTO{}

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return webinfra.BandRequest(err)
	}

	buffer, err := base64.StdEncoding.DecodeString(req.Conteudo)

	if err != nil {
		return webinfra.BandRequest(err)
	}

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	extensao := strings.ReplaceAll(fileType, "image/", "")

	nome := req.Nome + "-" + common.Uuid10() + "." + extensao

	cfg, err := dynamodbinfra.AwsConfig()

	if err != nil {
		return webinfra.BandRequest(err)
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(nome),
		Body:        fileBytes,
		ContentType: aws.String(fileType),
	})

	if err != nil {
		return webinfra.BandRequest(err)
	}

	return webinfra.Ok(Response{
		Url: result.Location,
	})
}
