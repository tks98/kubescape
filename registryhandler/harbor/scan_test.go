package harbor

import (
	"github.com/armosec/kubescape/cautils"
	"testing"
)

// TestScan tests scanning an artifact
func TestScan(t *testing.T) {

	var credentials cautils.ContainerImageRegistryCredentials
	credentials.BasicAuth = make(map[string]string)
	credentials.BasicAuth["username"] = "admin"
	credentials.BasicAuth["password"] = "Harbor12345"

	harbor, err := NewHarborRegistry("http://127.0.0.1:8443", credentials)
	if err != nil {
		t.Fatal(err)
	}

	imageURL := "http://127.0.0.1:8443/harbor/projects/2/repositories/tsubot/artifacts/sha256:eb51a0897f4525f2403d48169cc94bd0e52036508d1bb4181942fb888a32e318"

	result, err := harbor.ScanImage(imageURL)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result.Error())

}
