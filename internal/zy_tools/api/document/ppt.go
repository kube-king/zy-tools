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
	"path"
	"path/filepath"
	"zy-tools/internal/zy_tools/global"
	"zy-tools/internal/zy_tools/models/document"
	"zy-tools/pkg/common/response"
)

func (d *DocumentApi) PptToPdf(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		global.Log.Error(err, "获取form文件")
		response.R.ErrorWithMessage(c, err.Error())
		return
	}

	fileName := file.Filename
	ext := filepath.Ext(fileName)
	dst := path.Join(global.Config.Server.UploadPath, fmt.Sprintf("%v%v", uuid.New().String(), ext))
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		global.Log.Error(err, "保存文件失败")
		response.R.ErrorWithMessage(c, err.Error())
		return
	}

	result, err := documentService.PptToPdf(document.ConvertRequest{
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
