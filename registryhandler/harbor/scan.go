package harbor

import (
	"context"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/scan"
)

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
