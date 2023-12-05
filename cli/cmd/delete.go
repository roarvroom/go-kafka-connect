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
	"github.com/roarvroom/go-kafka-connect/lib/connectors"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing connector",
	RunE:  RunEDelete,
}

// RunEDelete ...
func RunEDelete(cmd *cobra.Command, args []string) error {
	req := connectors.ConnectorRequest{
		Name: connector,
	}

	resp, err := getClient().DeleteConnector(req, sync)
	if err != nil {
		return err
	}

	return printResponse(resp)
}

func init() {
	RootCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().BoolVarP(&sync, "sync", "y", false, "execute synchronously")
	deleteCmd.PersistentFlags().StringVarP(&connector, "connector", "n", "", "name of the target connector")
}
