package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/modles"
	"web_app/translate"
)

func SignUpHandler(c *gin.Context) {

	var user modles.UserSignUp

	err := c.ShouldBindJSON(&user)
	if err != nil {
		zap.L().Error("参数解析异常 err", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, translate.RemoveTopStruct(errs.Translate(translate.Trans)))
		//c.JSON(http.StatusInternalServerError, response.ErrorMsg(translate.RemoveTopStruct(errs.Translate(translate.Trans))))
		return
	}

	if err = logic.SignUp(&user); err != nil {
		//c.JSON(http.StatusInternalServerError, response.ErrorMsg(err.Error()))
		ResponseError(c, CodeUserExist)
		return
	}

	//c.JSON(http.StatusOK, response.Success())
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	var loginUser modles.LoginUser

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		zap.L().Error("参数解析异常 err", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//c.JSON(http.StatusInternalServerError, response.ErrorMsg(err.Error()))
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, translate.RemoveTopStruct(errs.Translate(translate.Trans)))
		//c.JSON(http.StatusInternalServerError, response.ErrorMsg(translate.RemoveTopStruct(errs.Translate(translate.Trans))))
		return
	}

	err = logic.LoginUser(&loginUser)
	if err != nil {
		ResponseError(c, CodeInvalidPassword)
		//c.JSON(http.StatusInternalServerError, response.ErrorMsg(err.Error()))
		return
	}
	ResponseSuccess(c, nil)

	//c.JSON(http.StatusOK, response.SuccessMsg("登陆成功"))

}
