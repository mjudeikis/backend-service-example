package validator

import "github.com/mjudeikis/backend-service/pkg/models"

// ValidateUserRequest is a function that validates the user request

// TODO: THis is very basic validator. We should be using tags/lables/comments based validation
// in example: `json:"error,omitempty" validate:"string"`

func ValidateUserRequest(user models.UserUpdateRequest) (bool, models.APIError) {
	if user.Email == "" { // TODO: check if email is valid
		return false, models.ErrEmailRequired
	}
	if user.Password == "" { // TODO: check if password lenght is valid
		return false, models.ErrEmailRequired
	}
	// ...
	return true, ""
}
