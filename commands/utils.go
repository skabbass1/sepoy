package commands

import (
	"errors"
	"strings"
	"strconv"
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
