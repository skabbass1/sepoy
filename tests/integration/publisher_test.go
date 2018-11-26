package integration

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	goplist "github.com/DHowett/go-plist"
	"github.com/skabbass1/sepoy/plist"
)

func TestPublishPlist(t *testing.T) {

	myplist := plist.NewPlist(
		"mytask",
		false,
		false,
		false,
		true,
		[]string{"hello.py", "1", "2"},
		[]map[string]int{
			map[string]int{
				"month":   1,
				"day":     15,
				"weekday": 0,
				"hour":    15,
				"minute":  0,
			},

			map[string]int{
				"month":   1,
				"day":     15,
				"weekday": 0,
				"hour":    15,
				"minute":  0,
			},
		},
		map[string]string{"PATH": "/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:"},
		"/tmp/mytask.stdout",
		"/tmp/mytask.stderr",
	)

	location := "/tmp/test.plist"

	defer os.Remove(location)

	err := plist.PublishPlist(*myplist, location)
	if err != nil {
		t.Error(err)
	}

	got := plist.Plist{}
	bytes, _ := ioutil.ReadFile(location)
	goplist.Unmarshal(bytes, &got)

	if !reflect.DeepEqual(got, *myplist) {
		t.Error("published plist file does not equal expected")
	}
}
