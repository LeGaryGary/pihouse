package main

import (
	"github.com/spf13/viper"
)

func GetSqlConnectionString() (connectionString string) {
	return viper.GetString("SqlServerConnectionString")
}
