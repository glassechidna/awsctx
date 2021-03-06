// Code generated by internal/generate/main.go. DO NOT EDIT.

package iotthingsgraphctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph/iotthingsgraphiface"
	"github.com/glassechidna/awsctx"
)

type IoTThingsGraph interface {
	AssociateEntityToThingWithContext(ctx context.Context, input *iotthingsgraph.AssociateEntityToThingInput, opts ...request.Option) (*iotthingsgraph.AssociateEntityToThingOutput, error)
	CreateFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.CreateFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.CreateFlowTemplateOutput, error)
	CreateSystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.CreateSystemInstanceInput, opts ...request.Option) (*iotthingsgraph.CreateSystemInstanceOutput, error)
	CreateSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.CreateSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.CreateSystemTemplateOutput, error)
	DeleteFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.DeleteFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.DeleteFlowTemplateOutput, error)
	DeleteNamespaceWithContext(ctx context.Context, input *iotthingsgraph.DeleteNamespaceInput, opts ...request.Option) (*iotthingsgraph.DeleteNamespaceOutput, error)
	DeleteSystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.DeleteSystemInstanceInput, opts ...request.Option) (*iotthingsgraph.DeleteSystemInstanceOutput, error)
	DeleteSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.DeleteSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.DeleteSystemTemplateOutput, error)
	DeploySystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.DeploySystemInstanceInput, opts ...request.Option) (*iotthingsgraph.DeploySystemInstanceOutput, error)
	DeprecateFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.DeprecateFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.DeprecateFlowTemplateOutput, error)
	DeprecateSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.DeprecateSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.DeprecateSystemTemplateOutput, error)
	DescribeNamespaceWithContext(ctx context.Context, input *iotthingsgraph.DescribeNamespaceInput, opts ...request.Option) (*iotthingsgraph.DescribeNamespaceOutput, error)
	DissociateEntityFromThingWithContext(ctx context.Context, input *iotthingsgraph.DissociateEntityFromThingInput, opts ...request.Option) (*iotthingsgraph.DissociateEntityFromThingOutput, error)
	GetEntitiesWithContext(ctx context.Context, input *iotthingsgraph.GetEntitiesInput, opts ...request.Option) (*iotthingsgraph.GetEntitiesOutput, error)
	GetFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.GetFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.GetFlowTemplateOutput, error)
	GetFlowTemplateRevisionsWithContext(ctx context.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput, opts ...request.Option) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error)
	GetFlowTemplateRevisionsPagesWithContext(ctx context.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput, cb func(*iotthingsgraph.GetFlowTemplateRevisionsOutput, bool) bool, opts ...request.Option) error
	GetNamespaceDeletionStatusWithContext(ctx context.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput, opts ...request.Option) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error)
	GetSystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.GetSystemInstanceInput, opts ...request.Option) (*iotthingsgraph.GetSystemInstanceOutput, error)
	GetSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.GetSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.GetSystemTemplateOutput, error)
	GetSystemTemplateRevisionsWithContext(ctx context.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput, opts ...request.Option) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error)
	GetSystemTemplateRevisionsPagesWithContext(ctx context.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput, cb func(*iotthingsgraph.GetSystemTemplateRevisionsOutput, bool) bool, opts ...request.Option) error
	GetUploadStatusWithContext(ctx context.Context, input *iotthingsgraph.GetUploadStatusInput, opts ...request.Option) (*iotthingsgraph.GetUploadStatusOutput, error)
	ListFlowExecutionMessagesWithContext(ctx context.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput, opts ...request.Option) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error)
	ListFlowExecutionMessagesPagesWithContext(ctx context.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput, cb func(*iotthingsgraph.ListFlowExecutionMessagesOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *iotthingsgraph.ListTagsForResourceInput, opts ...request.Option) (*iotthingsgraph.ListTagsForResourceOutput, error)
	ListTagsForResourcePagesWithContext(ctx context.Context, input *iotthingsgraph.ListTagsForResourceInput, cb func(*iotthingsgraph.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error
	SearchEntitiesWithContext(ctx context.Context, input *iotthingsgraph.SearchEntitiesInput, opts ...request.Option) (*iotthingsgraph.SearchEntitiesOutput, error)
	SearchEntitiesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchEntitiesInput, cb func(*iotthingsgraph.SearchEntitiesOutput, bool) bool, opts ...request.Option) error
	SearchFlowExecutionsWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowExecutionsInput, opts ...request.Option) (*iotthingsgraph.SearchFlowExecutionsOutput, error)
	SearchFlowExecutionsPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowExecutionsInput, cb func(*iotthingsgraph.SearchFlowExecutionsOutput, bool) bool, opts ...request.Option) error
	SearchFlowTemplatesWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowTemplatesInput, opts ...request.Option) (*iotthingsgraph.SearchFlowTemplatesOutput, error)
	SearchFlowTemplatesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowTemplatesInput, cb func(*iotthingsgraph.SearchFlowTemplatesOutput, bool) bool, opts ...request.Option) error
	SearchSystemInstancesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemInstancesInput, opts ...request.Option) (*iotthingsgraph.SearchSystemInstancesOutput, error)
	SearchSystemInstancesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemInstancesInput, cb func(*iotthingsgraph.SearchSystemInstancesOutput, bool) bool, opts ...request.Option) error
	SearchSystemTemplatesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemTemplatesInput, opts ...request.Option) (*iotthingsgraph.SearchSystemTemplatesOutput, error)
	SearchSystemTemplatesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemTemplatesInput, cb func(*iotthingsgraph.SearchSystemTemplatesOutput, bool) bool, opts ...request.Option) error
	SearchThingsWithContext(ctx context.Context, input *iotthingsgraph.SearchThingsInput, opts ...request.Option) (*iotthingsgraph.SearchThingsOutput, error)
	SearchThingsPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchThingsInput, cb func(*iotthingsgraph.SearchThingsOutput, bool) bool, opts ...request.Option) error
	TagResourceWithContext(ctx context.Context, input *iotthingsgraph.TagResourceInput, opts ...request.Option) (*iotthingsgraph.TagResourceOutput, error)
	UndeploySystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.UndeploySystemInstanceInput, opts ...request.Option) (*iotthingsgraph.UndeploySystemInstanceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *iotthingsgraph.UntagResourceInput, opts ...request.Option) (*iotthingsgraph.UntagResourceOutput, error)
	UpdateFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.UpdateFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.UpdateFlowTemplateOutput, error)
	UpdateSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.UpdateSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.UpdateSystemTemplateOutput, error)
	UploadEntityDefinitionsWithContext(ctx context.Context, input *iotthingsgraph.UploadEntityDefinitionsInput, opts ...request.Option) (*iotthingsgraph.UploadEntityDefinitionsOutput, error)
}

