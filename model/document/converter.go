package document

import (
	"zy-tools/config"
	"zy-tools/utils"
)

// FileConverter 文件转换器
func FileConverter(srcFile string, srcType string, dstType string) (result string, err error) {
	switch srcType {
	case config.FileTypePdf:
		switch dstType {
		case config.FileTypeText:
			result, err = PDFToText(srcFile)
			return
		}
	case config.FileTypeImage:
		switch dstType {
		case config.FileTypeDocx:
			result, err = ImageToWord(srcFile)
			return
		}

	case config.FileTypeJson:
		switch dstType {
		case config.FileTypeExcel:

			return
		}
	}

	result, err = utils.Convert(srcFile, dstType)
	return
}
