package s3hdlr

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"
	"stori/internal/core/domain"
	"stori/internal/ports"
	"stori/pkg/cloud"
	"stori/pkg/response"
	"stori/pkg/validator"

	"github.com/gin-gonic/gin"
)

const BucketName = "gopher-stori-files"

type S3Hdlr struct {
	AccountSrv   ports.AccountServicePort
	AccountS3Srv ports.AccountS3ServicePort
	S3           cloud.BucketClient
}

func ProvideS3Handler(
	account ports.AccountServicePort,
	accountS3 ports.AccountS3ServicePort,
	client cloud.BucketClient) *S3Hdlr {
	return &S3Hdlr{
		AccountSrv:   account,
		AccountS3Srv: accountS3,
		S3:           client,
	}
}

type Parameters struct {
	AccountID string                `form:"account_id" validate:"required"`
	File      *multipart.FileHeader `form:"file"`
}

func (hdl *S3Hdlr) FindHandler(ctx *gin.Context) {
	accountId := ctx.Params.ByName("account")

	list, err := hdl.AccountS3Srv.Find(accountId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.Failure(err.Error()))
	}

	ctx.JSON(http.StatusOK, response.Success(list))
}

func (hdl *S3Hdlr) UploadToS3AndSaveHandler(ctx *gin.Context) {
	var input Parameters
	if err := ctx.Bind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Failure(err.Error()))
		return
	}

	errs := validator.ValidateStructure(input)
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.FailureMappingErrors(errs))
		return
	}

	// Getting account info
	account, err := hdl.AccountSrv.GetByID(input.AccountID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if account == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New("getting account error"))
		return
	}

	// First, we need to save in s3 bucket
	url, err := hdl.uploadObject(ctx, input.File)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("uploading file error"))
		return
	}

	request := domain.AccountS3DTO{
		AccountID: input.AccountID,
		URL:       url,
		Filename:  input.File.Filename,
	}

	s3, err := hdl.AccountS3Srv.Create(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.Failure(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(s3.URL))
}

func (hdl *S3Hdlr) uploadObject(ctx context.Context, fh *multipart.FileHeader) (string, error) {
	mpfile, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer mpfile.Close()

	url, err := hdl.S3.UploadObject(ctx, BucketName, fh.Filename, mpfile)
	if err != nil {
		return "", err
	}

	return url, nil
}
