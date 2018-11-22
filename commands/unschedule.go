package commands

import (
	"github.com/skabbass1/sepoy/launchctl"
	"github.com/urfave/cli"
)

var Unschedule = cli.Command{
	Name:  "unschedule",
	Usage: "unschedule a servie or batch task",
	Action: func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			return cli.NewExitError("task name must be provided as first argument", 1)
		}

		err := UnscheduleTask(c.Args().Get(0))
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		return nil
	},
}

func UnscheduleTask(taskName string) error {
	_, err := launchctl.Unload(PlistLocation(taskName))
	return err
}
