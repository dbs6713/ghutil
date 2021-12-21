// Package cmd provides commands for the application.
// Copyright 2020 Don B. Stringham All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.
//
// @author donbstringham <donbstringham@icloud.com>
//
package cmd

import (
	"fmt"

	"github.com/dbs67/ghutil/ver"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number of application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("VERSION:\t %s\n", ver.Version)
		fmt.Printf("BUILD TIME:\t %s\n", ver.Buildtime)
		fmt.Printf("BUILD USER:\t %s\n", ver.Builduser)
	},
}
