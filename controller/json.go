package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zy-tools/model/document"
	"zy-tools/utils"
)

func JsonToExcel(c *gin.Context) {

	params := &struct {
		Json string `json:"json" validate:"json"`
	}{}

	if !utils.BindJson(c, params) {
		return
	}

	excel, err := document.JsonToExcel(params.Json)
	if err != nil {
		utils.R.Custom(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.R.Success(c, excel)
	return
}
