package commands

import (
	"errors"
	"fmt"
	"os/user"
	"path"
	"strconv"
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
					Name:  "schedule",
					Usage: "specify schedule as month-weekday-hour-minute. E.g 0-0-7-12-0. 0 denotes a wildcard entry",
				},
			},
			Action: scheduleBatch,
		},
	},
}

func scheduleBatch(c *cli.Context) error {
	if len(c.Args()) < 0 {
		return cli.NewExitError("run command not provided", 1)
	}

	jobName := fmt.Sprintf("com.%s", c.Args().Get(0))
	jobCommand := strings.Split(c.Args().Get(1), " ")
	jobPlist := plist.NewPlist(
		jobName,
		false,
		false,
		false,
		false,
		jobCommand,
		[]map[string]int{
			map[string]int{
				"month":   0,
				"day":     0,
				"weekday": 7,
				"hour":    12,
				"minute":  0,
			},
		},
		map[string]string{"PATH": "/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:"},
		fmt.Sprintf("/tmp/%s.stdout", jobName),
		fmt.Sprintf("/tmp/%s.stderr", jobName),
	)
	err := plist.PublishPlist(*jobPlist, plistLocation(jobName))
	if err != nil {
		panic(err)
	}

	out, err := launchctl.Load(plistLocation(jobName))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	return nil

}

func plistLocation(jobName string) string {
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, fmt.Sprintf("Library/LaunchAgents/%s.plist", jobName))
}

func ParseSchedule(schedule string) (map[string]int, error) {
	parts := strings.Split(schedule, "-")
	const numValues = 5
	if len(parts) < numValues {
		return nil, errors.New("invalid schedule format")
	}
	month, monthErr := strconv.Atoi(parts[0])
	day, dayErr := strconv.Atoi(parts[1])
	weekday, weekdayErr := strconv.Atoi(parts[2])
	hour, hourErr := strconv.Atoi(parts[3])
	minute, minuteErr := strconv.Atoi(parts[4])

	if monthErr != nil || dayErr != nil || weekdayErr != nil || hourErr != nil || minuteErr != nil {
		return nil, errors.New("invalid types specified in schedule")
	}

	return map[string]int{
		"month":   month,
		"day":     day,
		"weekday": weekday,
		"hour":    hour,
		"minute":  minute,
	}, nil
}