type Client struct {
	iotthingsgraphiface.IoTThingsGraphAPI
	Contexter awsctx.Contexter
}

func New(base iotthingsgraphiface.IoTThingsGraphAPI, ctxer awsctx.Contexter) IoTThingsGraph {
	return &Client{
		IoTThingsGraphAPI: base,
		Contexter: ctxer,
	}
}

var _ IoTThingsGraph = (*iotthingsgraph.IoTThingsGraph)(nil)
var _ IoTThingsGraph = (*Client)(nil)

func (c *Client) AssociateEntityToThingWithContext(ctx context.Context, input *iotthingsgraph.AssociateEntityToThingInput, opts ...request.Option) (*iotthingsgraph.AssociateEntityToThingOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "AssociateEntityToThing",
		Input:   input,
		Output:  (*iotthingsgraph.AssociateEntityToThingOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.AssociateEntityToThingWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.AssociateEntityToThingOutput), req.Error
}

func (c *Client) CreateFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.CreateFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.CreateFlowTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "CreateFlowTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.CreateFlowTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.CreateFlowTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.CreateFlowTemplateOutput), req.Error
}

func (c *Client) CreateSystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.CreateSystemInstanceInput, opts ...request.Option) (*iotthingsgraph.CreateSystemInstanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "CreateSystemInstance",
		Input:   input,
		Output:  (*iotthingsgraph.CreateSystemInstanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.CreateSystemInstanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.CreateSystemInstanceOutput), req.Error
}

