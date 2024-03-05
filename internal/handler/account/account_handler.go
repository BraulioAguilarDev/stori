package accounthdlr

import (
	"errors"
	"net/http"
	"stori/internal/core/domain"
	"stori/internal/ports"
	"stori/pkg/response"
	"stori/pkg/validator"

	"github.com/gin-gonic/gin"
)

type AccountHdlr struct {
	AccountSrv   ports.AccountServicePort
	ProfileSrv   ports.ProfileServicePort
	AccountS3Srv ports.AccountS3ServicePort
}

// Account params
type Parameters struct {
	Owner  string `json:"owner" validate:"required"`
	Bank   string `json:"bank" validate:"required"`
	Type   string `json:"type" validate:"required"`
	Number int    `json:"number" validate:"required"`
	User   string `json:"user_id" validate:"required"`
}

func ProvideAccountHandler(
	account ports.AccountServicePort,
	profile ports.ProfileServicePort,
	s3 ports.AccountS3ServicePort,
) *AccountHdlr {
	return &AccountHdlr{
		AccountSrv: account,
		ProfileSrv: profile,
	}
}

func (hdl *AccountHdlr) CreateHandler(ctx *gin.Context) {
	var input Parameters
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Failure(err.Error()))
		return
	}

	accErrors := validator.ValidateStructure(input)
	if len(accErrors) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.FailureMappingErrors(accErrors))
		return
	}

	// Getting profile info
	profile, err := hdl.ProfileSrv.GetByID(input.User)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if profile == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New("getting profile error"))
		return
	}

	request := domain.AccountDTO{
		Owner:     input.Owner,
		Bank:      input.Bank,
		Type:      input.Type,
		Number:    input.Number,
		ProfileID: profile.ID,
	}

	account, err := hdl.AccountSrv.Create(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.Failure(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(account))
}
