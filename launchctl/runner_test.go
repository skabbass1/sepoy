package launchctl

import (
	"os"
	"testing"
)

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
	_, err := Load("com.spendthrif")

	if err != nil {
		t.Error(err)
	}
}

func setUpPlist() {

}

func tearDownPlist() {

}
