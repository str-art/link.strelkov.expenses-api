package handlers

import (
	"expenses.strelkov.link/lambda/application"
	"expenses.strelkov.link/lambda/helpers"
	"expenses.strelkov.link/lambda/models"
)



func GetCategories(app *application.Application,respond application.MakeResponse)(){
	categories,err := models.GetAllCategories(app.Database)

	if err != nil {
		respond(nil,err)
	}

	for index,category := range categories {
		categories[index].Name = helpers.CCaseToString(category.Name)
	}
	
	respond(categories,nil)
}