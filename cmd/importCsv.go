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
	"fmt"
	"github.com/andersonribeir0/market/model"
	"github.com/andersonribeir0/market/parser"
	"github.com/andersonribeir0/market/repository"
	"os"

	"github.com/spf13/cobra"
)

// importCsvCmd represents the importCsv command
var importCsvCmd = &cobra.Command{
	Use:   "importCsv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var ps parser.CsvFileParser
		repo := &repository.MarketRepository {}
		ps = *ps.New("./market/csv/", "DEINFO_AB_FEIRASLIVRES_2014.csv")
		result, err  := ps.Parse()
		if err != nil {
			fmt.Fprintf(os.Stdout, "Failed to import %s", err.Error())
			return
		}
		err = repo.New()
		var record model.Record
		records, err := record.FromRecordMapList(result)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Failed to parse to model map list %s", err.Error())
			return
		}

		for i := range records {
			err = repo.Save(records[i])
			if err != nil {
				fmt.Fprintf(os.Stdout, "Failed to put item %#v error: %s", records[i], err.Error())
			}
		}
		fmt.Fprintf(os.Stdout, "Successfully saved\n")
	},
}


func init() {
	rootCmd.AddCommand(importCsvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCsvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCsvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
