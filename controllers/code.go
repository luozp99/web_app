package controllers

type ResCode int

const (
	CodeSuccess ResCode = 10000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeInvalidationToken
	CodeInvalidationLogin
	CodeInvalidationTokenEmpty
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:                "success",
	CodeInvalidParam:           "参数无效",
	CodeUserExist:              "用户已经存在",
	CodeUserNotExist:           "用户不存在",
	CodeInvalidPassword:        "密码错误",
	CodeServerBusy:             "系统繁忙",
	CodeInvalidationToken:      "无效的token",
	CodeInvalidationLogin:      "登陆失效",
	CodeInvalidationTokenEmpty: "token为空",
}

func (c ResCode) getMsg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
