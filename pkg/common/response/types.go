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

// ResponseBody API返回结构体
type ResponseBody struct {
	Code    int         `json:"code,omitempty"`    // 错误码
	Message string      `json:"message,omitempty"` // 返回失败消息
	Data    interface{} `json:"data"`              // 请求返回数据
}

// PageResult 分页返回结构体
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
