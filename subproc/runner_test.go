package subproc

import (
	"testing"
)

func TestRunSubprocess(t *testing.T) {
	_, err := RunSubprocess("ls", []string{"-ltr"})

	if err != nil {
		t.Error(err)
	}
}
