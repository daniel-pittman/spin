// Copyright (c) 2019, Google, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package execution

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/daniel-pittman/spin/cmd/gateclient"
	"github.com/daniel-pittman/spin/util"
)

var (
	statusExecutionShort = "Get the status of an execution"
	statusExecutionLong  = "Get the status of an execution with the provided event id "
)

func NewStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: statusExecutionShort,
		Long:  statusExecutionLong,
		RunE:  getExecutionStatus,
	}
	return cmd
}

func getExecutionStatus(cmd *cobra.Command, args []string) error {
	gateClient, err := gateclient.NewGateClient(cmd.InheritedFlags())
	if err != nil {
		return err
	}

	id, err := util.ReadArgsOrStdin(args)
	if err != nil {
		return err
	}

	executions := make([]interface{}, 0)
	executions, resp, err := gateClient.ExecutionsControllerApi.SearchForPipelineExecutionsByTriggerUsingGET(
		gateClient.Context,
		"*",
		map[string]interface{}{
			"eventId": id,
		})

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error getting execution %s, status code: %d\n",
			id,
			resp.StatusCode)
	}

	util.UI.JsonOutput(executions[0], util.UI.OutputFormat)
	return nil
}
