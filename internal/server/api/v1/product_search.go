package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"shop-search-api/internal/pkg/errcode"
	"shop-search-api/internal/server/api/api_response"
)

func ProductSearch(c *gin.Context) {
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}
	appG := api_response.Gin{C: c}
	appG.ResponseOk(errcode.ErrCodes.ErrNo, nil)
}
