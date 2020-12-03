package cusUtils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func init() {
	if FileExists("./config.yaml") {
		initConfig("./config.yaml")
	}
}

var configMap map[interface{}]interface{}

func initConfig(configPath string) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		ErrInfo("配置读取失败", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal([]byte(data), &configMap)
	if err != nil {
		ErrInfo("配置读取失败", err)
		os.Exit(1)
	}
}

func GetKeyValue(key interface{}) string {
	val, ok := configMap[key]
	if ok {
		val, ok := GetStringValue(val)
		if ok {
			return val
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func GetLogLevel() string {
	return GetKeyValue("logLevel")
}
