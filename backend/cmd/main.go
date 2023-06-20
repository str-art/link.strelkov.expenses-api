package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"expenses.strelkov.link/lambda/application"
	"expenses.strelkov.link/lambda/handlers"
)



func main(){

	handlersMap := map[string]application.ApplicationRequestHandler{
		"read_category":handlers.GetCategoryRequest,
		"add_expense":handlers.AddExpense,
		"get_categories":handlers.GetCategories,
	}

	handler := handlersMap[os.Getenv("HANDLER")]

	lambda.Start(application.NewRequestHandler(handler))
}