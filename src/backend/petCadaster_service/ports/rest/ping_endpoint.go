package rest

import "github.com/gin-gonic/gin"

func (s Server) PingEndpoint(gctx *gin.Context) {
	gctx.JSON(200, gin.H{
		"Response": "Pong",
	})
}
