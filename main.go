package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang-coupang-backend.com/m/handler"

	"github.com/aws/aws-sdk-go-v2/config"
)

const (
	port = ":8080"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(cfg)

	r := gin.Default()
	handler.RegisterParcelRoutes(r)
	r.Run(port)
}
