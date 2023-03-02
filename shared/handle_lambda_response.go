package shared

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func HandleLambdaResponse(statusCode int, data any) (events.LambdaFunctionURLResponse, error) {
	respJSON, err := json.Marshal(data)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			Body:       string(err.Error()),
			StatusCode: statusCode,
		}, nil
	}
	
	return events.LambdaFunctionURLResponse{
		Body:       string(respJSON),
		StatusCode: statusCode,
	}, nil
}