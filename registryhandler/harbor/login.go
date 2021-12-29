package harbor

import (
	"fmt"
	"net/http"
	"net/url"
)

type harborRegistry struct {
	url      url.URL
	username string
	password string
	client   *http.Client
}

// NewHarborRegistry tests the connection to the provided harbor container registry using the provided credentials
// It returns a pointer to the harborRegistry type which can be used to perform actions on the registry
func NewHarborRegistry(rawURL, username, password string) (*harborRegistry, error) {

	var hr harborRegistry

	registryURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	hr.url = *registryURL
	hr.username = username
	hr.password = password

	err = hr.Login()
	if err != nil {
		return nil, err
	}

	return &hr, nil

}

// Login tests the connection to the harbor registry and if successful saves the http client to the caller type for future use
func (h *harborRegistry) Login() error {
	client := &http.Client{}
	endpoint := fmt.Sprintf("%s/health", h.url.String())
	req, err := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(h.username, h.password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("could not connect to harbor url, status: %s", resp.Status)
	}

	h.client = client

	return nil

}
