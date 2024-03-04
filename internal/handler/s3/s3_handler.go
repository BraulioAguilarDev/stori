package s3hdlr

import (
	"errors"
	"fmt"
	"net/http"
	"stori/internal/core/domain"
	"stori/internal/ports"
	"stori/pkg/response"
	"stori/pkg/validator"

	"github.com/gin-gonic/gin"
)

type S3Hdlr struct {
	AccountSrv   ports.AccountServicePort
	AccountS3Srv ports.AccountS3ServicePort
}

func ProvideS3Handler(account ports.AccountServicePort, s3 ports.AccountS3ServicePort) *S3Hdlr {
	return &S3Hdlr{
		AccountSrv:   account,
		AccountS3Srv: s3,
	}
}

type S3Parameters struct {
	AccountID string `json:"account_id" validate:"required"`
	URL       string `json:"url" validate:"required"`
}

func (hdl *S3Hdlr) UploadS3(ctx *gin.Context) {
	var input S3Parameters

	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Failure(err.Error()))
		return
	}

	errs := validator.ValidateStructure(input)
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.FailureMappingErrors(errs))
		return
	}

	fmt.Println("paso", input.AccountID)
	// Getting account info
	account, err := hdl.AccountSrv.GetByID(input.AccountID)
	if err != nil {
		fmt.Println("GET Accounr ERROR:", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if account == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New("getting account error"))
		return
	}

	request := domain.AccountS3DTO{
		AccountID: input.AccountID,
		URL:       input.URL,
	}

	s3, err := hdl.AccountS3Srv.Create(&request)
	if err != nil {
		fmt.Println("ERROR:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.Failure(err.Error()))
		return
	}

	s3URL := fmt.Sprintf("%v", s3.URL)
	ctx.JSON(http.StatusOK, response.Success(s3URL))
}
