package main

import (
	"net/http"
	"numbergame/backend/utils"
	"numbergame/backend/endpoints"
	"cloud.google.com/go/firestore"
)

var client *firestore.Client

func init() {
	utils.FirestoreInit()
}

func main() {
	http.HandleFunc("/getNumbers", utils.Logging(utils.ApiKey(endpoints.GetNumbers)))
	http.HandleFunc("/userCreate", utils.Logging(utils.ApiKey(endpoints.UserCreate)))
	http.HandleFunc("/isNameTaken", utils.Logging(utils.ApiKey(endpoints.IsNameTaken)))
	http.HandleFunc("/userLogin", utils.Logging(utils.ApiKey(endpoints.UserLogin)))
	
	http.ListenAndServe(":80", nil)
}
