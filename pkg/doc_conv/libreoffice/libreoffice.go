package libreoffice

import (
	"log"
	"os"
	"os/exec"
	"path"
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

	cmd := exec.Command(l.Command, "--invisible", "--language=zh-CN", "--convert-to", fileType, srcFile, "--outdir", l.OutputPath)
	var output []byte
	output, err = cmd.Output()
	log.Println("转换结果:" + string(output))
	if err != nil {
		return nil, err
	}

	return &types.ConvertResponse{
		Content:  string(output),
		FilePath: path.Join(l.OutputPath, srcFileHandle.Name(), fileType),
	}, nil
}

//
//// ConvertHeadless 文档转换
//func ConvertHeadless(srcFile string, fileType string) (result string, err error) {
//	var srcFileHandle, outPutPath *os.File
//	srcFileHandle, err = os.Open(srcFile)
//	if err != nil && os.IsNotExist(err) {
//		return "", err
//	}
//
//	outPutPath, err = os.Open(constants.FileOutPutPath)
//	if err != nil && os.IsNotExist(err) {
//		err = os.MkdirAll(constants.FileOutPutPath, os.ModePerm)
//		if err != nil {
//			return "", err
//		}
//	}
//
//	defer func() {
//		_ = srcFileHandle.Close()
//		_ = outPutPath.Close()
//	}()
//
//	args := make([]string, 0)
//	args = append(args, "--headless")
//	args = append(args, "--convert-to")
//	args = append(args, fileType)
//	args = append(args, srcFile)
//	args = append(args, "--outdir")
//	args = append(args, constants.FileOutPutPath)
//	cmd := exec.Command(CommandLibreoffice, args...)
//
//	var output []byte
//	output, err = cmd.Output()
//	fmt.Println("转换结果:" + string(output))
//	if err != nil {
//		return "", errors.New(err.Error())
//	}
//	result = strings.Split(path.Base(srcFile), ".")[0] + "." + fileType
//	return
//}
