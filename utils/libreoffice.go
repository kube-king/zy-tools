package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"zy-tools/config"
)

const (
	CommandLibreoffice = "/usr/bin/libreoffice7.4"
)

type OfficeOption struct {
}

// Convert 文档转换
func Convert(srcFile string, fileType string) (result string, err error) {
	var srcFileHandle, outPutPath *os.File
	srcFileHandle, err = os.Open(srcFile)
	if err != nil && os.IsNotExist(err) {
		return "", err
	}

	outPutPath, err = os.Open(config.FileOutPutPath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(config.FileOutPutPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	defer func() {
		_ = srcFileHandle.Close()
		_ = outPutPath.Close()
	}()

	cmd := exec.Command(CommandLibreoffice, "--invisible", "--language=zh-CN", "--convert-to", fileType, srcFile, "--outdir", config.FileOutPutPath)
	var output []byte
	output, err = cmd.Output()
	fmt.Println("转换结果:" + string(output))
	if err != nil {
		return "", errors.New(err.Error())
	}
	result = strings.Split(path.Base(srcFile), ".")[0] + "." + fileType
	return
}

// ConvertHeadless 文档转换
func ConvertHeadless(srcFile string, fileType string) (result string, err error) {
	var srcFileHandle, outPutPath *os.File
	srcFileHandle, err = os.Open(srcFile)
	if err != nil && os.IsNotExist(err) {
		return "", err
	}

	outPutPath, err = os.Open(config.FileOutPutPath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(config.FileOutPutPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	defer func() {
		_ = srcFileHandle.Close()
		_ = outPutPath.Close()
	}()

	args := make([]string, 0)
	args = append(args, "--headless")
	args = append(args, "--convert-to")
	args = append(args, fileType)
	args = append(args, srcFile)
	args = append(args, "--outdir")
	args = append(args, config.FileOutPutPath)
	cmd := exec.Command(CommandLibreoffice, args...)

	var output []byte
	output, err = cmd.Output()
	fmt.Println("转换结果:" + string(output))
	if err != nil {
		return "", errors.New(err.Error())
	}
	result = strings.Split(path.Base(srcFile), ".")[0] + "." + fileType
	return
}
