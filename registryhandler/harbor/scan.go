package harbor

import (
	"context"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/scan"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/scan_all"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"strings"
)

// ScanImage sends a scan request for the specified image
func (h harborRegistry) ScanImage(imageURL string) (*scan.ScanArtifactAccepted, error) {

	projectName, repositoryName, reference, err := h.ParseImageURL(imageURL)
	if err != nil {
		return nil, err
	}

	params := scan.ScanArtifactParams{
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		Context:        context.Background(),
	}

	result, err := h.client.V2().Scan.ScanArtifact(context.Background(), &params)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// ScanAll sends a manual request to scan all images in the registry
func (h harborRegistry) ScanAll() (*scan_all.CreateScanAllScheduleCreated, error) {

	params := scan_all.CreateScanAllScheduleParams{
		Schedule: &models.Schedule{Schedule: &models.ScheduleObj{Type: models.ScheduleObjTypeManual}},
	}

	schedule, err := h.client.V2().ScanAll.CreateScanAllSchedule(context.Background(), &params)
	if err != nil {
		return nil, err
	}

	return schedule, err
}

func (h harborRegistry) ParseImageURL(imageURL string) (string, string, string, error) {

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
