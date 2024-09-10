package document

import (
	"code.sajari.com/docconv"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
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

	//dir := filepath.Dir(req.FilePath)
	//wg := sync.WaitGroup{}
	//go func() {
	//	wg.Add(1)
	//	err := global.Pdfbox.ExtractImages(dir, req.FileName, "images")
	//	if err != nil {
	//		global.Log.Error(err, "提取图片失败")
	//		wg.Done()
	//		return
	//	}
	//	imageList := make([]string, 0)
	//	err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
	//		if strings.Index(info.Name(), "images-") != -1 {
	//			imageList = append(imageList, info.Name())
	//		}
	//		return nil
	//	})
	//	if err != nil {
	//		global.Log.Error(err, "获取图片列表失败")
	//		wg.Done()
	//		return
	//	}
	//
	//	if len(imageList) == 0 {
	//		global.Log.Error(errors.New("图片列表为空"), "检查imageList")
	//		wg.Done()
	//		return
	//	}
	//
	//	imageList = d.sortFilenames(imageList)
	//	for i, image := range imageList {
	//		err = os.Rename(filepath.Join(dir, image), filepath.Join(dir, "images-"+strconv.Itoa(i)+".jpg"))
	//		if err != nil {
	//			global.Log.Error(err, "重命名图片失败")
	//			continue
	//		}
	//	}
	//	wg.Done()
	//}()
	//
	//html, err := global.Tika.PdfToHtml(req.FilePath)
	//if err != nil {
	//	global.Log.Error(err, "PDF转换HTML失败")
	//	return nil, err
	//}
	//// 正则表达式匹配 embedded:image-01.jpg 或 embedded:image-01.png
	//re := regexp.MustCompile(`embedded:image(\d{1,10})\.(jpg|png|jpeg)`)
	//
	//// 替换为 image-01.jpg 保留编号
	//htmlStr := re.ReplaceAllString(string(html), `images-$1.jpg`)
	//
	////re1 := regexp.MustCompile(`image(\d{1,10})\.(jpg|png|jpeg)`)
	////
	////htmlStr = re1.ReplaceAllString(htmlStr, `images-$1.jpg`)
	//
	//err = os.WriteFile(filepath.Join(dir, fmt.Sprintf("%v.html", req.FileId)), []byte(htmlStr), os.ModePerm)
	//if err != nil {
	//	global.Log.Error(err, "写入HTML失败")
	//	return nil, err
	//}
	//wg.Wait()
	//
	//err = global.Pandoc.HtmlToWord(dir, fmt.Sprintf("%v.html", req.FileId), fmt.Sprintf("%v.docx", req.FileId))
	//if err != nil {
	//	global.Log.Error(err, "HTML转换WORD失败")
	//	return nil, err
	//}

	dir := filepath.Dir(req.FilePath)
	outputFile := fmt.Sprintf("%v.docx", req.FileId)
	err := global.Pdf2docx.PdfToWord(dir, req.FileName, outputFile)
	if err != nil {
		global.Log.Error(err, "PDF转换WORD失败")
		return nil, err
	}
	return &types.ConvertResponse{
		Filename: filepath.Join(dir, outputFile),
	}, nil
}

func (d *DocumentService) sortFilenames(filenames []string) []string {
	re := regexp.MustCompile(`\d+`) // 匹配文件名中的数字部分
	sort.Slice(filenames, func(i, j int) bool {
		// 提取数字部分
		numI := re.FindString(filenames[i])
		numJ := re.FindString(filenames[j])

		// 转换为整数进行比较
		numIInt, _ := strconv.Atoi(numI)
		numJInt, _ := strconv.Atoi(numJ)

		// 如果数字相同，则按字母顺序比较
		if numIInt == numJInt {
			return filenames[i] < filenames[j]
		}

		return numIInt < numJInt
	})
	return filenames
}
