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

package controller

import (
	"github.com/gin-gonic/gin"
)

// Router 实例化router结构体，可使用该对象点出首字母大写的方法(包外调用)
var Router router

// 创建router结构体
type router struct{}

// InitApiRouter 初始化路由规则，创建测试api接口
func (r *router) InitApiRouter(router *gin.Engine) {

	doc := router.Group("/document")
	doc.POST("/file-convert", FileConvert)  // 文件转pdf
	doc.POST("/json-to-excel", JsonToExcel) // 文件转pdf

	//devel := router.Group("/api/development")
	//devel.GET("/check-json", CheckJson) // 检查json格式

	router.GET("/download", Download) // 下载文件
}
