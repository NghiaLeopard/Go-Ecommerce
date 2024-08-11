package response

type AuthResponse struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}