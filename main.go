package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"golang-coupang-backend.com/m/handler"
	"golang-coupang-backend.com/m/repository"
	"golang-coupang-backend.com/m/service"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	port = ":8080"
)

func main() {
	context := context.TODO()
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(cfg)

	// Create DynamoDB client
	dynamoClient := dynamodb.NewFromConfig(cfg)

	// 의존성 생성 및 연결
	repo := repository.NewDynamoParcelRepository(dynamoClient)
	parcelService := service.NewParcelService(repo)

	r := gin.Default()
	handler := handler.NewParcelHandler(parcelService)
	handler.RegisterParcelRoutes(r)
	r.Run(port)
}
