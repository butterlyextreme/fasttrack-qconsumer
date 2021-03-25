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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type answer struct {
	ID     int `json:"id"`
	Answer int `json:"answer"`
}

type score struct {
	Result int     `json:"result"`
	Grade  float64 `json:"grade"`
}

type answers []answer

const ID = 0
const ANSWER = 1
const urlResult = "http://localhost:8010/result"

// qansCmd represents the qans command
var qansCmd = &cobra.Command{
	Use:   "qans",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		numArgs := len(args)
		ans := make(answers, numArgs)
		duplicate := make(map[string]int)

		if numArgs > 5 {
			log.Fatalf("No more then 5 answers are expected, you have submitted [%d]", numArgs)
		}

		for i := 0; i < numArgs; i++ {
			a := strings.Split(args[i], ",")
			id := a[ID]

			if duplicate[id] != 0 {
				log.Fatalf("Only one answer per question is permitted,question [%s] has been answered more then once", id)
			}
			duplicate[id] = 1

			ans[i].ID, err = strconv.Atoi(id)
			ans[i].Answer, err = strconv.Atoi(a[ANSWER])

			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Print(postResults(ans))
	},
}

func postResults(ans []answer) string {
	var s score

	payloadBuf := new(bytes.Buffer)
	_ = json.NewEncoder(payloadBuf).Encode(&ans)
	req, _ := http.NewRequest("POST", urlResult, payloadBuf)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	jsonErr := json.Unmarshal(body, &s)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return fmt.Sprintf("You had [%d] answers correct, and are in the top %3.0f%%", s.Result ,s.Grade)
}

func init() {
	rootCmd.AddCommand(qansCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qansCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qansCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
