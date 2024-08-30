package libreoffice

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"zy-tools/pkg/doc_conv/types"
)

const (
	CommandLibreoffice = "/usr/bin/libreoffice7.4"
)

type OfficeOption struct {
}

type LibreOffice struct {
	Command    string
	OutputPath string
}

type OpsFunc func(*LibreOffice)

func NewLibreOffice(opsFuncs ...OpsFunc) *LibreOffice {
	l := &LibreOffice{}
	for _, opsFunc := range opsFuncs {
		opsFunc(l)
	}
	return l
}

func WithCommand(command string) OpsFunc {
	return func(l *LibreOffice) {
		l.Command = command
	}
}

func WithOutputPath(outputPath string) OpsFunc {
	return func(l *LibreOffice) {
		l.OutputPath = outputPath
	}
}

// Convert 文档转换
func (l *LibreOffice) Convert(srcFile string, fileType string) (response *types.ConvertResponse, err error) {
	var srcFileHandle, outPutPath *os.File
	srcFileHandle, err = os.Open(srcFile)
	if err != nil && os.IsNotExist(err) {
		return nil, err
	}

	outPutPath, err = os.Open(l.OutputPath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(l.OutputPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	defer func() {
		_ = srcFileHandle.Close()
		_ = outPutPath.Close()
	}()

	args := make([]string, 0)
	args = append(args, "--headless")
	args = append(args, "--invisible")
	args = append(args, "--convert-to")
	args = append(args, fileType)
	args = append(args, srcFile)
	args = append(args, "--outdir")
	args = append(args, l.OutputPath)
	log.Println(l.Command, strings.Join(args, " "))
	cmd := exec.Command(l.Command, args...)
	var output []byte
	output, err = cmd.Output()
	log.Println("转换结果:" + string(output))
	if err != nil {
		return nil, err
	}
	filename := path.Base(srcFileHandle.Name())
	nameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
	return &types.ConvertResponse{
		Content:  string(output),
		Filename: fmt.Sprintf("%v.%v", nameWithoutExt, fileType),
	}, nil
}
