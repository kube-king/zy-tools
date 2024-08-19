package config

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
