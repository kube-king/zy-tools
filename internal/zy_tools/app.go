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

package zy_tools

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"zy-tools/internal/zy_tools/global"
	"zy-tools/internal/zy_tools/initialize"
	"zy-tools/internal/zy_tools/router"
)

// Gin 配置
const (
	ListenAddr     = "0.0.0.0:8080"    // Gin配置 服务端口
	ReadTimeout    = 180 * time.Second // Gin配置 读取数据超时时间
	WriteTimeout   = 180 * time.Second // Gin配置 写入超时时间
	MaxHeaderBytes = 1 << 20           // Gin配置 Header 最大字节大小
	IdleTimeout    = 180 * time.Second // Gin配置 空闲超时时间
)

func App() {

	fmt.Println(1)
	var err error
	global.Log = initialize.InitLogger()

	global.Log.Info("init logger success")
	global.Config, err = initialize.InitConfig()
	if err != nil {
		global.Log.Error(err, "init config error: ")
		return
	}
	global.Log.Info("config content:", "config", global.Config)
	global.Log.Info("init config success")
	//global.DB, err = initialize.Gorm()
	//if err != nil {
	//	global.Log.Error(err, "init db error")
	//	return
	//}
	//global.Log.Info("init database success")
	//initialize.RegisterTables()
	//global.Log.Info("register tables success")

	global.Office = initialize.InitOffice()
	global.Tika = initialize.InitTika()
	global.Pdfbox = initialize.InitPdfbox()
	global.Pandoc = initialize.InitPandoc()
	global.Pdf2docx = initialize.InitPdf2docx()
	initialize.InitValidator()

	g := router.InitRouters()
	srv := &http.Server{
		Addr:           ListenAddr,
		Handler:        g,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
		IdleTimeout:    IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Error(err, "listen server error")
		}
	}()

	global.Log.Info(fmt.Sprintf("listen: %v ,success!", ListenAddr))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.Log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Log.Error(err, "Server Shutdown:")
	}
	global.Log.Info("Server exiting")
}
