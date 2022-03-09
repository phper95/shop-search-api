package v1

import (
	"github.com/gin-gonic/gin"
	"shop-search-api/internal/pkg/errcode"
	"shop-search-api/internal/server/api/api_response"
)

type Product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ProductGet(c *gin.Context) {
	appG := api_response.Gin{C: c}
	p := Product{}
	//db.GetMysqlClient(config.DefaultMysqlClient).First(&p, 1)
	appG.ResponseOk(errcode.ErrCodes.ErrNo, p)
}