package harbor

import (
	"context"
	"fmt"
	gc "github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/health"
	"net/url"
)

type harborRegistry struct {
	url      url.URL
	username string
	password string
	client   *gc.ClientSet
}

// NewHarborRegistry returns a pointer to the harborRegistry type which can be used to perform actions on the registry
func NewHarborRegistry(rawURL string, credentials map[string]string) (*harborRegistry, error) {

	registryURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	var hr harborRegistry

	err = hr.Login(registryURL.String(), credentials)
	if err != nil {
		return nil, err
	}

	hr.url = *registryURL

	return &hr, nil

}

// Login verifies the user supplied a valid registry and login credentials and stores those in the local config
func (h *harborRegistry) Login(registry string, credentials map[string]string) error {

	if _, ok := credentials["username"]; !ok {
		return fmt.Errorf("you need to supply a username to login to the registry")
	}

	if _, ok := credentials["password"]; !ok {
		return fmt.Errorf("you need to supply a password to login to the registry")
	}

	c := &gc.ClientSetConfig{
		URL:      registry,
		Username: credentials["username"],
		Password: credentials["password"],
	}

	harborClient, err := gc.NewClientSet(c)
	if err != nil {
		return err
	}

	_, err = harborClient.V2().Health.GetHealth(context.Background(), &health.GetHealthParams{})
	if err != nil {
		return err
	}

	h.client = harborClient
	h.username = credentials["username"]
	h.password = credentials["password"]
	h.client = harborClient

	return nil

}
