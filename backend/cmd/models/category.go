package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"expenses.strelkov.link/lambda/constants"
	"expenses.strelkov.link/lambda/helpers"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type category struct {
	Name string `dynamodbav:"name" json:"name"`
	Entity string `dynamodbav:"PK" json:"-"`
	SK1 string `dynamodbav:"SK1" json:"-"`
	ID string `dynamodbav:"id" json:"-"`
}

func NewCategory(name string)(*category){
	categoryName := helpers.StringToCCase(name)

	category := &category{
		Name: categoryName,
		Entity: category_entity,
		SK1: categoryName,
	}
	setId(category)
	return category
}

func (c *category)setId(id string){
	if c.ID == ""{
		c.ID = id
	}
}

func FindCategoryByName(name string,db *dynamodb.Client)(*category,error){
	category := category{}
	var err error
	var response *dynamodb.GetItemOutput
	response,err = db.GetItem(context.TODO(),&dynamodb.GetItemInput{
			Key:                      map[string]types.AttributeValue{
				pk:&types.AttributeValueMemberS{
					Value: category_entity,
				},
				sk1:&types.AttributeValueMemberS{
					Value: name,
				},
			},
			TableName:                aws.String(os.Getenv(constants.DynamoTableName)),
		})
	
	if err != nil {
		helpers.LogError(err)
		return nil,err
	}

	err = attributevalue.UnmarshalMap(response.Item,&category)

	if err != nil {
		helpers.LogError(err)
		return nil,err
	}

	if category.ID == ""{
		return NewCategory(name),nil
	}

	return &category,nil
}

func (c *category) GetExpensesByDate(date time.Time,db *dynamodb.Client)([]expense,error){
	entityEx := expression.Key(pk).Equal(expression.Value(expense_entity))
	skCondition := fmt.Sprintf("%s%s%s%s",c.Name,constants.Delimiter,date.Format(constants.DateFormat),constants.Delimiter)
	log.Printf("%s",skCondition)
	skEx := expression.Key(sk1).BeginsWith(skCondition)
	var expr expression.Expression
	var err error
	expr,err = expression.NewBuilder().WithKeyCondition(entityEx.And(skEx)).Build()
	var expenses []expense
	var response *dynamodb.QueryOutput
	if err != nil {
		helpers.LogError(err)
		return nil,err
	} else {
		
		response,err = db.Query(context.TODO(),&dynamodb.QueryInput{
			TableName:                 getTableName(),
			ExpressionAttributeValues: expr.Values(),
			ExpressionAttributeNames:  expr.Names(),
			KeyConditionExpression:    expr.KeyCondition(),
		})
	}

	if err != nil {
		helpers.LogError(err)
		return nil,err
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items,&expenses)

	if err != nil {
		helpers.LogError(err)
		return nil,err
	}
	
	if expenses == nil {
		expenses = make([]expense, 0)
	}

	return expenses,nil
}

func (c *category) AddExpense(amount int,date time.Time,name string,db *dynamodb.Client)(*expense,error){
	expense := NewExpense(amount,name,c.Name,date)
	var categoryItem map[string]types.AttributeValue

	setId(c)

	var err error

	categoryItem,err = attributevalue.MarshalMap(c)

	if err != nil{
		helpers.LogError(err)
		return nil,err
	} else {
		_,err = db.PutItem(context.TODO(),&dynamodb.PutItemInput{
			TableName:                   getTableName(),
			Item:						 categoryItem,
		})
	}

	if err != nil {
		helpers.LogError(err)
		return nil,err
	}

	return expense.Save(db)
}

func GetAllCategories(db *dynamodb.Client)([]category,error){
	pkExpr := expression.Key(pk).Equal(expression.Value(category_entity))
	var response *dynamodb.QueryOutput
	var err error
	var expr expression.Expression
	var categories []category
	expr,err = expression.NewBuilder().WithKeyCondition(pkExpr).Build()

	if err != nil {
		helpers.LogError(err)
		return nil,err
	} else {
		response,err = db.Query(context.TODO(),&dynamodb.QueryInput{
			TableName:                 getTableName(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
		})
	}

	if err != nil {
		helpers.LogError(err)
		return nil,err
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items,&categories)

	if categories == nil {
		categories = make([]category,0)
	}

	return categories,err
}