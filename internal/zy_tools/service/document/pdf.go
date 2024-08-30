package document

import (
	"code.sajari.com/docconv"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/unidoc/unipdf/v3/extractor"
	"os"
	"zy-tools/internal/zy_tools/global"
	"zy-tools/internal/zy_tools/models/document"
	"zy-tools/pkg/doc_conv/types"
)

// PDFToText PDF 转文本
func (d *DocumentService) PDFToText(docFile string) (string, error) {

	fileHandle, err := os.Open(docFile)
	if err != nil {
		global.Log.Error(err, "PDF转换文本失败")
		return "", err
	}

	doc, _, err := docconv.ConvertPDF(fileHandle)
	if err != nil {
		global.Log.Error(err, "PDF转换文本失败")
		return "", err
	}

	return doc, nil
}

func (d *DocumentService) PdfToWord(req document.ConvertRequest) (*types.ConvertResponse, error) {
	//convertResponse, err := global.Office.Convert(req.FilePath, constants.FileTypeDocx)
	//if err != nil {
	//	return nil, err
	//}
	//content, err := pdf_api.ExtractF("input.pdf", nil, nil)
	//if err != nil {
	//	global.Log.Error(err, "PDF转换文本失败")
	//	return nil, err
	//}
	//
	//// 创建一个新的 Word 文档
	//doc := udoc.New()
	//
	//// 将提取的文本添加到 Word 文档中
	//para := doc.AddParagraph()
	//para.AddRun().AddText(content)
	//
	//// 保存文档
	//file, err := os.Create("output.docx")
	//if err != nil {
	//	log.Fatalf("Error creating output file: %v", err)
	//}
	//defer file.Close()
	//
	//doc.Save(file)
	//log.Println("Document created successfully")

	// Read PDF document
	pdfFile := "path/to/input.pdf"
	pdfReader, err := extractor.NewPdfReader(pdfFile)
	if err != nil {
		// Handle error
	}
	defer pdfReader.Close()

	// Create Word document
	wordDoc := unioffice.NewDocument()

	// Convert PDF document content to Word document content
	pages, err := pdfReader.GetPages()
	if err != nil {
		// Handle error
	}

	for _, page := range pages {
		text, err := page.GetText()
		if err != nil {
			// Handle error
		}

		paragraph := wordDoc.AddParagraph()
		paragraph.AddRun().AddText(text)
	}

	// Save Word document
	wordFile := "path/to/output.docx"
	err = wordDoc.SaveToFile(wordFile)
	if err != nil {
		// Handle error
	}

	return convertResponse, nil
}
