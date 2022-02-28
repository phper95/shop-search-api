package auth

import (
	"errors"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"shop-search-api/internal/config"
	"shop-search-api/internal/pkg/errcode"
	"shop-search-api/internal/pkg/sign"
	"shop-search-api/internal/server/api"
	"sort"
	"strconv"
	"strings"
	"time"
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
		ak := c.Query("ak")
		sn := c.Query("sn")
		ts := c.GetInt64("ts")
		if len(ak) == 0 {
			appG.ResponseErr(errcode.ErrCodes.ErrAppKey)
			c.Abort()
			return
		}
		if len(sn) == 0 {
			appG.ResponseErr(errcode.ErrCodes.ErrSign)
			c.Abort()
			return
		}
		if ts == 0 {
			appG.ResponseErr(errcode.ErrCodes.ErrParams)
			c.Abort()
			return
		}
		// 验证过期时间
		timestamp := time.Now().Unix()
		if ts > timestamp || timestamp-ts >= config.Cfg.App.AppSignExpire {
			appG.ResponseErr(errcode.ErrCodes.ErrAuthExpired)
			c.Abort()
			return
		}

		// 验证签名
		if sn == "" || sn != createSign(req) {
			return nil, errors.New("sn Error")
		}

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
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
}

// 创建签名
func createSign(params url.Values) string {
	// 自定义 MD5 组合
	return sign.MD5(AppSecret + createEncryptStr(params) + AppSecret)
}

func createEncryptStr(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" && k != "debug" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	return str
}
