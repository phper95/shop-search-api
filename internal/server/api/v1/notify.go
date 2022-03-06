package v1

import (
	"github.com/gin-gonic/gin"
	"shop-search-api/internal/pkg/errcode"
)

func ProductMsgCallback(c *gin.Context) {
	appG := api_response.Gin{C: c}
	appG.ResponseOk(errcode.ErrCodes.ErrNo, nil)
}

func ProductMsgBatchCallback(c *gin.Context) {
	appG := api_response.Gin{C: c}
	appG.ResponseOk(errcode.ErrCodes.ErrNo, nil)
}
