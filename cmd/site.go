// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "Prints the Control Board suite version that the cb program came from.",
	Long:  `Prints the Control Board suite version that the cb program came from.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("Control Board version 0.1")
	//},
}
var siteAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Prints the Control Board suite version that the cb program came from.",
	Long:  `Prints the Control Board suite version that the cb program came from.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add ...")
	},
}

var siteRmCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete website from server.",
	Long:  `Delete website from server.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting ...")
	},
}

func init() {
	RootCmd.AddCommand(siteCmd)
	siteCmd.AddCommand(siteAddCmd)
	siteCmd.AddCommand(siteRmCmd)
	siteRmCmd.Aliases = []string{"rm", "delete" }

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
