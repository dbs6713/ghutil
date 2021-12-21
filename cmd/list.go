// Package cmd provides commands for the application.
// Copyright 2021 Don B. Stringham All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.
//
// @author dbs67 <donstringham@weber.edu>
//
package cmd

import (
	"context"

	"github.com/google/go-github/v41/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	jww "github.com/spf13/jwalterweatherman"
)

type config struct {
	User string
	Pass string
	Org  string
	Url  string
}

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Repositories by Organizations",
	Long:  `List all the repositories for the user or organization.`,
	Run: func(cmd *cobra.Command, args []string) {
		listRepos(initConfig())
	},
}

func initConfig() *config {
	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	viper.SetConfigName("ghutil")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		jww.FATAL.Fatal(err)
	}

	return &config{
		User: viper.GetString("User"),
		Pass: viper.GetString("Pass"),
		Org:  viper.GetString("Org"),
		Url:  viper.GetString("Url"),
	}

}

func listRepos(c *config) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Pass},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	repos, _, err := client.Repositories.ListByOrg(ctx, "cs3210-fall2021", nil)
	if err != nil {
		jww.FATAL.Fatal(err)
	}

	jww.INFO.Printf("%v", repos[0])

	for _, repo := range repos {
		jww.INFO.Printf("%s", github.Stringify(repo.FullName))
	}

	return nil
}
