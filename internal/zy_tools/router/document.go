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

package router

import (
	"github.com/gin-gonic/gin"
	"zy-tools/internal/zy_tools/api"
)

// InitDocumentRouter 初始化Document路由
func InitDocumentRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	document := Router.Group("document")

	var documentApi = api.ApiGroupApp.DocumentApi
	{
		pdf := document.Group("pdf")
		{
			//pdf.POST("to-text", documentApi.PDFToText)
			//pdf.POST("to-image", documentApi.PDFToImage)
			pdf.POST("to-word", documentApi.PdfToword)
		}

		ppt := document.Group("ppt")
		{
			//ppt.POST("to-text", documentApi.PDFToText)
			ppt.POST("to-pdf", documentApi.PptToPdf)
		}

		image := document.Group("image")
		{
			//image.POST("to-pdf", documentApi.ImageToPDF)
			//image.POST("to-text", documentApi.ImageToText)
			//image.POST("to-excel,)
			image.POST("to-ppt", documentApi.ImageToPPT)

		}
		//
		excel := document.Group("excel")
		{
			excel.POST("to-pdf", documentApi.ExcelToPdf)
		}

		//
		word := document.Group("word")
		{
			//word.POST("to-image", documentApi.WordToImage)
			word.POST("to-pdf", documentApi.WordToPdf)
		}

		json := document.Group("json")
		{
			json.POST("to-excel", documentApi.JsonToExcel)
		}
	}
}
