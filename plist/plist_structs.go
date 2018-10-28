package plist

func NewPlist(
	lable string,
	runAtLoad bool,
	disabled bool,
	keepAlive bool,
	launchOnlyOnce bool,
	programArguments []string,
	startCalendarInterval []map[string]int,
) *Plist {

	return &Plist{
		Label{lable},
		RunAtLoad{runAtLoad},
		Disabled{disabled},
		KeepAlive{keepAlive},
		LaunchOnlyOnce{launchOnlyOnce},
		ProgramArguments{programArguments},
		StartCalendarInterval{Intervals: constructCalendarkeys(startCalendarInterval)},
	}
}

type Label struct {
	Value string `plist:"Label"`
}

type RunAtLoad struct {
	Value bool `plist:"RunAtLoad"`
}

type Disabled struct {
	Value bool `plist:"Disabled"`
}

type KeepAlive struct {
	Value bool `plist:"KeepAlive"`
}

type LaunchOnlyOnce struct {
	Value bool `plist:"LaunchOnlyOnce"`
}

type ProgramArguments struct {
	Value []string `plist:"ProgramArguments"`
}

type StartCalendarInterval struct {
	Intervals []CalendarKeys `plist:"StartCalendarInterval,omitempty"`
}

type CalendarKeys struct {
	Month   int `plist:"Month"`
	Day     int `plist:"Day"`
	Weekday int `plist:"Weekday"`
	Hour    int `plist:"Hour"`
	Minute  int `plist:"Minute"`
}

type Plist struct {
	Label
	RunAtLoad
	Disabled
	KeepAlive
	LaunchOnlyOnce
	ProgramArguments
	StartCalendarInterval
}

func constructCalendarkeys(calendarIntervals []map[string]int) []CalendarKeys {
	calendarKeys := []CalendarKeys{}
	for _, v := range calendarIntervals {
		calendarKeys = append(
			calendarKeys,
			CalendarKeys{
				Month:   v["month"],
				Day:     v["day"],
				Weekday: v["weekday"],
				Hour:    v["hour"],
				Minute:  v["minute"],
			},
		)
	}
	return calendarKeys
}
