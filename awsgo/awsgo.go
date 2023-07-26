package awsgo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/aws"
)

var Ctx context.Context
var Cfg context.Context
var err error


func InicializoAWS() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("eu-west-3"))

	if err != nil {
		panic("Error al cargar la configuración .aws/config "+err.Error())
	}
}