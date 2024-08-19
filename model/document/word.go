package document

import (
	"errors"
	"github.com/wonderivan/logger"
	"os"
	"zy-tools/utils"
	"zy-tools/utils/ocr"
)

func ImageToWord(imagePath string) (outputFile string, err error) {
	logger.Debug("图片文件路径:" + imagePath)
	_, err = os.Stat(imagePath)
	if err != nil {
		return "", errors.New("图片文件不存在")
	}
	text := ocr.ImageToText(imagePath)
	logger.Debug("ocr识别结果:" + text)
	tmpFile := "/tmp/" + utils.GetId() + ".txt"
	file, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_WRONLY, 0754)
	defer file.Close()
	defer os.Remove(tmpFile)
	logger.Debug("临时文件:" + tmpFile)
	if err != nil {
		return
	}
	file.WriteString(text)
	outputFile, err = utils.ConvertHeadless(tmpFile, "docx")
	if err != nil {
		return
	}
	return
}
