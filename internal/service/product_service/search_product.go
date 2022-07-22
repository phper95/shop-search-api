package auth_service

import (
	"shop-search-api/global"
	"shop-search-api/internal/repo/es/product_repo"
	"shop-search-api/internal/server/api/api_response"
)

func (s service) SearchProduct(ctx *api_response.Gin, keyword string) (date []*product_repo.ProductIndex, err error) {
	s.es.Query(ctx, global.ProductIndexName,[]cate)
}
