package harbor

import (
	"context"
	"fmt"
	"github.com/armosec/kubescape/cautils"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/artifact"
	"testing"
)

// TestImage tests scanning an artifact
func TestImage(t *testing.T) {

	var credentials cautils.ContainerImageRegistryCredentials
	credentials.BasicAuth = make(map[string]string)
	credentials.BasicAuth["username"] = "admin"
	credentials.BasicAuth["password"] = "Harbor12345"

	harbor, err := NewHarborRegistry("http://127.0.0.1:8443", credentials)
	if err != nil {
		t.Fatal(err)
	}

	imageURL := "http://127.0.0.1:8443/harbor/projects/2/repositories/tsubot/artifacts/sha256:eb51a0897f4525f2403d48169cc94bd0e52036508d1bb4181942fb888a32e318"

	imageID, err := harbor.GetImageIdentifier(imageURL)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve information about the specified image with the scan overview
	withScanOverview := true
	artifact, err := harbor.client.V2().Artifact.GetArtifact(context.Background(), &artifact.GetArtifactParams{
		ProjectName:      imageID.Project,
		Reference:        imageID.Tag,
		RepositoryName:   imageID.Repository,
		WithScanOverview: &withScanOverview,
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(artifact.GetPayload().ScanOverview["application/vnd.security.vulnerability.report; version=1.1"].EndTime)
}
