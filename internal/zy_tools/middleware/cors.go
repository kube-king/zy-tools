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

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cores 处理跨域请求，支持options访问
func Cores() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Cores")
		// 获取请求方法
		method := context.Request.Method

		// 添加跨域响应头
		//context.Header("Content-Type", "application/json")
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Max-Age", "86400")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Header("Access-Control-Allow-Headers", "X-Token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		context.Header("Access-Control-Allow-Credentials", "false")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		context.Next()
	}

}
