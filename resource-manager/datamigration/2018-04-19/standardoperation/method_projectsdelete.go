package standardoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectsDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ProjectsDeleteOperationOptions struct {
	DeleteRunningTasks *bool
}

func DefaultProjectsDeleteOperationOptions() ProjectsDeleteOperationOptions {
	return ProjectsDeleteOperationOptions{}
}

func (o ProjectsDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProjectsDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ProjectsDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DeleteRunningTasks != nil {
		out.Append("deleteRunningTasks", fmt.Sprintf("%v", *o.DeleteRunningTasks))
	}
	return &out
}

// ProjectsDelete ...
func (c StandardOperationClient) ProjectsDelete(ctx context.Context, id ProjectId, options ProjectsDeleteOperationOptions) (result ProjectsDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		Path:          id.ID(),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
