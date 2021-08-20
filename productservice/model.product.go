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
	Price     string
	Tag       string
	Quantity  string
	LocationN string
}

func addProduct(svc *dynamodb.Client, productId string, storeId string, title string, price string, tag string, quantity string, location string) {
	tableName := "products"
	produc := Product{ProductId: productId, StoreId: storeId, Title: title, Price: price, Tag: tag, Quantity: quantity, LocationN: location}
	av, err := attributevalue.MarshalMap(produc)

	if err != nil {
		fmt.Println(err)
	}
	svc.PutItem(context.TODO(), &dynamodb.PutItemInput{TableName: &tableName, Item: av})

}

func getProduct(svc *dynamodb.Client, productId string) Product {
	tableName := "products"

	produc := map[string]string{
		"ProductId": productId,
	}
	av, _ := attributevalue.MarshalMap(produc)
	res, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{TableName: &tableName, Key: av})
	if err != nil {
		fmt.Println(err)
	}
	product := Product{}
	attributevalue.UnmarshalMap(res.Item, &product)
	return product
}
