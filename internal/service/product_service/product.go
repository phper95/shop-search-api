package product_service

import (
	"shop-search-api/internal/repo/es/product_repo"
)

type Product struct {
	KeyWord  string
	IsHot    *int
	IsNew    *int
	IsGood   *int
	PageNum  int
	PageSize int
}

func (s *Product) SearchProduct() (date []*product_repo.ProductIndex, err error) {
	//global.ES.Query()
	return nil, nil
}
