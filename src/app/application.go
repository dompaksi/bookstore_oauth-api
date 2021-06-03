package app

import (
	"github.com/dompaksi/bookstore_oauth-api/src/http"
	"github.com/dompaksi/bookstore_oauth-api/src/repository/db"
	"github.com/dompaksi/bookstore_oauth-api/src/repository/rest"
	"github.com/dompaksi/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8001")
}
