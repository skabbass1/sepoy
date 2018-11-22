package unit

import (
	"reflect"
	"testing"

	"github.com/skabbass1/sepoy/commands"
)

func TestValidParseSchedule(t *testing.T) {
	schedule := "0-0-7-12-0"
	got, error := commands.ParseSchedule(schedule)

	if error != nil {
		t.Error(error)
	}
	expected := map[string]int{
		"month":   0,
		"day":     0,
		"weekday": 7,
		"hour":    12,
		"minute":  0,
	}
	eq := reflect.DeepEqual(expected, got)
	if !eq {
		t.Errorf("expected: %v got: %v", expected, got)
	}

}

func TestInValidParseSchedule(t *testing.T) {
	testCases := []string{
		"",
		"00",
		"0}12-13",
		"0-0-7-12-0-1",
		"a-b-c-d-e-f",
	}

	for _, tc := range testCases {
		_, err := commands.ParseSchedule(tc)
		if err == nil {
			t.Errorf("expected error to be not nil for test case:%s", tc)
		}
	}
}
