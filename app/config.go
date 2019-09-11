package app

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var CFG = &Config{}

// Config represents the config model
type Config struct {
	Mongo struct {
		Host         string        `mapstructure:"host"`
		Timeout      time.Duration `mapstructure:"timeout"`
		Username     string        `mapstructure:"username"`
		Password     string        `mapstructure:"password"`
		AuthDatabase string        `mapstructure:"auth_database"`
	} `mapstructure:"mongo"`
}

// InitConfig reads a config file and bind with Config model
func InitConfig() error {
	configName := fmt.Sprintf("config.%s", "local")

	v := viper.New()
	v.AddConfigPath("configs")
	v.SetConfigName(configName)

	if err := v.ReadInConfig(); err != nil {
		log.Infof("Read config file error: %+v", err)
		return err
	}

	if err := bindingConfig(v, CFG); err != nil {
		log.Infof("Binding config error: %+v", err)
		return err
	}
	return nil
}

func bindingConfig(vp *viper.Viper, cfg *Config) error {
	if err := vp.Unmarshal(&cfg); err != nil {
		log.Infof("Unmarshal config error: %+v", err)
		return err
	}
	return nil
}
