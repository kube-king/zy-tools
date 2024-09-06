package main

import (
	"bytes"
	"fmt"
	_ "github.com/go-resty/resty/v2"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	//f, err := os.Open("./data/1.pdf")
	//if err != nil {
	//	return
	//}

	//bytes, err := os.ReadFile("./data/1.pdf")
	//if err != nil {
	//	return
	//}
	//
	//c := resty.New()
	//response, err := c.R().SetHeader("Accept", "image/*").SetBody(bytes).Put("http://127.0.0.1:9998/tika")
	//fmt.Println("status code", response.StatusCode())
	//fmt.Println("body:", string(response.Body()))
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	//url := "http://127.0.0.1:9998/tika"
	//method := "PUT"
	//
	//file, err := os.ReadFile("./data/1.pdf")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//payload := strings.NewReader(string(file))
	//
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, payload)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("Accept", "image/*")
	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	//req.Header.Add("Host", "127.0.0.1:9998")
	//req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Content-Type", "application/pdf")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))

	data, err := os.ReadFile("./data/1.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	body := bytes.NewReader(data)
	request, err := http.NewRequest("PUT", "http://127.0.0.1:9998/tika", body)
	if err != nil {
		return
	}

	request.Header.Set("Connection", "Keep-Alive")

	//hdr["X-Tika-PDFextractInlineImages"] = []string{"true"}
	request.Header.Set("Accept", "text/html")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	fmt.Println(string(b))
}
