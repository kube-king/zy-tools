package document

import (
	"github.com/gin-gonic/gin"
	document_model "zy-tools/internal/zy_tools/models/document"
	"zy-tools/internal/zy_tools/service"
	"zy-tools/pkg/common/request"
	"zy-tools/pkg/common/response"
)

var documentService = service.GroupAppService.DocumentService

func (d *DocumentApi) JsonToExcel(c *gin.Context) {

	jsonToExcel := &document_model.JsonToExcel{}
	if !request.ShouldBindJson(c, jsonToExcel) {
		return
	}

	excel, err := documentService.JsonToExcel(jsonToExcel.Json)
	if err != nil {
		response.R.ErrorWithMessage(c, err.Error())
		return
	}

	response.R.SuccessWithData(c, excel)
	return
}
