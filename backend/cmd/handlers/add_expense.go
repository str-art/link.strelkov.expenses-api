package handlers

import (
	"expenses.strelkov.link/lambda/application"
	"expenses.strelkov.link/lambda/helpers"
	"expenses.strelkov.link/lambda/models"
	"github.com/go-playground/validator"
)

type expenseDto struct {
	Amount int `json:"amount" validate:"min=1,required"`
	Name string `json:"name" validate:"required,min=1"`
}

type addExpenseResponse struct {
	Id string `json:"id"`
}


func AddExpense(app *application.Application,makeResponse application.MakeResponse){

	dto,parseError := application.ParseBody(expenseDto{},app.Request.Body)

	if parseError != nil {
		helpers.LogError(parseError)
		makeResponse(nil,parseError)
		return
	}

	validationError := validator.New().Struct(dto)

	if validationError != nil {
		helpers.LogError(validationError)
		makeResponse(nil,validationError)
		return
	}

	categoryName,categoryError := helpers.GetCategoryFromPath(app.Request)

	if categoryError != nil {
		helpers.LogError(categoryError)
		makeResponse(nil,categoryError)
		return
	}

	time,timeError := helpers.GetDateFromPath(app.Request)

	if timeError != nil {
		helpers.LogError(timeError)
		makeResponse(nil,timeError)
		return
	}

	category,findErr:= models.FindCategoryByName(*categoryName,app.Database)

	if findErr != nil {
		helpers.LogError(findErr)
		makeResponse(nil,findErr)
		return
	}


	expense,writeError := category.AddExpense(dto.Amount,*time,dto.Name,app.Database)

	if writeError != nil {
		helpers.LogError(writeError)
		makeResponse(nil,writeError)
		return
	} 

	makeResponse(expense,nil)
}