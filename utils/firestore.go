package utils

import(
	"log"
	"context"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

var Client *firestore.Client
	
func FirestoreInit() {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
	  log.Fatalln(err)
	}
	
	Client, err = app.Firestore(ctx)
	if err != nil {
	  log.Fatalln(err)
	}
}

