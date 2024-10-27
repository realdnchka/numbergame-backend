package endpoints

import (
	"net/http"
	"numbergame/backend/utils"
	"encoding/json"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	var leaderboard []utils.LeaderboardEntry
	ctx := r.Context()
	
	cachedLeaderboard, found := utils.GetFromCache("leaderboard")
	if found {
		json.NewEncoder(w).Encode(cachedLeaderboard)
		return
	}

	iter := utils.Client.Collection("users").OrderBy("highscore", firestore.Desc).Limit(10).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Failed to retrieve documents", http.StatusInternalServerError)
			return
		}
		var entry utils.LeaderboardEntry
		if err := doc.DataTo(&entry); err != nil {
			http.Error(w, "Failed to parse document", http.StatusInternalServerError)
			return
		}
		leaderboard = append(leaderboard, entry)
	}
	
	utils.SetToCache("leaderboard", leaderboard)
	json.NewEncoder(w).Encode(leaderboard)
}