package nodes

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
