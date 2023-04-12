package code

// 1000～1999 区间表示系统错误
// 2000～2999 区间表示用户错误
var (
	Success              = int64(200)  // 执行成功
	SystemError          = int64(1001) // 系统错误
	ResourceNotAvailable = int64(1002) // 服务端资源不可用
	RemoteServiceError   = int64(1003) // 远程服务出错
	ParamInvalid         = int64(1004) // 错误:参数错误
	SystemBusy           = int64(1005) // 任务过多系统繁忙
	Timeout              = int64(1006) // 任务超时
	RPCError             = int64(1007) // RPC错误
	BadRequest           = int64(1008) // 非法请求
	SqlError             = int64(2001) // SQL错误

	//状态码对应的信息
	statusText = map[int64]string{
		Success:              "成功",
		SystemError:          "系统错误",
		ResourceNotAvailable: "服务端资源不可用",
		RemoteServiceError:   "远程服务出错",
		ParamInvalid:         "参数错误",
		SystemBusy:           "任务过多系统繁忙",
		Timeout:              "任务超时",
		RPCError:             "RPC错误",
		BadRequest:           "非法请求",
		SqlError:             "SQL错误",
	}
)

func StatusText(code int64) string {
	return statusText[code]
}
