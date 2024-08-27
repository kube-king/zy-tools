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

package config

import (
	"strconv"
	"strings"
)

type Config struct {
	Mysql       Mysql       `json:"mysql" yaml:"mysql"`
	System      System      `json:"system" yaml:"system"`
	Server      Server      `json:"server" yaml:"server"`
	LibreOffice LibreOffice `yaml:"libreoffice" json:"libreoffice"`
}

type Server struct {
	UploadMaxSize string `json:"uploadMaxSize" yaml:"uploadMaxSize"`
	UploadPath    string `json:"uploadPath" yaml:"uploadPath"`
}

func (s *Server) UploadMaxSizeValue() (result int64) {
	val := strings.Trim(s.UploadMaxSize, " ")
	val = strings.ToLower(val)
	if len(val) >= 2 {
		num, err := strconv.ParseInt(val[:len(val)-2], 10, 64)
		if err != nil {
			return
		}
		unit := val[len(val)-2:]
		switch unit {
		case "m":
			result = num << 20
		}
	}
	return
}

type LibreOffice struct {
	Command string `json:"command" yaml:"command"`
}

type System struct {
	DBType string `json:"dbType" yaml:"dbType"`
}

type Mysql struct {
	Port         string ` json:"port" yaml:"port"`
	Config       string ` json:"config" yaml:"config"`     // 高级配置
	Database     string ` json:"database" yaml:"database"` // 数据库名
	Username     string ` json:"username" yaml:"username"` // 数据库密码
	Password     string ` json:"password" yaml:"password"` // 数据库密码
	Host         string ` json:"host" yaml:"host"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"` // 最大空闲连接数
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"` // 最大打开连接数
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Config
}
