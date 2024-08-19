package document

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/wonderivan/logger"
	"github.com/xuri/excelize/v2"
	_ "github.com/xuri/excelize/v2"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"reflect"
	"zy-tools/config"
	"zy-tools/utils"
)

// JsonToExcel json 转 Excel
func JsonToExcel(jsonText string) (result string, err error) {
	tmpMapList := make([]map[string]interface{}, 0)
	err = json.Unmarshal([]byte(jsonText), &tmpMapList)
	if err != nil {
		return
	}
	f := excelize.NewFile()
	s := utils.Set{}
	for i, tmpMap := range tmpMapList {
		j := 0
		flag := 65
		for key, v := range tmpMap {
			ty := reflect.TypeOf(v)
			s.Add(key)
			cellFlag := fmt.Sprintf("%c", flag+j)
			cellNum := cellFlag + fmt.Sprintf("%v", i+2)
			if ty.Kind() == reflect.Array || ty.Kind() == reflect.Slice || ty.Kind() == reflect.Map {
				marshal, err := json.Marshal(v)
				if err != nil {
					v = ""
				}
				v = string(marshal)
			}
			f.SetCellValue("Sheet1", cellNum, v)
			j++
		}
	}

	tableHeader := s.All()
	for i := 0; i < len(tableHeader); i++ {
		cellFlag := fmt.Sprintf("%c1", 65+i)
		f.SetCellValue("Sheet1", cellFlag, tableHeader[i])
	}

	filename := uuid.New().String() + ".xlsx"
	outPutFile := fmt.Sprintf("%v/%v", config.FileOutPutPath, filename)
	if err = f.SaveAs(outPutFile); err != nil {
		logger.Error(err)
		return
	}
	result = filename
	return
}

func ImageToExcel(srcFile string) {

	//text := ocr.ImageToText(srcFile)
	//
	//pattern := `*[\s*]`

}
