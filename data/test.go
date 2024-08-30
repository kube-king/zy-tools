package main

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	"os"
)

func main() {
	// PDF文件路径
	pdfFile := "./data/1.pdf"

	// Word文档路径
	//wordFile := "path/to/word_file.docx"

	// 读取PDF文件
	f, err := os.Open(pdfFile)
	if err != nil {
		fmt.Println("Error opening PDF file:", err)
		return
	}
	defer f.Close()

	reader, err := model.NewPdfReader(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pages, err := reader.GetNumPages()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Number of pages:", pages)

	for i := 0; i < pages; i++ {

		pageNum := i + 1
		page, err := reader.GetPage(pageNum)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		ex, err := extractor.New(page)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		text, err := ex.ExtractText()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(text)

		//// 创建Word文档
		//w, err := os.Create(wordFile)
		//if err != nil {
		//	fmt.Println("Error creating Word file:", err)
		//	return
		//}
		//defer w.Close()

		// 写入Word文档
	}

	//
	//// 创建PDF解析器
	//r, err := extractor.New()
	//if err != nil {
	//	fmt.Println("Error creating PDF parser:", err)
	//	return
	//}
	//
	//// 解析PDF文件
	//text, err := r.Text()
	//if err != nil {
	//	fmt.Println("Error parsing PDF file:", err)
	//	return
	//}
	//
	//// 创建Word文档
	//w, err := os.Create(wordFile)
	//if err != nil {
	//	fmt.Println("Error creating Word file:", err)
	//	return
	//}
	//defer w.Close()
	//
	//// 写入Word文档
	//w.Write([]byte(text))
	//
	//fmt.Println("PDF file converted to Word file successfully.")
}
