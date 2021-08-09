package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Product struct {
	ProductId string
	StoreId   string
	Title     string
	Price     float32
	Tag       string
}

func addProduct(svc *dynamodb.Client, productId string, storeId string, title string, price float32, tag string) {
	tableName := "products"
	produc := Product{ProductId: productId, StoreId: storeId, Title: title, Price: price, Tag: tag}
	av, err := attributevalue.MarshalMap(produc)

	if err != nil {
		fmt.Println(err)
	}
	svc.PutItem(context.TODO(), &dynamodb.PutItemInput{TableName: &tableName, Item: av})
}
