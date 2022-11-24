package pkg

import "testing"

func TestPrintVersion(t *testing.T) {
	v := PrintVersion()
	t.Log(v)
}
