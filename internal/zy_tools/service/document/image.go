/*
Copyright 2024 Faw Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package document

import (
	"zy-tools/internal/zy_tools/constants"
	"zy-tools/internal/zy_tools/global"
	"zy-tools/internal/zy_tools/models/document"
	"zy-tools/pkg/doc_conv/types"
)

func (d *DocumentService) ImageToPpt(req document.ConvertRequest) (*types.ConvertResponse, error) {
	convertResponse, err := global.Office.Convert(req.FilePath, constants.FileTypePpt)
	if err != nil {
		return nil, err
	}
	return convertResponse, nil
}
