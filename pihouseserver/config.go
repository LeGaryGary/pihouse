package main

import (
	"github.com/spf13/viper"
)

func GetSqlConnectionString() string {
	return viper.GetString("SqlServerConnectionString")
}
