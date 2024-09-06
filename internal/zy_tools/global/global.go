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

package global

import (
	"github.com/go-logr/logr"
	"gorm.io/gorm"
	"zy-tools/internal/zy_tools/config"
	"zy-tools/pkg/doc_conv"
	"zy-tools/pkg/pandoc"
	"zy-tools/pkg/pdfbox"
	"zy-tools/pkg/tika"
)

var (
	Log    logr.Logger
	Config *config.Config
	DB     *gorm.DB
	Office doc_conv.Office
	Pdfbox *pdfbox.Pdfbox
	Tika   *tika.Tika
	Pandoc *pandoc.Pandoc
)
