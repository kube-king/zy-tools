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
	"github.com/gin-gonic/gin"
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
