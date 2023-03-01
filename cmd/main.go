package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/buemura/golang-cron-jobs/handlers"
	"github.com/buemura/golang-cron-jobs/internal"
)

func init() {
	internal.LoadEnv()
}

func main() {
	lambda.Start(handlers.HandleRequest)
}