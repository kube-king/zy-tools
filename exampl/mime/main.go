package main

import (
	"fmt"
	"zy-tools/internal/zy_tools/utils"
)

func main() {

	//mime := utils.GetFileMime("/Users/xiangjiqiang/code/my/zy-tools/uploads/0bbfc083-7827-4713-b2d8-7c0bb6525b6e.docx")

	ok := utils.CheckMimeByExt("/Users/xiangjiqiang/code/my/zy-tools/uploads/0bbfc083-7827-4713-b2d8-7c0bb6525b6e.docx", "zip")

	fmt.Println(ok)
}
