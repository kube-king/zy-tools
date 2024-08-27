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

package request

import (
	"github.com/gin-gonic/gin"
	"zy-tools/pkg/common/response"
	"zy-tools/pkg/common/valid"
)

// ValidStruct 数据字段验证
func ValidStruct(c *gin.Context, data valid.ValidatorInterface) bool {
	result := valid.V.Struct(data)
	if result != "" {
		response.R.Custom(c, response.InvalidParams, result)
		return false
	}
	return true
}

// ShouldBindJson 绑定json数据
func ShouldBindJson(c *gin.Context, data valid.ValidatorInterface) bool {
	if err := c.ShouldBindJSON(data); err != nil {
		response.R.ErrorWithCode(c, response.InvalidParams)
		return false
	}
	if !ValidStruct(c, data) {
		return false
	}
	return true
}

// ShouldBindQuery 绑定json数据
func ShouldBindQuery(c *gin.Context, data valid.ValidatorInterface) bool {
	if err := c.ShouldBindQuery(data); err != nil {
		response.R.ErrorWithCode(c, response.InvalidParams)
		return false
	}
	if !ValidStruct(c, data) {
		return false
	}
	return true
}

// ShouldBindYaml 绑定yaml数据
func ShouldBindYaml(c *gin.Context, data valid.ValidatorInterface) bool {
	if err := c.ShouldBindYAML(data); err != nil {
		response.R.ErrorWithCode(c, response.InvalidParams)
		return false
	}
	if !ValidStruct(c, data) {
		return false
	}
	return true
}

// ValidVar 验证单个变量由gin处理
func ValidVar(c *gin.Context, data interface{}, tag string, flag valid.ValidatorMessages, key ...string) bool {

	result := valid.V.Var(data, tag, flag, key...)
	if result != "" {
		response.R.Custom(c, response.InvalidParams, result)
		return false
	}
	return true
}

// ValidMap 验证map字段
func ValidMap(c *gin.Context, data map[string]interface{}, rule map[string]interface{}, flag valid.ValidatorMessages) bool {
	var message string
	message = valid.V.Map(data, rule, flag)
	if message != "" {
		response.R.Custom(c, response.InvalidParams, message)
		return false
	}
	return true
}
