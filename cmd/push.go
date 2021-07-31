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
	"github.com/rainss/ossupload/src/upload"
	"github.com/spf13/cobra"
)

var (
	//ossFilePath     string
	//localFilePath string
	isRecursive bool = false
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push file to oss bucket",

	Long: `example:
ossupload push ossPath localFile/localPath `,
	//必须填写2个Args
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		upload.OssUpload(args[0], args[1])

	},
}

func init() {

	rootCmd.AddCommand(pushCmd)

	//pushCmd.Flags().BoolVarP(&isRecursive, "r", "r",
	//	false, "Upload the file recursively")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
