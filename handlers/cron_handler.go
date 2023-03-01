package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/buemura/golang-cron-jobs/internal"
)

type HandlerResponse struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	var resp HandlerResponse
	
	if request.RequestContext.HTTP.Method == "POST" && request.RequestContext.HTTP.Path == "/api/cron" {
		err := internal.UpdateAllExpenses()
		if err != nil {
			resp.Message = "Something went wrong"
			return events.LambdaFunctionURLResponse{
				Body:       fmt.Sprint(resp),
				StatusCode: http.StatusInternalServerError,
			}, nil
		}

		resp.Message = "Successfully updated expenses"
		return events.LambdaFunctionURLResponse{
			Body:       fmt.Sprint(resp),
			StatusCode: http.StatusOK,
		}, nil
	}

	resp.Message = "Route not found"
	return events.LambdaFunctionURLResponse{
		Body:       fmt.Sprint(resp),
		StatusCode: http.StatusNotFound,
	}, nil
}