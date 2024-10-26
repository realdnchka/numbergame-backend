package endpoints

import (
	"encoding/json"
	"net/http"
	"numbergame/backend/utils"
	"time"
	"encoding/base64"
		"google.golang.org/api/iterator"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user utils.User
	
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    if user.Name == "" {
        http.Error(w, "Name is required", http.StatusBadRequest)
        return
    }
    iter := utils.Client.Collection("users").Where("username", "==", user.Name).Limit(1).Documents(ctx)
	defer iter.Stop()
	
    exists := false
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		exists = true
		break
	}
    
	if exists {
		http.Error(w, "Name is already taken", http.StatusForbidden)
		return
	}
    
    user.Token = EncodeToken(user.Name)
	utils.Client.Collection("users").NewDoc().Set(ctx, map[string]interface{}{
				"username": user.Name,
		        "highscore": 0,
		        "totalscores":  0,
				"token": user.Token,
		})
	
	userResponse := utils.UserRegisterResponse {
		Token: user.Token,
	}
	
	json.NewEncoder(w).Encode(userResponse)
}

func EncodeToken(username string) string {
 	currentTime := time.Now()
    timeString := currentTime.Format(time.RFC3339) + username

    return base64.StdEncoding.EncodeToString([]byte(timeString))
}