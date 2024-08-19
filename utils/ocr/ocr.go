package ocr

import (
	"github.com/otiai10/gosseract/v2"
)

func ImageToText(imgPath string) string {
	client := gosseract.NewClient()
	client.SetLanguage("chi", "eng", "chi_sim")
	defer client.Close()
	client.SetImage(imgPath)
	text, _ := client.Text()
	return text
}
