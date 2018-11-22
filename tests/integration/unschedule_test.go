package integration

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/skabbass1/sepoy/commands"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestUnscheduleTask(t *testing.T) {
	err := commands.UnscheduleTask(taskName())
	if err != nil {
		t.Error(err)
	}
	got, _ := commands.TaskInfo(taskName())
	expected := `Could not find service "com.testls" in domain for`
	got = strings.TrimSpace(got)
	if got != expected {
		t.Errorf("expected output to be: %s. Got: %s.", expected, got)
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
	os.Remove(commands.PlistLocation(taskName()))
}

func taskName() string {
	return "com.testls"
}
