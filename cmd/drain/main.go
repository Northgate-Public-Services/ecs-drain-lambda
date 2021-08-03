package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	drain "github.com/Northgate-Public-Services/ecs-drain-lambda"
)

func main() {
	lambda.Start(drain.HandleRequest)
}
