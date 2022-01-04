package harbor

import (
	"context"
	"fmt"
	"github.com/armosec/kubescape/cautils"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/artifact"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"strings"
	"time"
)

// TODO - add ability for user to supply direct image URL or image pull string
// ex: kubescape registry scan-status --image=http://127.0.0.1:8443/tks98/tsubot:1.0
// ex: kubescape registry scan-status --image=http://127.0.0.1:8443/harbor/projects/2/repositories/tsubot/artifacts/sha256:eb51a0897f4525f2403d48169cc94bd0e52036508d1bb4181942fb888a32e318
func (h *harborRegistry) ParseImageURL(imageURL string) (string, string, string, string, bool, error) {

	var projectName, repositoryName, reference, tag string
	var tagSpecified bool

	if strings.Contains(imageURL, "artifacts/sha256:") {
		imageURLSplit := strings.SplitAfter(imageURL, "projects")
		if len(imageURLSplit) > 1 {
			imageURLSplit1 := strings.Split(imageURLSplit[1], "/")
			if len(imageURLSplit1) > 5 {
				projectID := imageURLSplit1[1]
				projectInfo, err := h.client.V2().Project.GetProject(context.Background(), &project.GetProjectParams{ProjectNameOrID: projectID})
				if err != nil {
					return "", "", "", "", false, err
				}
				projectName = projectInfo.Payload.Name
				repositoryName = imageURLSplit1[3]
				reference = imageURLSplit1[5]
			}
		}
	} else {
		imageURLSplit := strings.Split(imageURL, "/")

		if len(imageURLSplit) < 5 {
			return "", "", "", "", false, fmt.Errorf("image url is invalid %s", imageURL)
		}

		projectName = imageURLSplit[3]
		repoTag := imageURLSplit[4]

		repoTagSplit := strings.Split(repoTag, ":")
		if len(repoTagSplit) != 2 {
			return "", "", "", "", false, fmt.Errorf("image url is invalid. Did you forget the tag? %s", imageURL)
		}

		repositoryName = repoTagSplit[0]
		tag = repoTagSplit[1]
		tagSpecified = true

		artifact, err := h.client.V2().Artifact.GetArtifact(context.Background(), &artifact.GetArtifactParams{
			ProjectName:    projectName,
			Reference:      tag,
			RepositoryName: repositoryName,
			WithTag:        &tagSpecified,
		})
		if err != nil {
			return "", "", "", "", false, err
		}

		reference = artifact.GetPayload().Digest
	}

	return projectName, repositoryName, reference, tag, tagSpecified, nil

}

// GetImageIdentifier receives an ImageURL and returns a ContainerImageIdentifier type
func (h *harborRegistry) GetImageIdentifier(imageURL string) (*cautils.ContainerImageIdentifier, error) {
	projectName, repositoryName, referenceName, tag, tagSpecified, err := h.ParseImageURL(imageURL)
	if err != nil {
		return nil, err
	}

	// if the user supplied an imageURL referencing a tag, use that to retrieve the artifact info
	if tagSpecified {
		referenceName = tag
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

	// determine the last time the image was scanned, and/or if scanning is available for this image
	if scanOverview, ok := artifact.GetPayload().ScanOverview["application/vnd.security.vulnerability.report; version=1.1"]; ok {
		scanStatus.LastScanDate = time.Time(scanOverview.EndTime)

		if scanOverview.CompletePercent != 100 {
			scanStatus.IsScanAvailable = false
		} else {
			scanStatus.IsScanAvailable = true
		}

	} else {
		scanStatus.LastScanDate = time.Time{}
	}

	return &scanStatus, nil

}
