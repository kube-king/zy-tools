package main

import (
	"zy-tools/internal/zy_tools/models/document"
	"zy-tools/internal/zy_tools/service"
)

func main() {
	//global.Office = libreoffice.NewLibreOffice(
	//	libreoffice.WithCommand("global.Config.LibreOffice.Command"),
	//	libreoffice.WithOutputPath("global.Config.Server.UploadPath"),
	//)
	//
	////fmt.Println("1")
	service.GroupAppService.ImageToPpt(document.ConvertRequest{})

}
