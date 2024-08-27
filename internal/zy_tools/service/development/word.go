package development

//
//func (d *DocumentService) ImageToWord(imagePath string) (outputFile string, err error) {
//	_, err = os.Stat(imagePath)
//	if err != nil {
//		return "", errors.New("图片文件不存在")
//	}
//	text := ocr.ImageToText(imagePath)
//	tmpFile := "/tmp/" + utils.GetId() + ".txt"
//	file, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_WRONLY, 0754)
//	defer file.Close()
//	defer os.Remove(tmpFile)
//	if err != nil {
//		return
//	}
//	file.WriteString(text)
//	convertResponse, err := global.Office.Convert(tmpFile, "docx")
//	if err != nil {
//		return "", err
//	}
//	return convertResponse.FilePath, nil
//}
