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

package constants

const (
	RouterPrefix = "/api"
)

const (
	DefaultConfigFileName     = "config"
	ResourceServiceConfigPath = "./internal/zy_tools/conf"
)

// gin 配置
var (
	ListenAddr = ":8080"
	GinMode    = "debug"
)

// 文件处理全局配置
const (
	HomePath         = "/opt/code/zy-tools"
	FileOutPutPath   = HomePath + "/output"
	OutPutStaticPath = "/output"
)

// FileType 文件类型
const (
	FileTypePdf   = "pdf"
	FileTypeDoc   = "doc"
	FileTypeDocx  = "docx"
	FileTypeExcel = "xlsx"
	FileTypePpt   = "ppt"
	FileTypeJson  = "json"
	FileTypeText  = "text"
	FileTypeImage = "image"
)

// 上传配置
const (
	UploadMaxSize = 10 << 20 // 上传限制10Mb
	UploadPath    = "./upload/"
)
