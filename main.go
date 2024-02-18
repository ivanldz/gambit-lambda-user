package main

import (
	"context"
	"fmt"
	"gambit-lambda-user/aws"
	"gambit-lambda-user/exceptions"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(RunLambda)
}

func RunLambda(
	ctx context.Context,
	event events.CognitoEventUserPoolsPostConfirmation,
) (events.CognitoEventUserPoolsPostConfirmation, error) {

	aws.InitAws()
	if !ValidParams() {
		fmt.Println("Error en los parametros. Debe enviar 'SecretName'")
		return event, exceptions.PARAM_NOT_PROVIDER
	}

	return event, nil
}

func ValidParams() bool {
	_, exist := os.LookupEnv("SecretName")
	return exist
}
