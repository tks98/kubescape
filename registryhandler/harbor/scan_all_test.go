package harbor

import (
	"github.com/armosec/kubescape/cautils"
	"testing"
)

// TestScanAll tests scanning all harbor artifacts
func TestScanAll(t *testing.T) {

	var credentials cautils.ContainerImageRegistryCredentials
	credentials.BasicAuth = make(map[string]string)
	credentials.BasicAuth["username"] = "admin"
	credentials.BasicAuth["password"] = "Harbor12345"
	harbor, err := NewHarborRegistry("http://127.0.0.1:8443", credentials)
	if err != nil {
		t.Fatal(err)
	}

	result, err := harbor.ScanAll()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result.Error())
}
