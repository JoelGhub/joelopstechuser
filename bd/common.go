package bd

import (
	"Go_Project/Proyecto_2/secretm"
	"Go_Project/Proyecto_2/models"
	"fmt"
	"os"


)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret () error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}