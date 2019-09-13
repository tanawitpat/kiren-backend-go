package app

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// CFG contains general config values. It can be accessed after initialized with function InitConfig.
var CFG = &GeneralConfig{}

// ERR contains error messages. It can be accessed after initialized with function InitConfig.
var ERR = &Errors{}

// GeneralConfig represents the general config model.
type GeneralConfig struct {
	Mongo struct {
		Host         string        `mapstructure:"host"`
		Timeout      time.Duration `mapstructure:"timeout"`
		Username     string        `mapstructure:"username"`
		Password     string        `mapstructure:"password"`
		AuthDatabase string        `mapstructure:"auth_database"`
	} `mapstructure:"mongo"`
}

// Errors represents the errors config model.
type Errors struct {
	BadRequest          Error `mapstructure:"bad_request"`
	InternalServerError Error `mapstructure:"internal_server_error"`
}

// Error represents the error config model.
type Error struct {
	Code string `mapstructure:"code"`
	Name string `mapstructure:"name"`
}

// InitConfig reads a config file and bind with config model.
func InitConfig() error {
	v := viper.New()
	v.AddConfigPath("configs")
	// Read general config file
	generalConfigPath := fmt.Sprintf("config.%s", "local")
	v.SetConfigName(generalConfigPath)

	if err := v.ReadInConfig(); err != nil {
		log.Infof("Read config file error: %+v", err)
		return err
	}
	// Bind general config to viper
	if err := bindingGeneralConfig(v, CFG); err != nil {
		log.Infof("Binding config error: %+v", err)
		return err
	}

	// Bind error config to viper
	if err := bindingErrorConfig(v, ERR); err != nil {
		log.Infof("Binding config error: %+v", err)
		return err
	}
	return nil
}

func bindingGeneralConfig(vp *viper.Viper, cfg *GeneralConfig) error {
	if err := vp.Unmarshal(&cfg); err != nil {
		log.Infof("Unmarshal config error: %+v", err)
		return err
	}
	return nil
}

func bindingErrorConfig(vp *viper.Viper, err *Errors) error {
	if err := vp.Unmarshal(&err); err != nil {
		log.Infof("Unmarshal config error: %+v", err)
		return err
	}
	return nil
}
