package harbor

import (
	"context"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/scan"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/scan_all"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// ScanImage sends a scan request for the specified image
func (h harborRegistry) ScanImage(projectName, repositoryName, referenceName string) (*scan.ScanArtifactAccepted, error) {

	params := scan.ScanArtifactParams{
		ProjectName:    projectName,
		Reference:      referenceName,
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
