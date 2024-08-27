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

package common

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func NewConfig(conf any, configFileName string, configPathList ...string) error {

	viper.SetConfigName(configFileName)
	for _, configPath := range configPathList {
		viper.AddConfigPath(configPath)
	}

	var err error
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			err = fmt.Errorf("error parsing configuration file %s", err)
		}
	}

	err = viper.Unmarshal(conf)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(conf); err != nil {
			log.Println("config reload error", err)
		}
	})

	return nil
}
