package api

import (
	"context"
	"net/http"
	accounthdlr "stori/internal/handler/account"
	profilehdlr "stori/internal/handler/profile"
	s3hdlr "stori/internal/handler/s3"
	transactionhdl "stori/internal/handler/transaction"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

type Stori struct {
	TransactionHandler *transactionhdl.TransactionHdlr
	ProfileHandler     *profilehdlr.ProfileHdlr
	AccountHandler     *accounthdlr.AccountHdlr
	AccountS3Handler   *s3hdlr.S3Hdlr
	Router             *gin.Engine
	ginLambda          *ginadapter.GinLambda
}

// Custom CORS
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

func (api *Stori) SetupRouter() {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors(),
	)

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20

	api.Router = router
	api.Router.POST("/signup", api.ProfileHandler.SignUpHandler)
	api.Router.POST("/accounts", api.AccountHandler.CreateHandler)
	api.Router.POST("/upload", api.AccountS3Handler.UploadToS3AndSaveHandler)
	api.Router.POST("/transactions", api.TransactionHandler.ExecuteProcessHdlr)
}

func (api *Stori) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	api.ginLambda = ginadapter.New(api.Router)
	return api.ginLambda.ProxyWithContext(ctx, req)
}
