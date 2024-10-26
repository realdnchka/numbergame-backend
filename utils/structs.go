package utils

// Struct for numbers and summary of them for core game
type NumberResponse struct {
	Numbers []int `json:"numbers"`
	Sum int `json:"sum"`
} 

type User struct {
	Name string `json:"username"`
	HighScore int `json:"highscore"`
	TotalScores int `json:"total_scores"`
	Token string `json:"token"`
}

type UserRegisterResponse struct {
	Token string `json:"token"`
}
