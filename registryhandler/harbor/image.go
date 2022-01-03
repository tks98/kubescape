package harbor

import (
	"context"
	"github.com/armosec/kubescape/cautils"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/artifact"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"strings"
	"time"
)

func (h *harborRegistry) ParseImageURL(imageURL string) (string, string, string, error) {

	var projectName, repositoryName, reference string

	imageURLSplit := strings.SplitAfter(imageURL, "projects")

	if len(imageURLSplit) > 1 {
		imageURLSplit1 := strings.Split(imageURLSplit[1], "/")
		if len(imageURLSplit1) > 5 {
			projectID := imageURLSplit1[1]
			projectInfo, err := h.client.V2().Project.GetProject(context.Background(), &project.GetProjectParams{ProjectNameOrID: projectID})
			if err != nil {
				return "", "", "", err
			}
			projectName = projectInfo.Payload.Name
			repositoryName = imageURLSplit1[3]
			reference = imageURLSplit1[5]

		}
	}

	return projectName, repositoryName, reference, nil

}

// GetImageIdentifier receives an ImageURL and returns a ContainerImageIdentifier type
func (h *harborRegistry) GetImageIdentifier(imageURL string) (*cautils.ContainerImageIdentifier, error) {
	projectName, repositoryName, referenceName, err := h.ParseImageURL(imageURL)
	if err != nil {
		return nil, err
	}

	// get artifact to retrieve digest hash
	artifactInfo, err := h.client.V2().Artifact.GetArtifact(context.Background(), &artifact.GetArtifactParams{ProjectName: projectName, RepositoryName: repositoryName, Reference: referenceName})
	if err != nil {
		return nil, err
	}

	var imageIdentifier cautils.ContainerImageIdentifier
	imageIdentifier.Registry = cautils.Harbor.String()
	imageIdentifier.Project = projectName
	imageIdentifier.Repository = repositoryName
	imageIdentifier.Tag = referenceName
	imageIdentifier.Hash = artifactInfo.Payload.Digest

	return &imageIdentifier, nil
}

// GetImagesScanStatus returns information about an image's scan status
func (h harborRegistry) GetImagesScanStatus(imageURL string) (*cautils.ContainerImageScanStatus, error) {

	var scanStatus cautils.ContainerImageScanStatus

	imageID, err := h.GetImageIdentifier(imageURL)
	if err != nil {
		return nil, err
	}

	scanStatus.ImageID = *imageID

	// retrieve information about the specified image with the scan overview
	withScanOverview := true
	artifact, err := h.client.V2().Artifact.GetArtifact(context.Background(), &artifact.GetArtifactParams{
		ProjectName:      imageID.Project,
		Reference:        imageID.Tag,
		RepositoryName:   imageID.Repository,
		WithScanOverview: &withScanOverview,
	})

	if err != nil {
		return nil, err
	}

	// determine the last time the image was scanned
	if scanOverview, ok := artifact.GetPayload().ScanOverview["application/vnd.security.vulnerability.report; version=1.1"]; ok {
		scanStatus.LastScanDate = time.Time(scanOverview.EndTime)
	} else {
		scanStatus.LastScanDate = time.Time{}
	}

	return &scanStatus, nil

}
