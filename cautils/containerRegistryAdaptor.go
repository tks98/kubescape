package cautils

import (
	"github.com/containers/image/manifest"
	"time"
)

type ContainerImageRegistryCredentials struct {
	BasicAuth map[string]string
}

type ContainerImageIdentifier struct {
	Project string
	Registry   string
	Repository string
	Tag        string
	Hash       string
}

type ContainerImageScanStatus struct {
	ImageID         ContainerImageIdentifier
	IsScanAvailable bool
	IsBomAvailable  bool
	LastScanDate    time.Time
}

type ContainerImageVulnerability struct {
	ImageID ContainerImageIdentifier
	CVEs    map[string]interface{}
}

type ContainerImageInformation struct {
	ImageID       ContainerImageIdentifier
	Bom           []string
	ImageManifest manifest.Manifest // will use here Docker package definition
}

type IContainerImageVulnerabilityAdaptor interface {
	// Credentials are coming from user input (CLI or configuration file) and they are abstracted at string to string map level
	// so and example use would be like registry: "simpledockerregistry:80" and credentials like {"username":"joedoe","password":"abcd1234"}
	Login(registry string, credentials ContainerImageRegistryCredentials) error

	// For "help" purposes
	DescribeAdaptor() string

	GetImagesScanStatus(imageIDs []ContainerImageIdentifier) ([]ContainerImageScanStatus, error)

	GetImagesVulnerabilities(imageIDs []ContainerImageIdentifier) ([]ContainerImageVulnerability, error)

	GetImagesInformation(imageIDs []ContainerImageIdentifier) ([]ContainerImageInformation, error)
}
