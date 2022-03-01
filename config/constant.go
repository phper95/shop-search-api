package config

const (
	//Header 中传递的参数字段，其携带的值为接口的签名
	HeaderAuthField = "Authorization"

	//Header 中传递的参数字段，其携带的值为发起请求的时间，用于签名失效验证
	HeaderAuthDateField = "Authorization-Date"
)
