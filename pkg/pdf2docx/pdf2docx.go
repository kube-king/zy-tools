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

package pdf2docx

import (
	"os"
	"zy-tools/pkg/command"
)

type Pdf2Docx struct {
	Command string
}

func NewPdf2Docx(command string) *Pdf2Docx {
	return &Pdf2Docx{
		Command: command,
	}
}

func (p *Pdf2Docx) PdfToWord(dir, filename string, outputPath string) error {
	e := command.NewExec(p.Command, "convert", filename, outputPath)
	e.SetTimeout(0)
	e.SetDir(dir)
	e.SetStderr(os.Stderr)
	e.SetStdout(os.Stdout)
	_, err := e.Run()
	if err != nil {
		return err
	}
	return nil
}
