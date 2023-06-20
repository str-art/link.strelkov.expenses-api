package application

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"expenses.strelkov.link/lambda/constants"
	"expenses.strelkov.link/lambda/helpers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var httpCodeStatusMap = map[string]int{
	"Bad request":http.StatusBadRequest,
	"Not found":http.StatusNotFound,
	"Unauthorized":http.StatusUnauthorized,
}

type Application struct {
	Request events.APIGatewayProxyRequest
	Database *dynamodb.Client
	responseChannel chan interface{}
	errorChannel chan error
}


type MakeResponse func(body interface{},err error)

type ApplicationRequestHandler func (
	application *Application,
	makeResponse MakeResponse,
)()

type errorMessage struct {
	Message string `json:"message"`
}

func (a *Application)closeDatabase(){
	a.Database = nil
}

func handleError(err error)(events.APIGatewayProxyResponse,error){
	message := err.Error()

	if len(message) == 0 {
		return handleResponse(
			errorMessage{
				Message: constants.DefaultErrorMessage,
			},
			http.StatusInternalServerError)
	}
	splitted := strings.SplitN(message,":",2)
	
	if len(splitted) < 2 {
		return handleResponse(
			errorMessage{
				Message: message,
			},
			http.StatusInternalServerError,
		)
	}

	status := splitted[0]

	code, found := httpCodeStatusMap[status]

	if !found {
		return handleResponse(
			errorMessage{
				Message: splitted[1],
			},
			http.StatusInternalServerError,
		)
	}

	return handleResponse(
		errorMessage{
			Message: splitted[1],
		},
		code,
	)

}

func handleResponse(body interface{},code int)(events.APIGatewayProxyResponse,error){
	json,err := json.Marshal(body)

	if err != nil {
		helpers.LogError(err)
	}


	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body: string(json),
		Headers: map[string]string{
			"Content-type":"application/json",
		},
	},nil
}

func NewApplication(request events.APIGatewayProxyRequest)(*Application,error){
	awsConfig,err := config.LoadDefaultConfig(context.TODO())
	if err != nil{
		helpers.LogError(err)
		return nil,err
	}
	return &Application{
		Request: request,
		Database: dynamodb.NewFromConfig(awsConfig),
		responseChannel: make(chan interface{}),
		errorChannel: make(chan error),
	},nil
}

func NewRequestHandler(handler ApplicationRequestHandler)(func (request events.APIGatewayProxyRequest)(events.APIGatewayProxyResponse,error)){
	
	return NewHandlerLogger(
		func(request events.APIGatewayProxyRequest)(events.APIGatewayProxyResponse, error) {
			application,bootStrapError := NewApplication(request)

			if bootStrapError != nil {
				helpers.LogError(bootStrapError)
				return handleError(bootStrapError)
			}

			makeResponse := func (body interface{},err error)(){
				application.closeDatabase()
				application.responseChannel <- body
				application.errorChannel <- err
			}

			go handler(application,makeResponse)

			body := <- application.responseChannel
			err := <- application.errorChannel

			if err != nil{
				helpers.LogError(err)
				return handleError(err)
			}

			return handleResponse(body,http.StatusOK)
	})
}
