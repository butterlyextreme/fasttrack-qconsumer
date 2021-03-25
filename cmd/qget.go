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
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

type question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Option0  string `json:"option0"`
	Option1  string `json:"option1"`
	Option2  string `json:"option2"`
}

type questions []question

const urlQuestions = "http://localhost:8010/questions"

// qgetCmd represents the qget command
var qgetCmd = &cobra.Command{
	Use:   "qget",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(urlQuestions)

		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		qs := questions{}

		jsonErr := json.Unmarshal(body, &qs)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for _, q := range qs {
			fmt.Printf("Question no [%d]\n", q.ID)
			fmt.Printf(" %s\n", q.Question)
			fmt.Printf(" [0] %s\n", q.Option0)
			fmt.Printf(" [1] %s\n", q.Option1)
			fmt.Printf(" [2] %s\n", q.Option2)
		}
	},
}

func init() {
	rootCmd.AddCommand(qgetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qgetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qgetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
