package cmd

import (
	"testing"
	"bytes"
	"fmt"
)

func TestVersion(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := NewRootCmd()
	cmd.SetOutput(buf)
	cmd.SetArgs([]string{"version"})

	if err := cmd.Execute(); err != nil {
		t.Errorf("??????")
	}

	actual := buf.String()

	expect := fmt.Sprintf("version %s", "0.1")
	//expect := "version "
	if actual != expect {
		t.Errorf("unexpected response: expect:%+v, actual:%+v", expect, actual)
	}
}


