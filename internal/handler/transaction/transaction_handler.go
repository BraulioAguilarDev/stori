package transactionhdlr

import (
	"encoding/csv"
	"net/http"
	"stori/internal/ports"
	"stori/pkg/cloud"
	"stori/pkg/response"
	"stori/pkg/validator"

	"github.com/gin-gonic/gin"
)

const BucketName = "gopher-stori-files"

type TransactionHdlr struct {
	TransactionSrv ports.TransactionServicePort
	AccountSrv     ports.AccountServicePort
	AccountS3      ports.AccountS3ServicePort
	S3             cloud.BucketClient
}

func ProvideTransactionHandler(
	TxnSrv ports.TransactionServicePort,
	AccountSrv ports.AccountServicePort,
	AccountS3 ports.AccountS3ServicePort,
	Client cloud.BucketClient,
) *TransactionHdlr {
	return &TransactionHdlr{
		TransactionSrv: TxnSrv,
		AccountSrv:     AccountSrv,
		AccountS3:      AccountS3,
		S3:             Client,
	}
}

type Parameters struct {
	AccountID string `json:"account_id" validate:"required"`
	File      string `json:"file"`
}

func (hdl *TransactionHdlr) ExecuteProcessHdlr(ctx *gin.Context) {
	input := Parameters{
		AccountID: ctx.Params.ByName("account"),
	}

	errs := validator.ValidateStructure(input)
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.FailureMappingErrors(errs))
		return
	}

	// Getting account info
	email := hdl.AccountSrv.GetEmail(input.AccountID)
	if len(email) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	filename, err := hdl.AccountS3.GetFileByAccountID(input.AccountID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Reading txns file
	file, err := hdl.S3.FetchObject(ctx, BucketName, filename)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.Failure(err.Error()))
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.Failure(err.Error()))
		return
	}

	// Sending to service
	if err := hdl.TransactionSrv.Create(input.AccountID, email, records); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.Failure(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(nil))
}
