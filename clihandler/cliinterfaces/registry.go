package cliinterfaces

import (
	"github.com/containers/image/manifest"
	"time"
)

type ContainerImageRegistryCredentials struct {
	Username string
	Password string
	Tag      string
	Hash     string
}

type ContainerImageIdentifier struct {
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
	Login(registry string, credentials map[string]string) error

	// For "help" purposes
	DescribeAdaptor() string

	GetImagesScanStatus(imageIDs []ContainerImageIdentifier) ([]ContainerImageScanStatus, error)

	GetImagesVulnerabilities(imageIDs []ContainerImageIdentifier) ([]ContainerImageVulnerability, error)

	GetImagesInformation(imageIDs []ContainerImageIdentifier) ([]ContainerImageInformation, error)
}
