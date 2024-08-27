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

package initialize

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"zy-tools/pkg/common/valid"
)

func InitValidator() {
	valid.V.Registry("contains_array", CheckContainsArr)
}

// CheckContainsArr 验证数据是否包含指定字符串元素
func CheckContainsArr(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()
	val := field.String()
	flag := false
	for _, elem := range strings.Split(param, ";") {
		if elem == val {
			flag = true
		}
	}
	return flag
}
