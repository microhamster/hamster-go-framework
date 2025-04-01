package apiServer

const (
	API_CODE_SUCCESS             = 0   // 成功返回
	API_CODE_NONCE_ERROR         = 100 // 编号错误
	API_CODE_PARAMS_ERROR        = 110 // 参数错误
	API_CODE_AUTHORIZE_ERROR     = 120 // 授权错误
	API_CODE_SIGNATURE_ERROR     = 130 // 签名错误
	API_CODE_REQUEST_LIMIT_ERROR = 140 // 请求超限

)
