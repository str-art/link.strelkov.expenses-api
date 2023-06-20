package application

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type genericLambdaHandler func(request events.APIGatewayProxyRequest)(events.APIGatewayProxyResponse,error)


func NewHandlerLogger(handler genericLambdaHandler)(genericLambdaHandler){
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {
		log.Printf("Recieved request \n")
		log.Printf("Body: %#v \n",request.Body)
		log.Printf("Path: %#v \n",request.PathParameters)
		log.Printf("Query: %#v \n",request.QueryStringParameters)
		result,err := handler(request)
		log.Printf("Result: %#v \n",result)
		log.Printf("Error: %#v \n",err)
		return result,err
	}
}