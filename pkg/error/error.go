// description: 通用错误
//
// author: vignetting
// time: 2021/5/10

package error

import "strconv"

type VError struct {
	// 错误码，与 http 状态码一致
	Code int `json:"code"`

	// 错误诱因
	Cause string `json:"cause"`
}

func (ve VError) Error() string {
	return strconv.Itoa(ve.Code) + " : " + ve.Cause
}

// description: 因用户端错误操作操作而产生的 error
func Fail(cause string) error {
	return &VError{Code: 400, Cause: cause}
}

// description: 服务端错误运行而产生的 error
func Error(cause string) error {
	return &VError{Code: 500, Cause: cause}
}

// description: 任意 error
func Any(code int, cause string) error {
	return &VError{Code: code, Cause: cause}
}
