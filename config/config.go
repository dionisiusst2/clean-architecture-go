package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	JWTSecret   string
	CookieHash  string
	CookieBlock string
}

func readConf() (*config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	conf := &config{}

	err := viper.ReadInConfig()
	if err != nil {
		return conf, err
	}

	err = viper.Unmarshal(conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func LoadEnv() {
	conf, err := readConf()
	if err != nil {
		panic("Failed reading config file")
	}

	os.Setenv("JWT_SECRET", conf.JWTSecret)
	os.Setenv("COOKIE_HASH", conf.CookieHash)
	os.Setenv("COOKIE_BLOCK", conf.CookieBlock)

	fmt.Println("Success")
}
