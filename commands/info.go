package commands

import (
	"fmt"

	"github.com/skabbass1/sepoy/launchctl"
	"github.com/urfave/cli"
)

var Info = cli.Command{
	Name:  "info",
	Usage: "display info about a scheduled task",
	Action: func(c *cli.Context) error {

		if len(c.Args()) == 0 {
			return cli.NewExitError("task name must be provided as first argument", 1)
		}
		info, err := TaskInfo(c.Args().Get(0))
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		fmt.Println(info)
		return nil
	},
}

func TaskInfo(taskName string) (string, error) {
	return launchctl.List(taskName)
}
