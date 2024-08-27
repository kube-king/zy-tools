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

package initialize

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"os"
	"zy-tools/internal/zy_tools/global"
)

func Gorm() (*gorm.DB, error) {
	switch global.Config.System.DBType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

type DbFieldPlugin struct{}

func (op *DbFieldPlugin) Name() string {
	return "dbFieldPlugin"
}

func (op *DbFieldPlugin) Initialize(db *gorm.DB) (err error) {
	// 创建字段的时候雪花算法生成id
	db.Callback().Create().Before("gorm:create").Replace("id", func(db *gorm.DB) {
		node, _ := snowflake.NewNode(1)
		id := node.Generate()
		db.Statement.SetColumn("id", fmt.Sprintf("%d", id))
	})
	return
}

func RegisterTables() {
	db := global.DB
	db.Use(&DbFieldPlugin{})
	err := db.AutoMigrate()
	if err != nil {
		global.Log.Error(err, "register table failed")
		os.Exit(0)
	}
	global.Log.Info("register table success")
}
