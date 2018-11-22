package launchctl

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
	"testing"
)

const serviceName = "com.spendthrift"

func TestMain(m *testing.M) {
	setUpPlist()
	code := m.Run()
	tearDownPlist()
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

}

func tearDownPlist() {
	Unload(plistLocation())
	os.Remove(plistLocation())
}

func plistLocation() string {
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, fmt.Sprintf("Library/LaunchAgents/%s.plist", serviceName))
}
