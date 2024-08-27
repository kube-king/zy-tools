package ocr_test

import (
	"testing"
	"zy-tools/pkg/ocr"
)

// go test -v
func TestImageToText(t *testing.T) {
	type args struct {
		imgPath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "测试ocr图片转文字",
			args: args{
				imgPath: "/opt/code/zy-tools/table.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ocr.ImageToText(tt.args.imgPath)
		})
	}
}
