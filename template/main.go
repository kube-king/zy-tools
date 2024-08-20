package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"os"
)

const (
	OutPutPath    = "./gen/output/page/"
	TemplatePath  = "gen/static/page/document/*.html"
	IndexHtmlPath = "./gen/output/page/index.html"
)

var (
	templateNameList = []string{
		"pdf-to-word",
		"word-to-pdf",
		"ppt-to-pdf",
		"json-to-excel",
		"image-to-word",
		"image-to-pdf",
		"image-to-excel",
		"excel-to-pdf",
		"pdf-to-text",
	}
)

func createTemplate(name string) {
	fileHandle, err := os.OpenFile(fmt.Sprintf("%v%v.html", OutPutPath, name), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t := template.Must(template.ParseGlob(TemplatePath))
	err = t.ExecuteTemplate(fileHandle, name, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateTemplateAll(name ...string) {
	for _, n := range name {
		createTemplate(n)
	}
}

func CopyFile(src string, dest string) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//defer srcFile.Close()
	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//defer destFile.Close()
	// 创建 Reader
	r := bufio.NewReader(srcFile)
	// 每次读取 1024 个字节
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
		}
		if n == 0 {
			break
		}
		destFile.WriteString(string(buf[:n]))
	}
}

func main() {
	CreateTemplateAll(templateNameList...)
	CopyFile(fmt.Sprintf("%v%v.html", OutPutPath, "word-to-pdf"), IndexHtmlPath)
}
