package main

import (
	"golox-playground/code"
	"golox-playground/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	codeService := code.NewService()
	codeHandler := code.NewHandler(
		codeService,
	)

	server := gin.Default()
	server.Use(
		response.CORSMiddleware(),
	)

	server.POST("/run", codeHandler.Run)
	server.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	server.Run(":8080")

}
