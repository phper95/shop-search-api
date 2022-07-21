package auth_service

import (
	"gitee.com/phper95/pkg/cache"
	"gorm.io/gorm"
	"shop-search-api/internal/server/api/api_response"
)

type Service interface {
	DetailByKey(ctx *api_response.Gin, key string) (data *CacheAuthorizedData, err error)
}

type service struct {
	db    *gorm.DB
	cache *cache.Redis
}

func New(db *gorm.DB, cache *cache.Redis) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}