func (c *Client) CreateSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.CreateSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.CreateSystemTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "CreateSystemTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.CreateSystemTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.CreateSystemTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.CreateSystemTemplateOutput), req.Error
}

func (c *Client) DeleteFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.DeleteFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.DeleteFlowTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "DeleteFlowTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.DeleteFlowTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.DeleteFlowTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.DeleteFlowTemplateOutput), req.Error
}

func (c *Client) DeleteNamespaceWithContext(ctx context.Context, input *iotthingsgraph.DeleteNamespaceInput, opts ...request.Option) (*iotthingsgraph.DeleteNamespaceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "DeleteNamespace",
		Input:   input,
		Output:  (*iotthingsgraph.DeleteNamespaceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.DeleteNamespaceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.DeleteNamespaceOutput), req.Error
}

func (c *Client) DeleteSystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.DeleteSystemInstanceInput, opts ...request.Option) (*iotthingsgraph.DeleteSystemInstanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "DeleteSystemInstance",
		Input:   input,
		Output:  (*iotthingsgraph.DeleteSystemInstanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.DeleteSystemInstanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.DeleteSystemInstanceOutput), req.Error
}

func (c *Client) DeleteSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.DeleteSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.DeleteSystemTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "DeleteSystemTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.DeleteSystemTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.DeleteSystemTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.DeleteSystemTemplateOutput), req.Error
}

func (c *Client) DeploySystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.DeploySystemInstanceInput, opts ...request.Option) (*iotthingsgraph.DeploySystemInstanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "DeploySystemInstance",
		Input:   input,
		Output:  (*iotthingsgraph.DeploySystemInstanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.DeploySystemInstanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.DeploySystemInstanceOutput), req.Error
}

func (c *Client) DeprecateFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.DeprecateFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.DeprecateFlowTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "DeprecateFlowTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.DeprecateFlowTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.DeprecateFlowTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.DeprecateFlowTemplateOutput), req.Error
}

func (c *Client) DeprecateSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.DeprecateSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.DeprecateSystemTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "DeprecateSystemTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.DeprecateSystemTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.DeprecateSystemTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.DeprecateSystemTemplateOutput), req.Error
}

func (c *Client) DescribeNamespaceWithContext(ctx context.Context, input *iotthingsgraph.DescribeNamespaceInput, opts ...request.Option) (*iotthingsgraph.DescribeNamespaceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "DescribeNamespace",
		Input:   input,
		Output:  (*iotthingsgraph.DescribeNamespaceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.DescribeNamespaceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.DescribeNamespaceOutput), req.Error
}

func (c *Client) DissociateEntityFromThingWithContext(ctx context.Context, input *iotthingsgraph.DissociateEntityFromThingInput, opts ...request.Option) (*iotthingsgraph.DissociateEntityFromThingOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "DissociateEntityFromThing",
		Input:   input,
		Output:  (*iotthingsgraph.DissociateEntityFromThingOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.DissociateEntityFromThingWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.DissociateEntityFromThingOutput), req.Error
}

func (c *Client) GetEntitiesWithContext(ctx context.Context, input *iotthingsgraph.GetEntitiesInput, opts ...request.Option) (*iotthingsgraph.GetEntitiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetEntities",
		Input:   input,
		Output:  (*iotthingsgraph.GetEntitiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.GetEntitiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.GetEntitiesOutput), req.Error
}

func (c *Client) GetFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.GetFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.GetFlowTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetFlowTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.GetFlowTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.GetFlowTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.GetFlowTemplateOutput), req.Error
}

