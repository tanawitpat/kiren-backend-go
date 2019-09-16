package app

import (
	"strings"

	"github.com/spf13/viper"
)

// CFG contains app config values. It can be accessed after initialized with function InitConfig.
var CFG = &appConfig{}

// ERR contains error config values. It can be accessed after initialized with function InitConfig.
var ERR = &errorConfig{}

// Config represents the config model.
type appConfig struct {
	App struct {
		Port int    `mapstructure:"port"`
		Env  string `mapstructure:"env"`
	} `mapstructure:"app"`
	AWS struct {
		AccessKeyID     string `mapstructure:"access_key_id"`
		SecretAccessKey string `mapstructure:"secret_access_key"`
	} `mapstructure:"aws"`
}

// errorConfig represents the error config model.
type errorConfig struct {
	BadRequest          errorDetailConfig `mapstructure:"bad_request"`
	InternalServerError errorDetailConfig `mapstructure:"internal_server_error"`
}

// errorDetailConfig represents the error detail config model.
type errorDetailConfig struct {
	Code    int    `mapstructure:"code"`
	Name    string `mapstructure:"name"`
	Message string `mapstructure:"message"`
}

// InitConfig binds values from environment variables and config files with config models.
func InitConfig() error {
	v := viper.New()

	// Define a config directory
	v.AddConfigPath("configs")
	v.SetConfigType("yml")

	// Read environment variables and concatenate config hierarchy with underscore
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read and bind app config
	v.SetConfigName("app")
	if err := v.ReadInConfig(); err != nil {
		log.Infof("Read app config file error: %+v", err)
		return err
	}
	if err := bindingAppConfig(v, CFG); err != nil {
		log.Infof("Bind app config error: %+v", err)
		return err
	}

	// Read and bind error config
	v.SetConfigName("error")
	if err := v.ReadInConfig(); err != nil {
		log.Infof("Read error config file error: %+v", err)
		return err
	}
	if err := bindingErrorConfig(v, ERR); err != nil {
		log.Infof("Bind error config error: %+v", err)
		return err
	}

	return nil
}

func bindingAppConfig(vp *viper.Viper, cfg *appConfig) error {
	if err := vp.Unmarshal(&cfg); err != nil {
		log.Infof("Unmarshal config error: %+v", err)
		return err
	}
	return nil
}

func bindingErrorConfig(vp *viper.Viper, err *errorConfig) error {
	if err := vp.Unmarshal(&err); err != nil {
		log.Infof("Unmarshal config error: %+v", err)
		return err
	}
	return nil
}
