package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"zy-tools/internal/zy_tools/constants"
	"zy-tools/pkg/common/response"
)

type CommonApi struct {
}

func (ca *CommonApi) Download(c *gin.Context) {
	fileName := c.Query("fileName")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")

	filePath := constants.FileOutPutPath + "/" + fileName
	if _, err := os.Stat(filePath); err != nil {
		response.R.Custom(c, http.StatusNotFound, "资源不存在!")
		return
	}
	c.File(filePath)
	return
}
