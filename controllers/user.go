package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"web_app/logic"
	"web_app/modles"
	"web_app/response"
	"web_app/translate"
)

func SignUpHandler(c *gin.Context) {

	var user modles.UserSignUp

	err := c.ShouldBindJSON(&user)
	if err != nil {
		zap.L().Error("参数解析异常 err", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusInternalServerError, response.ErrorMsg(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorMsg(translate.RemoveTopStruct(errs.Translate(translate.Trans))))
		return
	}

	logic.SignUp(&user)

	c.JSON(http.StatusOK, response.Success())
}
