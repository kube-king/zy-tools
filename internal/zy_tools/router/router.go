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

package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zy-tools/internal/zy_tools/constants"
	"zy-tools/internal/zy_tools/global"
)

func InitRouters() *gin.Engine {

	Router := gin.New()
	Router.Use(gin.Recovery())
	Router.MaxMultipartMemory = global.Config.Server.UploadMaxSizeValue()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	PublicGroup := Router.Group(constants.RouterPrefix)
	{

	}

	PrivateGroup := Router.Group(constants.RouterPrefix)
	{

	}

	InitDevelopmentRouter(PublicGroup, PrivateGroup)
	InitDocumentRouter(PublicGroup, PrivateGroup)
	global.Log.Info("init router success")
	return Router

}
