package auth_service

import (
	"encoding/json"
	"gitee.com/phper95/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"shop-search-api/config"
	"shop-search-api/internal/repo/mysql/auth_repo"
	"shop-search-api/internal/server/api/api_response"
)

// CacheAuthorizedData 缓存结构
type CacheAuthorizedData struct {
	Key    string `json:"key"`     // 调用方 key
	Secret string `json:"secret"`  // 调用方 secret
	IsUsed int32  `json:"is_used"` // 调用方启用状态 1=启用 -1=禁用
}

type cacheApiData struct {
	Method string `json:"method"` // 请求方式
	Api    string `json:"api"`    // 请求地址
}

func (s *service) DetailByKey(ctx *api_response.Gin, key string) (cacheData *CacheAuthorizedData, err error) {
	// 查询缓存
	cacheKey := config.RedisKeyPrefixSignature + key
	var ok bool
	ok, err = s.cache.Exists(cacheKey)
	if err != nil {
		errors.Wrap(err, "redis error key"+cacheKey+"error "+err.Error())
	}
	if !ok {
		// 查询调用方信息
		authorizedInfo, err := auth_repo.NewQueryBuilder().
			WhereIsDeleted("=", -1).
			WhereBusinessKey("=", key).
			First(s.db)

		if err != nil {
			return nil, err
		}
		logger.Debug("authorizedInfo", zap.Any("", authorizedInfo))
		// 设置缓存 data
		cacheData = new(CacheAuthorizedData)
		cacheData.Key = key
		cacheData.Secret = authorizedInfo.BusinessSecret
		cacheData.IsUsed = authorizedInfo.IsUsed

		cacheDataByte, _ := json.Marshal(cacheData)

		err = s.cache.Set(cacheKey, string(cacheDataByte), config.RedisSignatureCacheSeconds)
		if err != nil {
			return nil, err
		}

		return cacheData, nil
	}

	value, err := s.cache.GetStr(cacheKey)
	if err != nil {
		return nil, err
	}

	cacheData = new(CacheAuthorizedData)
	err = json.Unmarshal([]byte(value), cacheData)
	if err != nil {
		return nil, err
	}

	return

}
