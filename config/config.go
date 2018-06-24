/*
 * Copyright (c) 2018
 * time:   6/24/18 3:22 PM
 * author: linhuanchao
 * e-mail: 873085747@qq.com
 */

package config

import (
	"path/filepath"
	"os"
	"github.com/BurntSushi/toml"
	"sync"
)

type Config struct {
	DictionaryPath string
	Port           string
	PidFilePath    string
}

var once = sync.Once{}
var config *Config

func GetConfig() *Config{
	once.Do(func() {
		currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		toml.DecodeFile(currentPath+"/config/config.toml", &config)
	})
	return config
}