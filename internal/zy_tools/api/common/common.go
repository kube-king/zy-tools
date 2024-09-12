package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"zy-tools/internal/zy_tools/global"
	"zy-tools/pkg/common/response"
)

type CommonApi struct {
}

func (ca *CommonApi) Download(c *gin.Context) {
	fileName := c.Query("fileName")
	path := c.Query("path")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")

	fmt.Println("文件名", fileName)
	filePath := filepath.Join(global.Config.Server.UploadPath, path)
	fmt.Println("文件路径", filePath)
	if _, err := os.Stat(filePath); err != nil {
		response.R.Custom(c, http.StatusNotFound, "文件不存在!")
		return
	}
	c.File(filePath)
	return
}
