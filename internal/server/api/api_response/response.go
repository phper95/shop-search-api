package api_response

import (
	"github.com/gin-gonic/gin"
	"shop-search-api/internal/pkg/errcode"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) ResponseOk(errCode errcode.ErrCode, data interface{}) {
	g.C.JSON(errCode.HTTPCode, Response{
		Success: true,
		Code:    errCode.Code,
		Msg:     errCode.Desc,
		Data:    data,
	})
}

func (g *Gin) ResponseErr(errCode errcode.ErrCode) {
	g.C.JSON(errCode.HTTPCode, Response{
		Success: false,
		Code:    errCode.Code,
		Msg:     errCode.Desc,
		Data:    nil,
	})
}
