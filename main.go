package main

import (
	"net/http"
	"numbergame/backend/utils"
	"numbergame/backend/endpoints"
	"cloud.google.com/go/firestore"
	"os"
)

var client *firestore.Client

func init() {
	utils.FirestoreInit()
}

func main() {
	http.HandleFunc("/getNumbers", utils.Logging(endpoints.GetNumbers))
	http.HandleFunc("/createUser", utils.Logging(endpoints.CreateUser))
	
	port := os.Getenv("PORT")
    if port == "" {
            port = "80"
    }
	http.ListenAndServe(":" + port, nil)
}
