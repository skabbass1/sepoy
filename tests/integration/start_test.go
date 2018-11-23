package integration

import (
	"log"
	"os"
	"testing"

	"github.com/skabbass1/sepoy/commands"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestTaskStart(t *testing.T) {
	out, err := commands.TaskStart(taskName())
	if err != nil {
		t.Errorf("%s : %v", out, err)
	}
}

func setUp() {
	err := commands.ScheduleBatchTask(
		taskName(),
		[]string{"ls", "-ltr"},
		map[string]int{
			"month":   0,
			"day":     0,
			"weekday": 1,
			"hour":    13,
			"minute":  0,
		},
		map[string]string{},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	defer os.Remove(commands.PlistLocation(taskName()))
	err := commands.UnscheduleTask(taskName())
	if err != nil {
		log.Fatal(err)
	}
}

func taskName() string {
	return "com.testls"
}
