package utils

import (
	"context"
	"log"
	"net/http"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
    "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func getSecret() [] byte{
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("\nCannot setup client: %v", err)
	}
	defer client.Close()
	
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
                Name: "projects/862600208132/secrets/api-key-auth-go/versions/1",
	}
    
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("\nCannot access to secret: %v", err)
	}

	return result.Payload.Data
}

func ApiKey(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		
		if apiKey != string(getSecret())  {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		
		f(w, r)
	}
}