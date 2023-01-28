package code

var (
	Success            = int64(0)    // 执行成功
	Fail               = int64(1)    // 执行失败
	ParamIsInvalid     = int64(1001) // 参数无效
	ParamIsBlank       = int64(1002) // 参数为空
	ParamTypeBindError = int64(1003) // 参数类型错误
	//状态码对应的信息
	statusText = map[int64]string{
		Success:            "Success",
		Fail:               "Fail",
		ParamIsInvalid:     "Param Is Invalid",
		ParamIsBlank:       "Param Is Blank",
		ParamTypeBindError: "Param Type Bind Error",
	}
)

func StatusText(code int64) string {
	return statusText[code]
}
