package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hngprojects/hng_boilerplate_golang_web/external/request"
	"github.com/hngprojects/hng_boilerplate_golang_web/internal/models"
	"github.com/hngprojects/hng_boilerplate_golang_web/pkg/repository/storage"
	"github.com/hngprojects/hng_boilerplate_golang_web/pkg/repository/storage/postgresql"
	"github.com/hngprojects/hng_boilerplate_golang_web/services/actions"
	"github.com/hngprojects/hng_boilerplate_golang_web/services/actions/names"
	"github.com/hngprojects/hng_boilerplate_golang_web/utility"
	"gorm.io/gorm"
)

func ValidateSqueezeUserRequest(req models.SqueezeUserReq, db *gorm.DB) (models.SqueezeUserReq, int, error) {

	squeezeUser := models.SqueezeUser{}

	if req.Email != "" {
		req.Email = strings.ToLower(req.Email)
		formattedMail, checkBool := utility.EmailValid(req.Email)
		if !checkBool {
			return req, http.StatusUnprocessableEntity, fmt.Errorf("email address is invalid")
		}
		req.Email = formattedMail
		exists := postgresql.CheckExists(db, &squeezeUser, "email = ?", req.Email)
		if exists {
			return req, http.StatusBadRequest, errors.New("user already exists with the given email")
		}
	}

	return req, 0, nil
}

func CreateSqueeze(db *gorm.DB, extReq request.ExternalRequest, req models.SqueezeUserReq) (*models.SqueezeUser, error) {
	squeezeUser := &models.SqueezeUser{
		ID:        utility.GenerateUUID(),
		FirstName: strings.ToLower(req.FirstName),
		LastName:  strings.ToLower(req.LastName),
		Email:     req.Email,
	}

	err := squeezeUser.Create(db)

	if err != nil {
		return nil, err
	}

	squeezeUserReq := models.SendSqueeze{
		FirstName: squeezeUser.FirstName,
		LastName:  squeezeUser.LastName,
		Email:     squeezeUser.Email,
	}

	err = actions.AddNotificationToQueue(storage.DB.Redis, names.SendSqueeze, squeezeUserReq)
	if err != nil {
		extReq.Logger.Error("Failed to send welcome email:", err)
	}

	return squeezeUser, nil
}
