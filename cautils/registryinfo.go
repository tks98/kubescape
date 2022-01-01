package cautils

type Registry string

const (
	Harbor Registry = ""
)

func (r Registry) String() string {
	switch r {
	case Harbor:
		return "harbor"
	}

	return "unknown"
}

type RegistryInfo struct {
	All      bool   // specify to scan all images
	Image    string // specify the image to scan
	URL      string // URL of the image registry
	Username string // Username to login to the registry
	Password string // Password to login to the registry
	Name     string // Name of the registry
}
