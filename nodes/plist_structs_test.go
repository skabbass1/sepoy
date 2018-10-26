package nodes

import (
	"fmt"
	"strings"
	"testing"

	goplist "github.com/DHowett/go-plist"
)

func TestNewPlistCreation(t *testing.T) {
	myplist := NewPlist(
		"com.vault",
		false,
		false,
		false,
		true,
		[]string{"hello.py", "1", "2"},
		"",
	)

	expected := `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Disabled</key>
		<false/>
		<key>KeepAlive</key>
		<false/>
		<key>Label</key>
		<string>com.vault</string>
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
	</dict>
`
	out, err := goplist.MarshalIndent(myplist, goplist.XMLFormat, "")
	if err != nil {
		t.Error(err)
	}

	got := string(out)
	expected = strings.Replace(strings.TrimSpace(expected), "\n", "", -1)
	fmt.Println(expected)
	if got != strings.TrimSpace(expected) {
		t.Fail()
	}

}
