package config

import (
	"strings"
	"github.com/spf13/viper"
)

type logConfig struct {
	LogFile  string
	LogLevel string
}

type eslConfig struct {
	Host string
	Port int
}

// Config - configuration object
type Config struct {
	Log       logConfig
	EslConfig eslConfig
}

var conf *Config

// GetConfig - Function to get Config
func GetConfig() *Config {
	if conf != nil {
		return conf
	}
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	logConf := logConfig{
		LogFile:  v.GetString("log.file"),
		LogLevel: v.GetString("log.level"),
	}

	eslConf := eslConfig{
		Host: v.GetString("esl.host"),
		Port: v.GetInt("esl.port"),
	}
	conf = &Config{
		Log:       logConf,
		EslConfig: eslConf,
	}
	return conf
}
