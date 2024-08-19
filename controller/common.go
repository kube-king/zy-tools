package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"zy-tools/config"
	"zy-tools/utils"
)

func Download(c *gin.Context) {
	fileName := c.Query("fileName")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")

	filePath := config.FileOutPutPath + "/" + fileName
	if _, err := os.Stat(filePath); err != nil {
		utils.R.Custom(c, http.StatusNotFound, "资源不存在!")
		return
	}
	c.File(filePath)
	return
}
