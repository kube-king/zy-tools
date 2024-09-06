package main

import (
	"context"
	"fmt"
	"github.com/google/go-tika/tika"
	"net/http"
	"os"
)

// 解析PDF文件
func ReadPdf(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	client := tika.NewClient(nil, "http://127.0.0.1:9998/tika")
	hdr := http.Header{}
	//hdr["Content-Type"] = []string{"application/pdf"}
	//hdr["X-Tika-OCRLanguage"] = []string{"eng"}
	hdr["X-Tika-PDFextractInlineImages"] = []string{"true"}
	hdr["Accept"] = []string{"text/html"}

	_, err = client.ParseWithHeader(context.Background(), nil, hdr)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return client.Parse(context.TODO(), f)
}
func main() {

	res, err := ReadPdf("./data/1.pdf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)

}
