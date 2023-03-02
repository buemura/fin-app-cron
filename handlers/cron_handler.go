package handlers

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/buemura/golang-cron-jobs/internal"
	"github.com/buemura/golang-cron-jobs/shared"
)

type HandlerResponse struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	if request.RequestContext.HTTP.Method == "POST" && request.RequestContext.HTTP.Path == "/api/cron" {
		err := internal.UpdateAllExpenses()
		if err != nil {
			return shared.HandleLambdaResponse(http.StatusInternalServerError, HandlerResponse{
				Message: "Something went wrong",
			})
		}

		return shared.HandleLambdaResponse(http.StatusOK, HandlerResponse{
			Message: "Successfully updated expenses",
		})
	}

	return shared.HandleLambdaResponse(http.StatusNotFound, HandlerResponse{
		Message: "Route not found",
	})
}