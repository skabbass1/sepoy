package integration

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
	"testing"

	"github.com/skabbass1/sepoy/launchctl"
)

const serviceName = "com.spendthrift"

func TestMain(m *testing.M) {
	setUpPlist()
	code := m.Run()
	tearDownPlist()
	os.Exit(code)
}

func TestRunSubprocess(t *testing.T) {
	_, err := launchctl.RunSubprocess("ls", []string{"-ltr"})

	if err != nil {
		t.Error(err)
	}
}

func TestLoad(t *testing.T) {
	_, err := launchctl.Load(plistLocation())

	if err != nil {
		t.Error(err)
	}

	out, err := launchctl.List(serviceName)
	if strings.Contains(out, "Could not find") {
		t.Error(out)
	}

}

func TestStart(t *testing.T) {
	launchctl.Load(plistLocation())
	_, err := launchctl.Start(serviceName)
	if err != nil {
		t.Error(err)
	}
}

func setUpPlist() {

}

func tearDownPlist() {
	launchctl.Unload(plistLocation())
	os.Remove(plistLocation())
}

func plistLocation() string {
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, fmt.Sprintf("Library/LaunchAgents/%s.plist", serviceName))
}
