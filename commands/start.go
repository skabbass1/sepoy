package commands

import (
	"github.com/skabbass1/sepoy/launchctl"
	"github.com/urfave/cli"
)

var Start = cli.Command{
	Name:  "start",
	Usage: "manually trigger a  scheduled task",
	Action: func(c *cli.Context) error {

		if len(c.Args()) == 0 {
			return cli.NewExitError("task name must be provided as first argument", 1)
		}
		_, err := TaskStart(c.Args().Get(0))
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		return nil
	},
}

func TaskStart(taskName string) (string, error) {
	return launchctl.Start(taskName)
}
