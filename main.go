package main

import (
	"Go_Project/Proyecto_2/awsgo"
	"Go_Project/Proyecto_2/models"
	"Go_Project/Proyecto_2/bd"
	"context"
	"fmt"
	"os"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	
	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros. debe enviar 'SecretName'")
		err := errors.New("error en los parametros debe enviar SecretName")
		return event, err
	}
	var datos models.Signup

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("email = "+datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = "+datos.UserUUID)
		}
	}
	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("error "+err.Error())
		return event, err
	}

	
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
