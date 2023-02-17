package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://your-api-url.com"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != "GET" {
		return events.APIGatewayProxyResponse{
			StatusCode: 405,
		}, nil
	}

	// Create an HTTP request based on the API Gateway request
	httpRequest, err := http.NewRequest(request.HTTPMethod, fmt.Sprintf("%s/v1/albums", DefaultHTTPGetAddress), nil)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Set the headers for the HTTP request
	for key, value := range request.Headers {
		httpRequest.Header.Set(key, value)
	}

	// Send the HTTP request and get the response
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Convert the HTTP response to an API Gateway response
	apiGatewayResponse := events.APIGatewayProxyResponse{
		StatusCode: httpResponse.StatusCode,
		//Headers:    httpResponse.Header.,
		Body: "",
	}
	if httpResponse.Body != nil {
		defer httpResponse.Body.Close()
		bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		apiGatewayResponse.Body = string(bodyBytes)
	}

	return apiGatewayResponse, nil
}

func main() {
	lambda.Start(handler)
}
