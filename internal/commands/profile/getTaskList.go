// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package profile

import (
	"fmt"

	"github.com/apache/skywalking-cli/internal/logger"
	"github.com/apache/skywalking-cli/pkg/display"
	"github.com/apache/skywalking-cli/pkg/display/displayable"
	"github.com/apache/skywalking-cli/pkg/graphql/metadata"
	"github.com/apache/skywalking-cli/pkg/graphql/profile"

	"github.com/urfave/cli"
)

var getTaskListCommand = cli.Command{
	Name:      "list",
	Aliases:   []string{"l"},
	Usage:     "query profile task list",
	ArgsUsage: "[parameters...]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "service-id",
			Usage: "`<service id>` whose profile task are to be searched",
		},
		cli.StringFlag{
			Name:  "service-name",
			Usage: "`<service name>` whose profile task are to be searched",
		},
		cli.StringFlag{
			Name:  "endpoint",
			Usage: "`<endpoint>` whose profile task are to be searched",
		},
	},
	Action: func(ctx *cli.Context) error {
		serviceID := ctx.String("service-id")
		if serviceID == "" {
			serviceName := ctx.String("service-name")
			if serviceName == "" {
				return fmt.Errorf(`either flags "service-id" or "service-name" must be set`)
			}
			service, err := metadata.SearchService(ctx, serviceName)
			if err != nil {
				return err
			}
			serviceID = service.ID
		}

		endpoint := ctx.String("endpoint")

		task, err := profile.GetTaskList(ctx, serviceID, endpoint)

		if err != nil {
			logger.Log.Fatalln(err)
		}

		return display.Display(ctx, &displayable.Displayable{Data: task, Condition: serviceID})
	},
}
