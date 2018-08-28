package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/negroni"
	"github.com/urfave/negroni"
	"gitlab.com/FernandoCagale/serverless-clean/api/handler"
	"gitlab.com/FernandoCagale/serverless-clean/api/infra"
)

var initialized = false

var negroniLambda *negroniadapter.NegroniAdapter

func handlers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if !initialized {
		router := infra.NewRouter()

		handler.MakeTaskHandlers(router, handler.MakeTaskGorm())

		n := negroni.Classic()
		n.UseHandler(router)

		negroniLambda = negroniadapter.New(n)
		initialized = true
	}

	return negroniLambda.Proxy(req)
}

func main() {
	lambda.Start(handlers)
}
