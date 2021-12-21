// Package cmd provides commands for the application.
// Copyright 2021 Don B. Stringham All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.
//
// @author dbs67 <donstringham@weber.edu>
//
package internal

import (
	"github.com/spf13/viper"

	jww "github.com/spf13/jwalterweatherman"
)

type Config struct {
	User string
	Pass string
	Org  string
	Url  string
}

func Init() *Config {
	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	viper.SetConfigName("ghutil")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		jww.FATAL.Fatal(err)
	}

	return &Config{
		User: viper.GetString("User"),
		Pass: viper.GetString("Pass"),
		Org:  viper.GetString("Org"),
		Url:  viper.GetString("Url"),
	}
}
