package api

import (
	"github.com/gin-gonic/gin"
	"shop-search-api/internal/server/middleware/auth"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	//防止panic发生，返回500
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	apiv1.Use(auth.Auth())

}
