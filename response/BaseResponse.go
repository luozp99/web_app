package response

type H map[string]interface{}

func SuccessMsg(msg interface{}) H {
	return H{
		"msg": msg,
	}
}

func Success() H {

	return H{
		"msg": "请求成功",
	}
}

func Error() H {

	return H{
		"msg": "程序处理异常",
	}
}

func ErrorMsg(msg interface{}) H {

	return H{
		"msg": msg,
	}
}
