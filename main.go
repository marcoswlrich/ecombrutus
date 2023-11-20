package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"github.com/marcoswlrich/ecombrutus/awsgo"
	"github.com/marcoswlrich/ecombrutus/bd"
	"github.com/marcoswlrich/ecombrutus/handlers"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(
	ctx context.Context,
	request events.APIGatewayV2HTTPRequest,
) (*events.APIGatewayProxyResponse, error) {
	awsgo.InitAWS()

	if !ValidateParameters() {
		panic(
			"Error nos parametros. Deve-se enviar 'SecretName', 'UrlPrefix'",
		)
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Handlers(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil
}

func ValidateParameters() bool {
	_, bringsParameters := os.LookupEnv("SecretName")
	if !bringsParameters {
		return bringsParameters
	}

	_, bringsParameters = os.LookupEnv("UrlPrefix")
	if !bringsParameters {
		return bringsParameters
	}

	return bringsParameters
}
