package response

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"projectzero/pkg/errorx"
)

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = fmt.Sprintf("%+v", err)
	}
	return res
}

func Success(data interface{}) Response {
	return Response{
		Code: 0,
		Msg:  "操作成功",
		Data: data,
	}
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(errorx.CodeDBError, msg, err)
}

// ParamErr 参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(errorx.CodeParamErr, msg, err)
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) Response {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return ParamErr(
			"请求失败",
			err,
		)
	}

	return ParamErr("参数错误", err)
}
