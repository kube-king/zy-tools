/*
Copyright 2022 qkp Authors

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

package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zy-tools/config"
)

var (
	R *Response
)

type Response struct {
}

//初始化对象
func init() {
	R = new(Response)
}

// Success 请求成功返回的对象
func (r *Response) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, config.ResponseBody{
		Code:    http.StatusOK,
		Message: "操作成功",
		Data:    data,
	})
}

// Error 请求失败返回的对象
func (r *Response) Error(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, config.ResponseBody{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	})
}

// Custom 自定义请求返回对象
func (r *Response) Custom(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, config.ResponseBody{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
