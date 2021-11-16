package service

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mjudeikis/backend-service/pkg/models"
	"github.com/mjudeikis/backend-service/pkg/validator"
	"go.uber.org/zap"
)

func (s *ApiService) registerUser(w http.ResponseWriter, r *http.Request) {
	s.log.Debug("registerUser")
	var updateUserRequest models.UserUpdateRequest
	if err := read(r, &updateUserRequest); err != nil {
		s.log.Error("failed to unmarshal", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok, aErr := validator.ValidateUserRequest(updateUserRequest)
	if !ok {
		result, err := json.Marshal(models.CloudError{
			StatusCode:     http.StatusBadRequest,
			CloudErrorBody: aErr.Error(),
		})
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, string(result))
			return
		}

	}

	if _, ok := s.store[updateUserRequest.Email]; ok {
		result, err := json.Marshal(models.CloudError{
			StatusCode:     http.StatusConflict,
			CloudErrorBody: "user already exists",
		})
		if err != nil {
			s.log.Error("failed to marshal", zap.Error(err))
			writeResponse(w, http.StatusInternalServerError, string(result))
			return
		}
		// handle here....
		return
	}

	user := models.UserInternal{
		ID:       uuid.New().String(),
		Email:    updateUserRequest.Email,
		Username: updateUserRequest.Username,
		Password: updateUserRequest.Password,
	}

	s.store[updateUserRequest.Email] = user

	result, err := json.Marshal(models.UserRespose{
		UserId: user.ID,
	})
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "ups...")
		return
	}
	writeResponse(w, http.StatusOK, string(result))
}
