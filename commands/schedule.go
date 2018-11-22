package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/skabbass1/sepoy/launchctl"
	"github.com/skabbass1/sepoy/plist"
	"github.com/urfave/cli"
)

var Schedule = cli.Command{
	Name:  "schedule",
	Usage: "schedule a servie or batch task",
	Subcommands: []cli.Command{
		{
			Name:  "batch",
			Usage: "schedule a batch task to run periodically",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "run-schedule",
					Usage: "specify run-schedule as month-day-weekday-hour-minute. E.g 0-0-7-12-0. 0 denotes a wildcard entry",
				},
			},
			Action: scheduleBatchTask,
		},
	},
}

func scheduleBatchTask(c *cli.Context) error {

	runSchedule := c.String("run-schedule")
	validationErr := ValidateScheduleBatchInput(c.Args(), runSchedule)
	if validationErr != nil {
		return cli.NewExitError(validationErr, 1)
	}

	parsedSchedule, parseErr := ParseSchedule(runSchedule)
	if parseErr != nil {
		return cli.NewExitError(parseErr, 1)
	}

	taskName := fmt.Sprintf("com.%s", c.Args().Get(0))
	taskCommand := strings.Split(c.Args().Get(1), " ")

	scheduleErr := ScheduleBatchTask(taskName, taskCommand, parsedSchedule, map[string]string{})
	if scheduleErr != nil {
		return cli.NewExitError(scheduleErr, 1)
	}
	return nil

}

func ScheduleBatchTask(
	taskName string,
	taskCommand []string,
	runSchedule map[string]int,
	taskEnvVars map[string]string) error {

	taskPlist := plist.NewPlist(
		taskName,
		false,
		false,
		false,
		false,
		taskCommand,
		[]map[string]int{
			runSchedule,
		},
		taskEnvVars,
		fmt.Sprintf("/tmp/%s.stdout", taskName),
		fmt.Sprintf("/tmp/%s.stderr", taskName),
	)
	err := plist.PublishPlist(*taskPlist, PlistLocation(taskName))
	if err != nil {
		panic(err)
	}

	_, err = launchctl.Load(PlistLocation(taskName))
	if err != nil {
		return err
	}
	return nil
}

func ValidateScheduleBatchInput(args []string, runSchedule string) error {

	if len(args) < 0 {
		return errors.New("run command not provided")
	}

	if runSchedule == "" {
		return errors.New("run-schedule must be provided")
	}
	return nil
}
