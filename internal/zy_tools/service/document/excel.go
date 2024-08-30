package document

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"reflect"
	"zy-tools/internal/zy_tools/constants"
	"zy-tools/internal/zy_tools/global"
	"zy-tools/internal/zy_tools/models/document"
	"zy-tools/internal/zy_tools/utils"
	"zy-tools/pkg/doc_conv/types"
)

type DocumentService struct {
}

// JsonToExcel json 转 Excel
func (d *DocumentService) JsonToExcel(jsonText string) (result string, err error) {
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
	outPutFile := fmt.Sprintf("%v/%v", constants.FileOutPutPath, filename)
	if err = f.SaveAs(outPutFile); err != nil {
		return
	}
	result = filename
	return
}

func (d *DocumentService) ExcelToPdf(req document.ConvertRequest) (*types.ConvertResponse, error) {
	convertResponse, err := global.Office.Convert(req.FilePath, constants.FileTypePdf)
	if err != nil {
		return nil, err
	}
	return convertResponse, nil
}
