package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"expenses.strelkov.link/lambda/constants"
	"expenses.strelkov.link/lambda/helpers"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PaginatedList[Interface interface{}] struct {
	rows []Interface
	edge string
}


type expense struct {
	SK1 string `json:"-" dynamodbav:"SK1"`
	Amount int `json:"amount" dynamodbav:"amount"`
	Category string `json:"-" dynamodbav:"-"`
	Name string `json:"name" dynamodbav:"name"`
	Date string `json:"date" dynamodbav:"date"`
	SK2 string `json:"-" dynamodbav:"SK2"`
	Entity string `dynamodbav:"PK" json:"-"`
	ID string `json:"id" dynamodbav:"id"`
}

func (e *expense) Save(db *dynamodb.Client)(*expense,error){
	var expenseItem map[string]types.AttributeValue
	var err error

	expenseItem,err = attributevalue.MarshalMap(e)

	if err != nil {
		helpers.LogError(err)
		return nil,err
	} else {
		_,err = db.PutItem(context.TODO(),&dynamodb.PutItemInput{
			Item:                        expenseItem,
			TableName:                   getTableName(),
		})
	}

	if err != nil {
		helpers.LogError(err)
		return nil,err
	}

	return e,nil
}

func NewExpense(amount int, name,category string,date time.Time) (*expense){
	e := &expense{
		Amount: amount,
		Category: category,
		Name: helpers.StringToCCase(name),
		Date: date.Format(constants.DateFormat),
		Entity: expense_entity,
	}

	e.setSortKey()

	return e
}

func (e *expense) setSortKey()(*expense){

	if e.ID == ""{
		setId(e)
	}

	if e.SK1 == ""{
		e.setSK1()
	}

	if e.SK2 == ""{
		e.setSK2()
	}

	return e
}

func (e *expense) setId(id string)(){
	e.ID = id
}

func (e *expense) setSK1()(*expense){
	e.SK1 = fmt.Sprintf("%s%s%s%s%s",e.Category,constants.Delimiter,e.Date,constants.Delimiter,e.ID)
	return e
}

func (e *expense) setSK2()(*expense){
	e.SK2 = fmt.Sprintf("%s%s%s%s%s",e.Name,constants.Delimiter,e.Date,constants.Delimiter,e.ID)
	return e
}


func FindExpenseById(id string,db dynamodb.Client)(*expense,error){
	var err error
	var response *dynamodb.QueryOutput
	var expenses []*expense

	keyExp := expression.Key(pk).Equal(expression.Value(expense_entity))
	idExp := expression.Key(id).Equal(expression.Value(id))

	exp,err := expression.NewBuilder().WithKeyCondition(keyExp.And(idExp)).Build()

	if err != nil {
		helpers.LogError(err)
		return nil,err
	}

	response,err = db.Query(context.TODO(),&dynamodb.QueryInput{
		TableName:                 getTableName(),
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
		IndexName:                 getByEntityIdIndexName(),
	})

	if len(response.Items) == 0 {
		err = errors.New(fmt.Sprintf("Not found:Expense with id %s is not found.",id))
		return nil,err
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items,expenses)

	if err != nil {
		helpers.LogError(err)
		return nil,err
	}

	return expenses[0],nil
}

func (e *expense) GetId()(string){
	return e.ID
}