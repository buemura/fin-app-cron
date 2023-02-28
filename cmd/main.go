package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/buemura/golang-cron-jobs/internal"
)

func init() {
	internal.LoadEnv()
}

func execute(ctx context.Context, event interface{}) error {
	internal.UpdateAllExpenses()
	return nil
}

func main() {
	lambda.Start(execute)
}
