package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	// Configurar a sessão do AWS SDK
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Endpoint: aws.String("http://localhost:4566"),
			Region:   aws.String("us-east-1"),
		},
		Profile: "default", // Perfil AWS
	}))

}
