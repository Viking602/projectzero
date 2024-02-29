package errorx

import "errors"

// 五位数错误编码为应用自定义错误
// 50XXX的五位数错误编码为服务器端错误
// 40XXX的五位数错误编码为客户端错误
const (
	// CodeCheckLogin 未登录
	CodeNoLoginErr = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeNotFoundErr 资源不存在
	CodeNotFoundErr = 404
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeTokenGenErr token生成失败
	CodeTokenGenErr = 50002
	//CodeParamErr 通用参数错误
	CodeParamErr = 40001
	// CodeTokenErr token错误
	CodeTokenErr = 40002
	// CodeTokenExpire token过期
	CodeTokenExpire = 40003
)

var ErrMissingParameter = errors.New("missing parameter")
