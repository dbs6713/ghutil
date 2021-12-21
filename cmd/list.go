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

	"github.com/dbs67/ghutil/internal"
	"github.com/google/go-github/v41/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	jww "github.com/spf13/jwalterweatherman"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Repositories by Organizations",
	Long:  `List all the repositories for the user or organization.`,
	Run: func(cmd *cobra.Command, args []string) {
		listRepos(internal.Init())
	},
}

func listRepos(c *internal.Config) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Pass},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	repos, _, err := client.Repositories.ListByOrg(ctx, c.Org, nil)
	if err != nil {
		jww.FATAL.Fatal(err)
	}

	for _, repo := range repos {
		jww.INFO.Printf("%s", github.Stringify(repo.Name))
	}

	return nil
}
