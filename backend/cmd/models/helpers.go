package models

import (
	"os"
	"strconv"
	"time"

	"expenses.strelkov.link/lambda/constants"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type withId interface {
	setId(id string)
}

func setId(model withId){
	id := strconv.FormatInt(time.Now().Unix(),10)
	model.setId(id)
}

func getTableName()(*string){
	return aws.String(os.Getenv(constants.DynamoTableName))
}

func getByEntityIdIndexName()(*string){
	return aws.String("ByEntityId")
}
