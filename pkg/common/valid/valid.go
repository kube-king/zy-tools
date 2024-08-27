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

package valid

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"log"
)

// ValidatorInterface 验证器接口
type ValidatorInterface interface {
	GetValidMessage() ValidatorMessages
}

// ValidatorMessages 验证器自定义错误信息字典
type ValidatorMessages map[string]string

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

var V = New()

func New() *Validator {
	v := &Validator{
		validate: validator.New(),
	}
	cn := zh.New()
	uni := ut.New(cn, cn)
	v.trans, _ = uni.GetTranslator("zh")
	// 注册翻译器
	err := zh_translations.RegisterDefaultTranslations(v.validate, v.trans)
	if err != nil {
		log.Printf("valid init error : %v \n", err.Error())
	}
	return v
}

func (v *Validator) Registry(tag string, fn func(fl validator.FieldLevel) bool, callValidationEvenIfNull ...bool) *Validator {
	v.validate.RegisterValidation(tag, fn, callValidationEvenIfNull...)
	return v
}

func (v *Validator) Struct(data ValidatorInterface) string {
	err := v.validate.Struct(data)
	if err == nil {
		return ""
	}
	return v.getErrorMsg("struct", data.GetValidMessage(), err)
}

func (v *Validator) Var(data any, tag string, request ValidatorMessages, key ...string) string {

	err := v.validate.Var(data, tag)
	if err == nil {
		return ""
	}

	return v.getErrorMsg("var", request, err, key...)
}

func (v *Validator) Map(data map[string]interface{}, rule map[string]interface{}, request ValidatorMessages) string {
	var msg string
	for k, e := range v.validate.ValidateMap(data, rule) {
		msg = v.getErrorMsg("map", request, e.(validator.ValidationErrors), k)
	}
	return msg
}

func (v *Validator) getErrorMsg(validType string, request ValidatorMessages, err error, keys ...string) string {
	var key string
	var field string
	for _, e := range err.(validator.ValidationErrors) {
		switch validType {
		case "struct":
			field = e.Field()
			key = e.StructNamespace() + "." + e.Tag()
		case "map":
			if len(keys) > 0 {
				field = keys[0]
			}
			key = field + "." + e.Tag()
		case "var":
			key = e.Tag()
			if len(keys) > 0 {
				field = keys[0]
			}
		}

		if message, exist := request[key]; exist {
			return message
		}

		if e.Field() != "" {
			return e.Translate(v.trans)
		}

		return field + e.Translate(v.trans)
	}
	return ""
}
