package endpoints

import (
	"encoding/json"
	"net/http"
	"numbergame/backend/utils"
)

//Endpoint for generating set of numbers and summary
func UserGetData(w http.ResponseWriter, r *http.Request) {
	var user utils.User
	ctx := r.Context()

	username := r.URL.Query().Get("username")

	db, err := utils.Client.Collection("users").Doc(username).Get(ctx)
	if err != nil {
		http.Error(w, "Incorrect name", http.StatusBadRequest)
        return
	}
	db.DataTo(&user)
	
	json.NewEncoder(w).Encode(user)
}