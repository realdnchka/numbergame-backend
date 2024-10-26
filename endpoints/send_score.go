package endpoints

import (
	"encoding/json"
	"net/http"
	"numbergame/backend/utils"
	"cloud.google.com/go/firestore"
)

//Endpoint for generating set of numbers and summary
func SendScore(w http.ResponseWriter, r *http.Request) {
	var user_current, user_db utils.User
	
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&user_current)
	if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    db, err := utils.Client.Collection("users").Doc(user_current.Name).Get(ctx)
    if err != nil {
	    http.Error(w, "Cannot find a name", http.StatusBadRequest)
	    return
    }
           
    db.DataTo(&user_db)
       
    if user_current.Token != user_db.Token {
	    http.Error(w, "Incorrect ", http.StatusForbidden)
		return
    }
	
    if user_current.HighScore == 0 {
	    http.Error(w, "Score is required", http.StatusBadRequest)
	    return
    }
    
	utils.Client.Collection("users").Doc(user_current.Name).Set(ctx, map[string]interface{}{
		        "highscore": user_current.HighScore,
		}, firestore.MergeAll)
}