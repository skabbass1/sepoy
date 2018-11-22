package commands

import (
	"errors"
	"fmt"
	"os/user"
	"path"
	"strings"

	"github.com/skabbass1/sepoy/launchctl"
	"github.com/skabbass1/sepoy/plist"
	"github.com/urfave/cli"
)

var Schedule = cli.Command{
	Name:  "schedule",
	Usage: "schedule a servie or batch job",
	Subcommands: []cli.Command{
		{
			Name:  "batch",
			Usage: "schedule a batch job to run periodically",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "run-schedule",
					Usage: "specify run-schedule as month-weekday-hour-minute. E.g 0-0-7-12-0. 0 denotes a wildcard entry",
				},
			},
			Action: scheduleBatch,
		},
	},
}

func scheduleBatch(c *cli.Context) *cli.ExitError {

	runSchedule := c.String("run-schedule")
	validationErr := ValidateScheduleBatchInput(c.Args(), runSchedule)
	if validationErr != nil {
		return cli.NewExitError(validationErr, 1)
	}

	parsedSchedule, parseErr := ParseSchedule(runSchedule)
	if parseErr != nil {
		return cli.NewExitError(parseErr, 1)
	}

	jobName := fmt.Sprintf("com.%s", c.Args().Get(0))
	jobCommand := strings.Split(c.Args().Get(1), " ")

	scheduleErr := ScheduleBatchJob(jobName, jobCommand, parsedSchedule, map[string]string{})
	if scheduleErr != nil {
		return cli.NewExitError(scheduleErr, 1)
	}
	return nil

}

func ScheduleBatchJob(
	jobName string,
	jobCommand []string,
	runSchedule map[string]int,
	jobEnvVars map[string]string) error {

	jobPlist := plist.NewPlist(
		jobName,
		false,
		false,
		false,
		false,
		jobCommand,
		[]map[string]int{
			runSchedule,
		},
		jobEnvVars,
		fmt.Sprintf("/tmp/%s.stdout", jobName),
		fmt.Sprintf("/tmp/%s.stderr", jobName),
	)
	err := plist.PublishPlist(*jobPlist, plistLocation(jobName))
	if err != nil {
		panic(err)
	}

	_, err = launchctl.Load(plistLocation(jobName))
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

func plistLocation(jobName string) string {
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, fmt.Sprintf("Library/LaunchAgents/%s.plist", jobName))
}
