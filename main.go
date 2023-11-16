package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(
	ctx context.Context,
	request events.APIGatewayV2HTTPRequest,
) (*events.APIGatewayProxyResponse, error) {
}
