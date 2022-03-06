package config

const (
	AppName = "shop-search-api"
	//Header 中传递的参数字段，其携带的值为接口的签名
	HeaderAuthField = "Authorization"

	//Header 中传递的参数字段，其携带的值为发起请求的时间，用于签名失效验证
	HeaderAuthDateField = "Authorization-Date"

	RunModeDev  = "dev"
	RunModeProd = "prod"

	DefaultMysqlClient = "default-mysql"
	DefaultRedisClient = "default-redis"

	RedisKeyPrefixSignature       = "sign:"
	RedisSignatureCacheSeconds    = 300
	HeaderSignTokenTimeoutSeconds = 300
)
