package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//global.Office = libreoffice.NewLibreOffice(
	//	libreoffice.WithCommand("global.Config.LibreOffice.Command"),
	//	libreoffice.WithOutputPath("global.Config.Server.UploadPath"),
	//)
	//
	////fmt.Println("1")

	filename := "uploads/efabb9a2-c143-4a6d-a4ad-3798cd695353.docx"

	ext := filepath.Ext(filename)
	base := filepath.Base(filename)
	abs, err := filepath.Abs(filename)
	if err != nil {
		return
	}

	filepath.Split(abs)

	fmt.Println(abs)
	fmt.Println(ext, base)
}
