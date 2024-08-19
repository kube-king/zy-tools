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
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/wonderivan/logger"
	"net/http"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

var ValidMapping = map[string]string{
	"max":      "{0} 长度超过限制!",
	"required": "{0} 为必填字段!",
	"yaml":     "{0} yaml 格式错误!",
}

func InitValidEngine() {
	//注册翻译器
	cn := zh.New()
	uni = ut.New(cn, cn)
	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate = validator.New()
	//注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		logger.Error(err.Error())
	}
	//注册验证map
	for key, val := range ValidMapping {
		err = customRegister(validate, key, val)
		if err != nil {
			logger.Error(err.Error())
		}
	}
}

//自定义重构数据验证信息
func customRegister(validate *validator.Validate, key string, val string) error {
	err := validate.RegisterTranslation(key, trans, func(ut ut.Translator) error {
		return ut.Add(key, val, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(key, fe.Field())
		return t
	})
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// GetValidMessage 获取验证信息
func GetValidMessage(data any) string {
	var result = ""
	err := validate.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			result = e.Translate(trans)
		}
	}
	return result
}

//CheckValid 数据字段验证
func CheckValid(c *gin.Context, data any) bool {
	result := GetValidMessage(data)
	if result != "" {
		R.Custom(c, http.StatusInternalServerError, result)
		return false
	}
	return true
}

// BindJson 绑定json数据
func BindJson(c *gin.Context, data any) bool {
	if err := c.ShouldBindJSON(data); err != nil {
		R.Custom(c, http.StatusInternalServerError, err.Error())
		return false
	}
	if !CheckValid(c, data) {
		return false
	}
	return true
}

// BindQuery 绑定json数据
func BindQuery(c *gin.Context, data any) bool {
	if err := c.ShouldBindQuery(data); err != nil {
		R.Custom(c, http.StatusInternalServerError, "服务器内部错误!")
		return false
	}
	if !CheckValid(c, data) {
		return false
	}
	return true
}

// BindYaml 绑定yaml数据
func BindYaml(c *gin.Context, data any) bool {
	if err := c.ShouldBindYAML(data); err != nil {
		R.Custom(c, http.StatusInternalServerError, "服务器内部错误!")
		return false
	}
	if !CheckValid(c, data) {
		return false
	}
	return true
}
