package harbor

import (
	"context"
	"github.com/armosec/kubescape/cautils"
	health2 "github.com/goharbor/go-client/pkg/sdk/v2.0/client/health"
	project2 "github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"testing"
)

// TestLogin tests the logic for authenticating with harbor
func TestLogin(t *testing.T) {

	var credentials cautils.ContainerImageRegistryCredentials
	credentials.BasicAuth = make(map[string]string)
	credentials.BasicAuth["username"] = "admin"
	credentials.BasicAuth["password"] = "Harbor12345"

	harbor, err := NewHarborRegistry("http://127.0.0.1:8443", credentials)
	if err != nil {
		t.Fatal(err)
	}

	// testing hitting endpoint that required no authentication
	health, err := harbor.client.V2().Health.GetHealth(context.Background(), &health2.GetHealthParams{})
	if err != nil {
		return
	}

	t.Log(health.GetPayload().Status)

	// testing hitting endpoint which requires authentication
	project, err := harbor.client.V2().Project.GetProject(context.Background(), &project2.GetProjectParams{ProjectNameOrID: "library"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(project.GetPayload().ProjectID)

}
