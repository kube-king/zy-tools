package development_test

import (
	"testing"
	"zy-tools/model/development"
)

//func TestCheckJson(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			development.CheckJson()
//		})
//	}
//}

func TestSqlToStruct(t *testing.T) {
	type args struct {
		sql string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name:       "测试sql 转结构体",
			args:       args{sql: "CREATE          TABLE \"t_tests\""},
			wantResult: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := development.SqlToStruct(tt.args.sql); gotResult == tt.wantResult {
				t.Errorf("SqlToStruct() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
