package firestore

import (
	"context"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// ContextForBot ... firestore information
type ContextForBot struct {
	Client  *firestore.Client
	Context context.Context
}

// Initialze ... sample firestore
func Initialze() *ContextForBot {

	// Use a service account
	var sa option.ClientOption
	if os.Getenv("ENV") == "LOCAL" {
		sa = option.WithCredentialsFile((os.Getenv("FIREBASEKEY")))
	} else {
		sa = option.WithCredentialsJSON([]byte(os.Getenv("FIREBASEKEY")))
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// defer client.Close()

	return &ContextForBot{
		Client:  client,
		Context: ctx,
	}

}
