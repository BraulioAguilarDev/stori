package api

import (
	"context"
	"net/http"
	transactionhdl "stori/internal/handler/transactions"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

type APIStori struct {
	TxnsHandler *transactionhdl.TransactionHandler
	Router      *gin.Engine
	ginLambda   *ginadapter.GinLambda
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func (api *APIStori) SetupRouter() {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors(),
	)

	api.Router = router

	api.Router.POST("/transactions-by-account", api.TxnsHandler.ReceiveFileToProcessHandler)
}

func (api *APIStori) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	api.ginLambda = ginadapter.New(api.Router)
	return api.ginLambda.ProxyWithContext(ctx, req)
}
