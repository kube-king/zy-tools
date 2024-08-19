package document

import (
	"code.sajari.com/docconv"
	"fmt"
	"os"
)

// PDFToText PDF 转文本
func PDFToText(docFile string) (string, error) {

	fileHandle, err := os.Open(docFile)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	doc, _, err := docconv.ConvertPDF(fileHandle)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return doc, nil
}
