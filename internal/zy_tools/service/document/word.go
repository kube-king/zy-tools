package document

import (
	"zy-tools/internal/zy_tools/constants"
	"zy-tools/internal/zy_tools/global"
	"zy-tools/internal/zy_tools/models/document"
	"zy-tools/pkg/doc_conv/types"
)

func (d *DocumentService) WordToPdf(req document.ConvertRequest) (*types.ConvertResponse, error) {
	convertResponse, err := global.Office.Convert(req.FilePath, constants.FileTypePdf)
	if err != nil {
		return nil, err
	}
	return convertResponse, nil
}
