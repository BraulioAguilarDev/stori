package transactionhdl

import (
	"stori/internal/ports"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	// transactionService
	txnsService ports.TransactionServicePort
}

func ProvideTransactionHandler(txnsSrv ports.TransactionServicePort) *TransactionHandler {

	return &TransactionHandler{
		txnsService: txnsSrv,
	}
}

func (hdl *TransactionHandler) ReceiveFileToProcessHandler(ctx *gin.Context) {

}
