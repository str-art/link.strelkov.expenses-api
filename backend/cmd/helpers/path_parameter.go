package helpers

import (
	"errors"
	"os"
	"time"

	"expenses.strelkov.link/lambda/constants"
	"github.com/aws/aws-lambda-go/events"
)

func GetCategoryFromPath(request events.APIGatewayProxyRequest)(*string,error){
	categoryName := request.PathParameters[os.Getenv(constants.CategoryPathParam)]
	if categoryName == "" {
		return nil,errors.New(BadRequestMessage(constants.InvalidCategoryNameMessage))
	}

	return &categoryName,nil
}

func GetDateFromPath(request events.APIGatewayProxyRequest)(*time.Time,error){
	date := request.PathParameters[os.Getenv(constants.DatePathParameter)]

	if date == ""{
		date = time.Now().Format(constants.DateFormat)
	}

	time,parseError := time.Parse(constants.DateFormat,date)

	if parseError != nil {
		return nil, errors.New(BadRequestMessage(constants.InvalidDateFormatMessage))
	}

	return &time,nil
}