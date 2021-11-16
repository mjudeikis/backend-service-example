package models

// UserUpdateRequest defines user update request payload
type UserUpdateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRespose defines user response payload
type UserRespose struct {
	UserId string `json:"user_id"`
}

type UserInternal struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
