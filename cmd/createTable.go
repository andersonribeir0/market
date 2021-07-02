/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/andersonribeir0/market/constants"
	"github.com/andersonribeir0/market/db"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createTableCmd represents the createTable command
var createTableCmd = &cobra.Command{
	Use:   "createTable",
	Short: "Command to create dynamoDB Market table",
	Long: `Responsible for storing open market infos with great response time on write with high scalability.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.WithFields(log.Fields{"app": "market"})
		conn, err := db.NewDB()
		logger.Info("Connection opened")
		if conn != nil {
			err = conn.CreateTable(constants.TableName)
		}
		if err != nil {
			logger.Error(err.Error())
		} else {
			logger.Info("Table created")
		}
	},
}

func init() {
	rootCmd.AddCommand(createTableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createTableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
