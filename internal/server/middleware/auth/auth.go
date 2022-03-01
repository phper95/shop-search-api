package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-search-api/config"
	"shop-search-api/internal/server/api"
	"strings"
)

var AppSecret string

/**
appKey     = "xxx"
secretKey  = "xxx"
encryptParamStr = "param_1=xxx&param_2=xxx&ak="+appKey+"&ts=xxx"

// 自定义验证规则
sn = MD5(secretKey + encryptParamStr + appKey)
*/

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := api.Gin{C: c}
		// 签名信息
		authorization := c.GetHeader(config.HeaderAuthField)
		if authorization == "" {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中缺少 Authorization 参数")),
			)
			return
		}

		// 时间信息
		date := c.GetHeader(configs.HeaderSignTokenDate)
		if date == "" {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中缺少 Date 参数")),
			)
			return
		}

		// 通过签名信息获取 key
		authorizationSplit := strings.Split(authorization, " ")
		if len(authorizationSplit) < 2 {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中 Authorization 格式错误")),
			)
			return
		}

		key := authorizationSplit[0]

		data, err := i.authorizedService.DetailByKey(c, key)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(err),
			)
			return
		}

		if data.IsUsed == authorized.IsUsedNo {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New(key + " 已被禁止调用")),
			)
			return
		}

		if len(data.Apis) < 1 {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New(key + " 未进行接口授权")),
			)
			return
		}

		if !whiteListPath[c.Path()] {
			// 验证 c.Method() + c.Path() 是否授权
			table := urltable.NewTable()
			for _, v := range data.Apis {
				_ = table.Append(v.Method + v.Api)
			}

			if pattern, _ := table.Mapping(c.Method() + c.Path()); pattern == "" {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.AuthorizationError,
					code.Text(code.AuthorizationError)).WithError(errors.New(c.Method() + c.Path() + " 未进行接口授权")),
				)
				return
			}
		}

		ok, err := signature.New(key, data.Secret, configs.HeaderSignTokenTimeout).Verify(authorization, date, c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(err),
			)
			return
		}

		if !ok {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中 Authorization 信息错误")),
			)
			return
		}
	}
	if code != e.SUCCESS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})

		c.Abort()
		return
	}

	c.Next()
}
