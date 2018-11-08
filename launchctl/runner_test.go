package launchctl

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
	"testing"

	"github.com/skabbass1/sepoy/plist"
)

const serviceName = "com.spendthrift"

func TestMain(m *testing.M) {
	setUpPlist()
	code := m.Run()
	// tearDownPlist()
	os.Exit(code)
}

func TestRunSubprocess(t *testing.T) {
	_, err := RunSubprocess("ls", []string{"-ltr"})

	if err != nil {
		t.Error(err)
	}
}

func TestLoad(t *testing.T) {
	_, err := Load(plistLocation())

	if err != nil {
		t.Error(err)
	}

	out, err := List(serviceName)
	if strings.Contains(out, "Could not find") {
		t.Error(out)
	}

}

func TestStart(t *testing.T) {
	Load(plistLocation())
	_, err := Start(serviceName)
	if err != nil {
		t.Error(err)
	}
}

func setUpPlist() {

	myplist := plist.NewPlist(
		"com.spendthrift",
		false,
		false,
		false,
		false,
		[]string{"/Applications/Docker.app/Contents/Resources/bin/docker", "run", "--rm", "skabbass1/spendthrift:v0.0.1"},
		[]map[string]int{
			map[string]int{
				"month":   1,
				"day":     15,
				"weekday": 0,
				"hour":    15,
				"minute":  0,
			},
		},
		map[string]string{"PATH": "/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:"},
		"/tmp/test_spendthrift.stdout",
		"/tmp/test_spendthrift.stderr",
	)
	err := plist.PublishPlist(*myplist, plistLocation())
	if err != nil {
		panic(err)
	}

}

func tearDownPlist() {
	Unload(plistLocation())
	os.Remove(plistLocation())
}

func plistLocation() string {
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, fmt.Sprintf("Library/LaunchAgents/%s.plist", serviceName))
}
