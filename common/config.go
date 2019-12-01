package common

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/ghodss/yaml"
)

//相应设置配置
type Setting struct {
	// RunMode string `yaml:"runMode"`
	Server   ServerConf   `yaml:"server"`
	Database DatabaseConf `yaml:"database"`
	// Log log `yaml:"log"`
}

//服务配置
type ServerConf struct {
	HTTPPort     int           `yaml:"HTTPPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

// 用于保存
type DatabaseConf struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

var AppConfig = &Setting{}

// Setup initialize the configuration instance
func ConfigSetup() {
	// var err error
	env := os.Getenv("ENV")

	if env == "" {
		env = "deploy"
	}

	yamlFile, _ := ioutil.ReadFile("./config/" + env + ".yaml")

	_ = yaml.Unmarshal(yamlFile, AppConfig)
}
