package document

type DocumentApi struct {
}

//func (d *DocumentApi) FileConvert(c *gin.Context) {
//	srcType, _ := c.GetPostForm("srcType")
//	dstType, _ := c.GetPostForm("dstType")
//
//	file, err := c.FormFile("file")
//	if err != nil {
//		response.R.Custom(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	dst := constants.UploadPath + file.Filename
//	err = c.SaveUploadedFile(file, dst)
//	if err != nil {
//		response.R.Custom(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	outputFile, err := documentService.FileConverter(dst, srcType, dstType)
//	if err != nil {
//		response.R.Custom(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	response.R.SuccessWithData(c, outputFile)
//}
