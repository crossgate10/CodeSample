package main

import (
	"apis/routers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin start")
	runtime.GOMAXPROCS(runtime.NumCPU())

	r := gin.Default()
	r = routers.InitRoutes(r)

	ginLambda = ginadapter.New(r)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("req: %+v", req)
	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
