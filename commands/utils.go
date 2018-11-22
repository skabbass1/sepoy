package commands

import (
	"errors"
	"fmt"
	"os/user"
	"path"
	"strconv"
	"strings"
)

func ParseSchedule(schedule string) (map[string]int, error) {
	parts := strings.Split(schedule, "-")
	const numValues = 5
	if len(parts) != numValues {
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

func PlistLocation(jobName string) string {
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, fmt.Sprintf("Library/LaunchAgents/%s.plist", jobName))
}
