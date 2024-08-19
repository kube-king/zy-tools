package document_test

import (
	"testing"
	"zy-tools/config"
	"zy-tools/model/document"
)

// go test -v model/document/word_test.go
func TestImageToWord(t *testing.T) {
	type args struct {
		imagePath string
	}
	tests := []struct {
		name           string
		args           args
		wantOutputFile string
		wantErr        bool
	}{
		{
			name: "图片转word",
			args: args{
				imagePath: config.HomePath + "/example/1.jpg",
			},
			wantOutputFile: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutputFile, err := document.ImageToWord(tt.args.imagePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageToWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutputFile == tt.wantOutputFile {
				t.Errorf("ImageToWord() gotOutputFile = %v, want %v", gotOutputFile, tt.wantOutputFile)
			}
		})
	}
}
