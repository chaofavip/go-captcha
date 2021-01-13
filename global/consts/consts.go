package consts

// 这里定义的常量，一般是具有错误代码+错误说明组成，一般用于接口返回
const (
	// 进程被结束
	ProcessKilled string = "收到信号，进程被结束"
	// 表单验证器前缀
	ValidatorPrefix              string = "Form_Validator_"
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	//服务器代码发生错误
	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误, "

	// CURD 常用业务状态码
	CurdStatusOkCode int = 200

	//验证码
	CaptchaGetParamsInvalidMsg    string = "获取验证码：提交的验证码参数无效,请检查验证码ID以及文件名后缀是否完整"
	CaptchaGetParamsInvalidCode   int    = -400350
	CaptchaCheckParamsInvalidMsg  string = "校验验证码：提交的参数无效，请确保提交的验证码ID和值有效"
	CaptchaCheckParamsInvalidCode int    = -400351
	CaptchaCheckOkMsg             string = "验证码校验通过"
	//CaptchaCheckOkCode            int    	= 200
	CaptchaCheckFailCode int    = -400355
	CaptchaCheckFailMsg  string = "验证码校验失败"
)
