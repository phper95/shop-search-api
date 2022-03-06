package v1

import (
	"github.com/gin-gonic/gin"
	"shop-search-api/internal/pkg/errcode"
	"shop-search-api/internal/server/api/api_response"
)

func ProductSearch(c *gin.Context) {
	appG := api_response.Gin{C: c}
	appG.ResponseOk(errcode.ErrCodes.ErrNo, nil)
}
