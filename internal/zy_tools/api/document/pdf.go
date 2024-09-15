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

package document

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"path"
	"path/filepath"
	"zy-tools/internal/zy_tools/global"
	"zy-tools/internal/zy_tools/models/document"
	"zy-tools/internal/zy_tools/utils"
	"zy-tools/pkg/common/response"
)

// PdfToExcel Pdf 转excel
func (d *DocumentApi) PdfToExcel() {

}

// PdfToText pdf转文字
func (d *DocumentApi) PdfToText(c *gin.Context) {
	//file, err := c.FormFile("file")
	//if err != nil {
	//	response.R.Custom(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//filename := file.Filename
	//dst := path.Join(global.Config.Server.UploadPath, uuid.New().String())
	//err = c.SaveUploadedFile(file, dst)
	//if err != nil {
	//	response.R.Custom(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//text, err := documentService.PDFToText(dst)
	//if err != nil {
	//	return
	//}
}

func (d *DocumentApi) PdfToword(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		global.Log.Error(err, "获取form文件")
		response.R.ErrorWithMessage(c, err.Error())
		return
	}

	f, err := file.Open()
	if err != nil {
		global.Log.Error(err, "文件打开失败")
		response.R.ErrorWithMessage(c, "文件类型必须是PDF格式")
	}
	if !utils.CheckFileBufferMimeByExt(f, "pdf") {
		global.Log.Error(err, "文件类型必须是PDF格式")
		response.R.ErrorWithMessage(c, "文件类型必须是PDF格式")
		return
	}

	fileName := file.Filename
	ext := filepath.Ext(fileName)
	fileId := uuid.New().String()

	dir := path.Join(global.Config.Server.UploadPath, fileId)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		global.Log.Error(err, "执行转换失败")
		response.R.ErrorWithMessage(c, err.Error())
		return
	}
	dstFileName := fmt.Sprintf("%v%v", fileId, ext)
	dst := path.Join(dir, dstFileName)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		global.Log.Error(err, "保存文件失败")
		response.R.ErrorWithMessage(c, err.Error())
		return
	}

	result, err := documentService.PdfToWord(document.ConvertRequest{
		FileId:   fileId,
		FilePath: dst,
	})
	if err != nil {
		global.Log.Error(err, "执行转换失败")
		response.R.ErrorWithMessage(c, err.Error())
		return
	}

	response.R.SuccessWithData(c, gin.H{
		"filePath": result.Filename,
		"fileName": fileName,
	})
}
