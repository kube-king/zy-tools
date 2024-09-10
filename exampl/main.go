package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 原始字符串
	text := `
	embedded:image265.jpg
	embedded:image01.jpg
	embedded:image-02.png
	embedded:image-03.jpg
	embedded:image-04.png
	`

	// 正则表达式匹配 embedded:image-01.jpg 或 embedded:image-01.png
	re := regexp.MustCompile(`embedded:image(\d{1,10})\.(jpg|png)`)

	// 替换为 image-01.jpg 保留编号
	result := re.ReplaceAllString(text, `images-$1.jpg`)

	fmt.Println(result)
}
