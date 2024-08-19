package document_test

import (
	"testing"
	"zy-tools/model/document"
)

// go test -v -run TestJsonToExcel model/document/excel_test.go
func TestJsonToExcel(t *testing.T) {
	type args struct {
		jsonText string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "测试json转换excel",
			args: args{jsonText: `[{"_id":"13001009656","name":"test1"},{"_id":"13001022163","name":"test2","sex":"男"},{"_id":"13001086801","name":"test3","user":{"userName":"test3","workUnit":"清华"}}]`},
		},
		{
			name: "测试json转换excel1",
			args: args{jsonText: `{"json":"[{\"_id\":\"13001009656\",\"name\":\"test1\"},{\"_id\":\"13001022163\",\"name\":\"test2\",\"sex\":\"男\"},{\"_id\":\"13001086801\",\"name\":\"test3\",\"user\":{\"userName\":\"test3\",\"workUnit\":\"清华\"}}]"}`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document.JsonToExcel(tt.args.jsonText)
		})
	}
}
