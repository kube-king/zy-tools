/*
Copyright 2024 Faw Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tika

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Tika struct {
	Host string
}

func NewTika(host string) *Tika {
	return &Tika{
		Host: host,
	}
}

func (t *Tika) PdfToHtml(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(data)
	request, err := http.NewRequest("PUT", fmt.Sprintf("%v%v", t.Host, "/tika"), body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Connection", "Keep-Alive")
	request.Header.Set("Accept", "text/html")
	request.Header.Set("X-Tika-PDFextractInlineImages", "true")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
