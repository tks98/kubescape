package harbor

import "testing"

// TestScanAll tests scanning all harbor artifacts
func TestScanAll(t *testing.T) {

	harbor, err := NewHarborRegistry("http://127.0.0.1:8443", "admin", "Harbor12345")
	if err != nil {
		t.Fatal(err)
	}

	result, err := harbor.ScanAll()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result.Error())
}
