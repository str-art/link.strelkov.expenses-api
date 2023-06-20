package handlers

import "github.com/aws/aws-lambda-go/events"

type Handler func(request events.APIGatewayProxyRequest)(events.APIGatewayProxyResponse)
