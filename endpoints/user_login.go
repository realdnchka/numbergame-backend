package endpoints

import (
	"net/http"
	"numbergame/backend/utils"
	"encoding/json"
)

// Endpoint for checking is user registered by name
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var user_current, user_db utils.User
	
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&user_current)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	if user_current.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
        return
	}
	
	if user_current.Token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
        return
	}
	
	db, err := utils.Client.Collection("users").Doc(user_current.Name).Get(ctx)
    if err != nil {
    	http.Error(w, "Cannot find a name", http.StatusForbidden)
    	return
    }
        
    db.DataTo(&user_db)
    
    if user_current.Token == user_db.Token {
    	w.WriteHeader(http.StatusOK)
     	w.Write([]byte("Succesfully login"))
    } else {
   		http.Error(w, "Incorrect username or token", http.StatusForbidden)
     	return
    }
}