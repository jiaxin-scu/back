// description: 通用返回值
//
// author: vignetting
// time: 2021/5/10

package result

type Result struct {

	// 状态码
	// 2xx 表示请求成功，3xx 表示重定向，4xx 表示客户端错误，5xx 表示服务端错误
	Code int `json:"code"`

	// 状态码描述信息
	Msg string `json:"msg"`

	// 数据
	Data interface{} `json:"data"`
}

// description: 用于请求处理成功后的通用返回值
func Ok(message string, data interface{}) *Result {
	return &Result{
		Code: 200,
		Msg:  message,
		Data: data,
	}
}

// description: 客户端的请求错误
func Fail(message string, data interface{}) *Result {
	return &Result{
		Code: 400,
		Msg:  message,
		Data: data,
	}
}

// description: 未授权，用户未登录或登录信息错误
func Unauthorized(message string, data interface{}) *Result {
	return &Result{
		Code: 401,
		Msg:  message,
		Data: data,
	}
}

// description: 用户无权访问
func Forbidden(message string, data interface{}) *Result {
	return &Result{
		Code: 403,
		Msg:  message,
		Data: data,
	}
}

// description: 服务端发送错误
func Error(message string, data interface{}) *Result {
	return &Result{
		Code: 500,
		Msg:  message,
		Data: data,
	}
}

// description: 服务不可用，服务处于维护、超载、限流等时可以使用该状态
func Unavailable(message string, data interface{}) *Result {
	return &Result{
		Code: 503,
		Msg:  message,
		Data: data,
	}
}

// description: 通用
func Any(code int, message string, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  message,
		Data: data,
	}
}
