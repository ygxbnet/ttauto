package ttapi

const (
	// 基础 URL
	baseURL = "http://tiptime-api.com"

	// 登陆相关 URL
	getCaptchaRequestURL = "/api/v1/captcha/request"

	sendSMSURL = "/web/api/v2/login/code"

	loginURL = "/web/api/login"

	refreshTokenURL = "/api/v1/login"

	// 用户相关 URL
	getUserInfoURL = "/web/api/account/message/loading"
	// 签到
	signInURL = "/web/api/account/sign_in"
)
