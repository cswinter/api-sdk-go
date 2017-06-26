package smartling

import (
	"fmt"
)

const (
	endpointFileDelete = "/files-api/v2/projects/%s/file/delete"
)

func (client *Client) DeleteFile(
	projectID string,
	uri string,
) error {
	request := FileURIRequest{
		FileURI: uri,
	}

	contentType, body, err := request.GetForm()
	if err != nil {
		return fmt.Errorf(
			"failed to create file delete form: %s",
			err,
		)
	}

	_, _, err = client.Post(
		fmt.Sprintf(endpointFileDelete, projectID),
		body,
		nil,
		ContentTypeOption(contentType),
	)
	if err != nil {
		return fmt.Errorf(
			"failed to remove file: %s", err,
		)
	}

	return nil
}