package version

import "testing"

func TestVersion(t *testing.T) {

	t.Logf("UserAgent: %v", UserAgent())
	t.Logf("Version:   %v", String())
}
