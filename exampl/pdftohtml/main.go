package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	//srcFile := "data/1.pdf"
	//args := make([]string, 0)
	//args = append(args, "-c")
	//args = append(args, srcFile)
	//cmd := exec.Command("pdftohtml", args...)
	//var output []byte
	//var err error
	//output, err = cmd.Output()
	//log.Println("转换结果:" + string(output))
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	srcFile := "./data/1.html"
	args := make([]string, 0)
	args = append(args, "--headless")
	//args = append(args, "--invisible")
	args = append(args, "--convert-to")
	args = append(args, "docx")
	args = append(args, srcFile)
	args = append(args, "--outdir")
	args = append(args, "./data/")
	log.Println("soffice", strings.Join(args, " "))
	cmd := exec.Command("soffice", args...)
	var output []byte
	var err error
	output, err = cmd.Output()
	log.Println("转换结果:" + string(output))
	if err != nil {
		fmt.Println(err.Error())
	}

}
