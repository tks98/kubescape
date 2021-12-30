package harbor

import (
	gc "github.com/goharbor/go-client/pkg/harbor"
	"net/url"
)

type harborRegistry struct {
	url      url.URL
	username string
	password string
	client   *gc.ClientSet
}

// NewHarborRegistry returns a pointer to the harborRegistry type which can be used to perform actions on the registry
func NewHarborRegistry(rawURL, username, password string, insecure bool) (*harborRegistry, error) {

	registryURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	c := &gc.ClientSetConfig{
		URL:      registryURL.String(),
		Insecure: insecure,
		Username: username,
		Password: password,
	}

	harborClient, err := gc.NewClientSet(c)
	if err != nil {
		return nil, err
	}

	var hr harborRegistry

	hr.url = *registryURL
	hr.username = username
	hr.password = password
	hr.client = harborClient

	return &hr, nil

}
