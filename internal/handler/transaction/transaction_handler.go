package transactionhdlr

import (
	"stori/internal/ports"

	"github.com/gin-gonic/gin"
)

type TransactionHdlr struct {
	service ports.TransactionServicePort
}

func ProvideTransactionHandler(srv ports.TransactionServicePort) *TransactionHdlr {
	return &TransactionHdlr{
		service: srv,
	}
}

func (hdl *TransactionHdlr) ReceiveFileToProcessHdlr(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}
