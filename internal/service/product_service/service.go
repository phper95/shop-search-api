package auth_service

import (
	"gitee.com/phper95/pkg/es"
	"shop-search-api/internal/repo/es/product_repo"
	"shop-search-api/internal/server/api/api_response"
)

type Service interface {
	SearchProduct(ctx *api_response.Gin, keyword string) (date []*product_repo.ProductIndex, err error)
}

type service struct {
	es *es.Client
}

func New(es *es.Client) Service {
	return &service{
		es: es,
	}
}
