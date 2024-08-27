/*
Copyright 2024 Faw Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	R *Response
)

type Response struct {
	statusCodeMapping StatusCodeMapping
}

// 初始化对象
func init() {
	R = new(Response)
	R.statusCodeMapping = BaseStatusCodeMapping
}

func result(ctx *gin.Context, status string, code int, message string, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseBody{
		Data:    data,
		Message: message,
		Code:    code,
	})
}

func (r *Response) RegistryStatusCode(statusCodeMapping StatusCodeMapping) {
	for k, v := range statusCodeMapping {
		r.statusCodeMapping[k] = v
	}
}

// GetMessage 自定义返回信息函数方法
func (r *Response) GetMessage(code int) string {
	msg, ok := r.statusCodeMapping[code]
	if ok {
		return msg
	}
	return r.statusCodeMapping[Error]
}

// Success 请求成功返回的对象
func (r *Response) Success(ctx *gin.Context) {
	result(ctx, "Success", Success, r.GetMessage(Success), nil)
}

// SuccessWithData 请求成功,携带数据
func (r *Response) SuccessWithData(ctx *gin.Context, data interface{}) {
	result(ctx, "Success", Success, r.GetMessage(Success), data)
}

// SuccessWithMessage 请求成功,自定义消息内容
func (r *Response) SuccessWithMessage(ctx *gin.Context, message string) {
	result(ctx, "Success", Success, message, nil)
}

// SuccessWithDetailed 请求成功,自定义消息并携带数据
func (r *Response) SuccessWithDetailed(ctx *gin.Context, message string, data interface{}) {
	result(ctx, "Success", Success, message, data)
}

// Error 请求失败
func (r *Response) Error(ctx *gin.Context) {
	result(ctx, "Error", Error, r.GetMessage(Error), nil)
}

// ErrorWithMessage 请求失败,自定义消息
func (r *Response) ErrorWithMessage(ctx *gin.Context, message string) {
	result(ctx, "Error", Error, message, nil)
}

// ErrorWithCode 请求失败,自定义消息
func (r *Response) ErrorWithCode(ctx *gin.Context, code int) {
	result(ctx, "Error", code, r.GetMessage(code), nil)
}

// Custom 请求失败,自定义消息
func (r *Response) Custom(ctx *gin.Context, code int, msg string) {
	result(ctx, "Error", code, msg, nil)
}
