package config

import (
	"fmt"

	"exercises_go/3_exercise_go_viperNlog/logging"
	"github.com/spf13/viper"
)

func GetConfig(Key string) string {
	viper.AddConfigPath("3_exercise_go_viperNlog/config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		logging.LogERROR("Can not read config: " + err.Error())
	}

	get := viper.Get(Key)

	res := fmt.Sprintf("%v", get)
	return res
}
