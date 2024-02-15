package infrastructure

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type FirebaseClient struct {
	C *auth.Client
}

func NewFirebaseClient() *FirebaseClient {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return &FirebaseClient{C: client}
}
