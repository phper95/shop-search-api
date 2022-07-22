package v1

import (
	"github.com/gin-gonic/gin"
	"shop-search-api/internal/pkg/errcode"
	"shop-search-api/internal/server/api/api_response"
	"shop-search-api/internal/service/product_service"
)

func ProductSearch(c *gin.Context) {
	keyword := c.Param("keyword")
	productService := product_service.Product{
		KeyWord: keyword,
		//IsHot:    0,
		//IsNew:    0,
		//IsGood:   0,
		PageNum:  0,
		PageSize: 20,
	}
	productService.SearchProduct()

	appG := api_response.Gin{C: c}
	appG.ResponseOk(errcode.ErrCodes.ErrNo, nil)
}