func (c *Client) GetFlowTemplateRevisionsWithContext(ctx context.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput, opts ...request.Option) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetFlowTemplateRevisions",
		Input:   input,
		Output:  (*iotthingsgraph.GetFlowTemplateRevisionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.GetFlowTemplateRevisionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.GetFlowTemplateRevisionsOutput), req.Error
}

func (c *Client) GetFlowTemplateRevisionsPagesWithContext(ctx context.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput, cb func(*iotthingsgraph.GetFlowTemplateRevisionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetFlowTemplateRevisions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.GetFlowTemplateRevisionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetNamespaceDeletionStatusWithContext(ctx context.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput, opts ...request.Option) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetNamespaceDeletionStatus",
		Input:   input,
		Output:  (*iotthingsgraph.GetNamespaceDeletionStatusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.GetNamespaceDeletionStatusWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.GetNamespaceDeletionStatusOutput), req.Error
}

func (c *Client) GetSystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.GetSystemInstanceInput, opts ...request.Option) (*iotthingsgraph.GetSystemInstanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetSystemInstance",
		Input:   input,
		Output:  (*iotthingsgraph.GetSystemInstanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.GetSystemInstanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.GetSystemInstanceOutput), req.Error
}

func (c *Client) GetSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.GetSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.GetSystemTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetSystemTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.GetSystemTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.GetSystemTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.GetSystemTemplateOutput), req.Error
}

func (c *Client) GetSystemTemplateRevisionsWithContext(ctx context.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput, opts ...request.Option) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetSystemTemplateRevisions",
		Input:   input,
		Output:  (*iotthingsgraph.GetSystemTemplateRevisionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.GetSystemTemplateRevisionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.GetSystemTemplateRevisionsOutput), req.Error
}

