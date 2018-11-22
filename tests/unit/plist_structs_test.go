package unit

import (
	"testing"

	goplist "github.com/DHowett/go-plist"
	"github.com/skabbass1/sepoy/plist"
)

func TestNewPlistCreation(t *testing.T) {
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

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
 <dict>
  <key>Disabled</key>
  <false/>
  <key>EnvironmentVariables</key>
  <dict>
   <key>PATH</key>
   <string>/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:</string>
  </dict>
  <key>KeepAlive</key>
  <false/>
  <key>Label</key>
  <string>mytask</string>
  <key>LaunchOnlyOnce</key>
  <true/>
  <key>ProgramArguments</key>
  <array>
   <string>hello.py</string>
   <string>1</string>
   <string>2</string>
  </array>
  <key>RunAtLoad</key>
  <false/>
  <key>StandardErrorPath</key>
  <string>/tmp/mytask.stderr</string>
  <key>StandardOutPath</key>
  <string>/tmp/mytask.stdout</string>
  <key>StartCalendarInterval</key>
  <array>
   <dict>
    <key>Day</key>
    <integer>15</integer>
    <key>Hour</key>
    <integer>15</integer>
    <key>Minute</key>
    <integer>0</integer>
    <key>Month</key>
    <integer>1</integer>
    <key>Weekday</key>
    <integer>0</integer>
   </dict>
   <dict>
    <key>Day</key>
    <integer>15</integer>
    <key>Hour</key>
    <integer>15</integer>
    <key>Minute</key>
    <integer>0</integer>
    <key>Month</key>
    <integer>1</integer>
    <key>Weekday</key>
    <integer>0</integer>
   </dict>
  </array>
 </dict>
</plist>`

	out, err := goplist.MarshalIndent(myplist, goplist.XMLFormat, " ")
	if err != nil {
		t.Error(err)
	}
	got := string(out)
	if got != expected {
		t.Error("rendered plist xml does not equal expected plist xml")
	}
}
