package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.getMsg(),
		Data: data,
	}

	c.JSON(http.StatusOK, rd)
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &Response{
		Code: code,
		Msg:  code.getMsg(),
		Data: nil,
	}

	c.JSON(http.StatusInternalServerError, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}

	c.JSON(http.StatusInternalServerError, rd)
}
