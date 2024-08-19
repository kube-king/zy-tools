package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zy-tools/config"
	"zy-tools/model/document"
	"zy-tools/utils"
)

func FileConvert(c *gin.Context) {
	srcType, _ := c.GetPostForm("srcType")
	dstType, _ := c.GetPostForm("dstType")

	file, err := c.FormFile("file")
	if err != nil {
		utils.R.Custom(c, http.StatusInternalServerError, err.Error())
		return
	}

	dst := config.UploadPath + file.Filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		utils.R.Custom(c, http.StatusInternalServerError, err.Error())
		return
	}

	outputFile, err := document.FileConverter(dst, srcType, dstType)
	if err != nil {
		utils.R.Custom(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.R.Success(c, outputFile)
}
