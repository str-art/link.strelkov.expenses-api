package handlers

import (
	"expenses.strelkov.link/lambda/application"
	"expenses.strelkov.link/lambda/helpers"
	"expenses.strelkov.link/lambda/models"
)



func GetCategoryRequest(app *application.Application,respond application.MakeResponse)(){

	categoryName,categoryError := helpers.GetCategoryFromPath(app.Request)

	if categoryError != nil {
		respond(nil,categoryError)
		return
	}

	category,err := models.FindCategoryByName(*categoryName,app.Database)

	if err != nil {
		respond(nil,err)
		return
	}

	time,timeError := helpers.GetDateFromPath(app.Request)

	if timeError != nil {
		respond(nil,timeError)
		return
	}

	expenses,getError := category.GetExpensesByDate(*time,app.Database)


	if getError != nil {
		respond(nil,err)
		return
	}


	for index,expense := range expenses {
		expenses[index].Name = helpers.CCaseToString(expense.Name)
	}

	respond(expenses,nil)
}