func (c *Client) GetSystemTemplateRevisionsPagesWithContext(ctx context.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput, cb func(*iotthingsgraph.GetSystemTemplateRevisionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetSystemTemplateRevisions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.GetSystemTemplateRevisionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetUploadStatusWithContext(ctx context.Context, input *iotthingsgraph.GetUploadStatusInput, opts ...request.Option) (*iotthingsgraph.GetUploadStatusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "GetUploadStatus",
		Input:   input,
		Output:  (*iotthingsgraph.GetUploadStatusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.GetUploadStatusWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.GetUploadStatusOutput), req.Error
}

func (c *Client) ListFlowExecutionMessagesWithContext(ctx context.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput, opts ...request.Option) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "ListFlowExecutionMessages",
		Input:   input,
		Output:  (*iotthingsgraph.ListFlowExecutionMessagesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.ListFlowExecutionMessagesWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.ListFlowExecutionMessagesOutput), req.Error
}

func (c *Client) ListFlowExecutionMessagesPagesWithContext(ctx context.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput, cb func(*iotthingsgraph.ListFlowExecutionMessagesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "ListFlowExecutionMessages",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.ListFlowExecutionMessagesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *iotthingsgraph.ListTagsForResourceInput, opts ...request.Option) (*iotthingsgraph.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*iotthingsgraph.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTagsForResourcePagesWithContext(ctx context.Context, input *iotthingsgraph.ListTagsForResourceInput, cb func(*iotthingsgraph.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "ListTagsForResource",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.ListTagsForResourcePagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) SearchEntitiesWithContext(ctx context.Context, input *iotthingsgraph.SearchEntitiesInput, opts ...request.Option) (*iotthingsgraph.SearchEntitiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchEntities",
		Input:   input,
		Output:  (*iotthingsgraph.SearchEntitiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.SearchEntitiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.SearchEntitiesOutput), req.Error
}

func (c *Client) SearchEntitiesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchEntitiesInput, cb func(*iotthingsgraph.SearchEntitiesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchEntities",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.SearchEntitiesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) SearchFlowExecutionsWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowExecutionsInput, opts ...request.Option) (*iotthingsgraph.SearchFlowExecutionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchFlowExecutions",
		Input:   input,
		Output:  (*iotthingsgraph.SearchFlowExecutionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.SearchFlowExecutionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.SearchFlowExecutionsOutput), req.Error
}

func (c *Client) SearchFlowExecutionsPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowExecutionsInput, cb func(*iotthingsgraph.SearchFlowExecutionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchFlowExecutions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.SearchFlowExecutionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) SearchFlowTemplatesWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowTemplatesInput, opts ...request.Option) (*iotthingsgraph.SearchFlowTemplatesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchFlowTemplates",
		Input:   input,
		Output:  (*iotthingsgraph.SearchFlowTemplatesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.SearchFlowTemplatesWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.SearchFlowTemplatesOutput), req.Error
}

func (c *Client) SearchFlowTemplatesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowTemplatesInput, cb func(*iotthingsgraph.SearchFlowTemplatesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchFlowTemplates",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.SearchFlowTemplatesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) SearchSystemInstancesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemInstancesInput, opts ...request.Option) (*iotthingsgraph.SearchSystemInstancesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchSystemInstances",
		Input:   input,
		Output:  (*iotthingsgraph.SearchSystemInstancesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.SearchSystemInstancesWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.SearchSystemInstancesOutput), req.Error
}

func (c *Client) SearchSystemInstancesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemInstancesInput, cb func(*iotthingsgraph.SearchSystemInstancesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchSystemInstances",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.SearchSystemInstancesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) SearchSystemTemplatesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemTemplatesInput, opts ...request.Option) (*iotthingsgraph.SearchSystemTemplatesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchSystemTemplates",
		Input:   input,
		Output:  (*iotthingsgraph.SearchSystemTemplatesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.SearchSystemTemplatesWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.SearchSystemTemplatesOutput), req.Error
}

func (c *Client) SearchSystemTemplatesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemTemplatesInput, cb func(*iotthingsgraph.SearchSystemTemplatesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchSystemTemplates",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.SearchSystemTemplatesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) SearchThingsWithContext(ctx context.Context, input *iotthingsgraph.SearchThingsInput, opts ...request.Option) (*iotthingsgraph.SearchThingsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchThings",
		Input:   input,
		Output:  (*iotthingsgraph.SearchThingsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.SearchThingsWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.SearchThingsOutput), req.Error
}

func (c *Client) SearchThingsPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchThingsInput, cb func(*iotthingsgraph.SearchThingsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "SearchThings",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.IoTThingsGraphAPI.SearchThingsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *iotthingsgraph.TagResourceInput, opts ...request.Option) (*iotthingsgraph.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "TagResource",
		Input:   input,
		Output:  (*iotthingsgraph.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.TagResourceOutput), req.Error
}

func (c *Client) UndeploySystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.UndeploySystemInstanceInput, opts ...request.Option) (*iotthingsgraph.UndeploySystemInstanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "UndeploySystemInstance",
		Input:   input,
		Output:  (*iotthingsgraph.UndeploySystemInstanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.UndeploySystemInstanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.UndeploySystemInstanceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *iotthingsgraph.UntagResourceInput, opts ...request.Option) (*iotthingsgraph.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*iotthingsgraph.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.UntagResourceOutput), req.Error
}

func (c *Client) UpdateFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.UpdateFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.UpdateFlowTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "UpdateFlowTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.UpdateFlowTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.UpdateFlowTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.UpdateFlowTemplateOutput), req.Error
}

func (c *Client) UpdateSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.UpdateSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.UpdateSystemTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "UpdateSystemTemplate",
		Input:   input,
		Output:  (*iotthingsgraph.UpdateSystemTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.UpdateSystemTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.UpdateSystemTemplateOutput), req.Error
}

func (c *Client) UploadEntityDefinitionsWithContext(ctx context.Context, input *iotthingsgraph.UploadEntityDefinitionsInput, opts ...request.Option) (*iotthingsgraph.UploadEntityDefinitionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iotthingsgraph",
		Action:  "UploadEntityDefinitions",
		Input:   input,
		Output:  (*iotthingsgraph.UploadEntityDefinitionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoTThingsGraphAPI.UploadEntityDefinitionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*iotthingsgraph.UploadEntityDefinitionsOutput), req.Error
}
