package harbor

import (
	"testing"
)

// TestScan tests scanning an artifact
func TestScan(t *testing.T) {

	harbor, err := NewHarborRegistry("http://127.0.0.1:8443", "admin", "Harbor12345")
	if err != nil {
		t.Fatal(err)
	}

	project, repository, reference, err := harbor.ParseImageURL("http://127.0.0.1:8443/harbor/projects/2/repositories/tsubot/artifacts/sha256:eb51a0897f4525f2403d48169cc94bd0e52036508d1bb4181942fb888a32e318")
	if err != nil {
		return
	}

	result, err := harbor.ScanImage(project, repository, reference)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result.Error())
}
