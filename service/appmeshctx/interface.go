// Code generated by internal/generate/main.go. DO NOT EDIT.

package appmeshctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appmesh/appmeshiface"
	"github.com/glassechidna/awsctx"
)

type AppMesh interface {
	CreateGatewayRouteWithContext(ctx context.Context, input *appmesh.CreateGatewayRouteInput, opts ...request.Option) (*appmesh.CreateGatewayRouteOutput, error)
	CreateMeshWithContext(ctx context.Context, input *appmesh.CreateMeshInput, opts ...request.Option) (*appmesh.CreateMeshOutput, error)
	CreateRouteWithContext(ctx context.Context, input *appmesh.CreateRouteInput, opts ...request.Option) (*appmesh.CreateRouteOutput, error)
	CreateVirtualGatewayWithContext(ctx context.Context, input *appmesh.CreateVirtualGatewayInput, opts ...request.Option) (*appmesh.CreateVirtualGatewayOutput, error)
	CreateVirtualNodeWithContext(ctx context.Context, input *appmesh.CreateVirtualNodeInput, opts ...request.Option) (*appmesh.CreateVirtualNodeOutput, error)
	CreateVirtualRouterWithContext(ctx context.Context, input *appmesh.CreateVirtualRouterInput, opts ...request.Option) (*appmesh.CreateVirtualRouterOutput, error)
	CreateVirtualServiceWithContext(ctx context.Context, input *appmesh.CreateVirtualServiceInput, opts ...request.Option) (*appmesh.CreateVirtualServiceOutput, error)
	DeleteGatewayRouteWithContext(ctx context.Context, input *appmesh.DeleteGatewayRouteInput, opts ...request.Option) (*appmesh.DeleteGatewayRouteOutput, error)
	DeleteMeshWithContext(ctx context.Context, input *appmesh.DeleteMeshInput, opts ...request.Option) (*appmesh.DeleteMeshOutput, error)
	DeleteRouteWithContext(ctx context.Context, input *appmesh.DeleteRouteInput, opts ...request.Option) (*appmesh.DeleteRouteOutput, error)
	DeleteVirtualGatewayWithContext(ctx context.Context, input *appmesh.DeleteVirtualGatewayInput, opts ...request.Option) (*appmesh.DeleteVirtualGatewayOutput, error)
	DeleteVirtualNodeWithContext(ctx context.Context, input *appmesh.DeleteVirtualNodeInput, opts ...request.Option) (*appmesh.DeleteVirtualNodeOutput, error)
	DeleteVirtualRouterWithContext(ctx context.Context, input *appmesh.DeleteVirtualRouterInput, opts ...request.Option) (*appmesh.DeleteVirtualRouterOutput, error)
	DeleteVirtualServiceWithContext(ctx context.Context, input *appmesh.DeleteVirtualServiceInput, opts ...request.Option) (*appmesh.DeleteVirtualServiceOutput, error)
	DescribeGatewayRouteWithContext(ctx context.Context, input *appmesh.DescribeGatewayRouteInput, opts ...request.Option) (*appmesh.DescribeGatewayRouteOutput, error)
	DescribeMeshWithContext(ctx context.Context, input *appmesh.DescribeMeshInput, opts ...request.Option) (*appmesh.DescribeMeshOutput, error)
	DescribeRouteWithContext(ctx context.Context, input *appmesh.DescribeRouteInput, opts ...request.Option) (*appmesh.DescribeRouteOutput, error)
	DescribeVirtualGatewayWithContext(ctx context.Context, input *appmesh.DescribeVirtualGatewayInput, opts ...request.Option) (*appmesh.DescribeVirtualGatewayOutput, error)
	DescribeVirtualNodeWithContext(ctx context.Context, input *appmesh.DescribeVirtualNodeInput, opts ...request.Option) (*appmesh.DescribeVirtualNodeOutput, error)
	DescribeVirtualRouterWithContext(ctx context.Context, input *appmesh.DescribeVirtualRouterInput, opts ...request.Option) (*appmesh.DescribeVirtualRouterOutput, error)
	DescribeVirtualServiceWithContext(ctx context.Context, input *appmesh.DescribeVirtualServiceInput, opts ...request.Option) (*appmesh.DescribeVirtualServiceOutput, error)
	ListGatewayRoutesWithContext(ctx context.Context, input *appmesh.ListGatewayRoutesInput, opts ...request.Option) (*appmesh.ListGatewayRoutesOutput, error)
	ListGatewayRoutesPagesWithContext(ctx context.Context, input *appmesh.ListGatewayRoutesInput, cb func(*appmesh.ListGatewayRoutesOutput, bool) bool, opts ...request.Option) error
	ListMeshesWithContext(ctx context.Context, input *appmesh.ListMeshesInput, opts ...request.Option) (*appmesh.ListMeshesOutput, error)
	ListMeshesPagesWithContext(ctx context.Context, input *appmesh.ListMeshesInput, cb func(*appmesh.ListMeshesOutput, bool) bool, opts ...request.Option) error
	ListRoutesWithContext(ctx context.Context, input *appmesh.ListRoutesInput, opts ...request.Option) (*appmesh.ListRoutesOutput, error)
	ListRoutesPagesWithContext(ctx context.Context, input *appmesh.ListRoutesInput, cb func(*appmesh.ListRoutesOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *appmesh.ListTagsForResourceInput, opts ...request.Option) (*appmesh.ListTagsForResourceOutput, error)
	ListTagsForResourcePagesWithContext(ctx context.Context, input *appmesh.ListTagsForResourceInput, cb func(*appmesh.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error
	ListVirtualGatewaysWithContext(ctx context.Context, input *appmesh.ListVirtualGatewaysInput, opts ...request.Option) (*appmesh.ListVirtualGatewaysOutput, error)
	ListVirtualGatewaysPagesWithContext(ctx context.Context, input *appmesh.ListVirtualGatewaysInput, cb func(*appmesh.ListVirtualGatewaysOutput, bool) bool, opts ...request.Option) error
	ListVirtualNodesWithContext(ctx context.Context, input *appmesh.ListVirtualNodesInput, opts ...request.Option) (*appmesh.ListVirtualNodesOutput, error)
	ListVirtualNodesPagesWithContext(ctx context.Context, input *appmesh.ListVirtualNodesInput, cb func(*appmesh.ListVirtualNodesOutput, bool) bool, opts ...request.Option) error
	ListVirtualRoutersWithContext(ctx context.Context, input *appmesh.ListVirtualRoutersInput, opts ...request.Option) (*appmesh.ListVirtualRoutersOutput, error)
	ListVirtualRoutersPagesWithContext(ctx context.Context, input *appmesh.ListVirtualRoutersInput, cb func(*appmesh.ListVirtualRoutersOutput, bool) bool, opts ...request.Option) error
	ListVirtualServicesWithContext(ctx context.Context, input *appmesh.ListVirtualServicesInput, opts ...request.Option) (*appmesh.ListVirtualServicesOutput, error)
	ListVirtualServicesPagesWithContext(ctx context.Context, input *appmesh.ListVirtualServicesInput, cb func(*appmesh.ListVirtualServicesOutput, bool) bool, opts ...request.Option) error
	TagResourceWithContext(ctx context.Context, input *appmesh.TagResourceInput, opts ...request.Option) (*appmesh.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *appmesh.UntagResourceInput, opts ...request.Option) (*appmesh.UntagResourceOutput, error)
	UpdateGatewayRouteWithContext(ctx context.Context, input *appmesh.UpdateGatewayRouteInput, opts ...request.Option) (*appmesh.UpdateGatewayRouteOutput, error)
	UpdateMeshWithContext(ctx context.Context, input *appmesh.UpdateMeshInput, opts ...request.Option) (*appmesh.UpdateMeshOutput, error)
	UpdateRouteWithContext(ctx context.Context, input *appmesh.UpdateRouteInput, opts ...request.Option) (*appmesh.UpdateRouteOutput, error)
	UpdateVirtualGatewayWithContext(ctx context.Context, input *appmesh.UpdateVirtualGatewayInput, opts ...request.Option) (*appmesh.UpdateVirtualGatewayOutput, error)
	UpdateVirtualNodeWithContext(ctx context.Context, input *appmesh.UpdateVirtualNodeInput, opts ...request.Option) (*appmesh.UpdateVirtualNodeOutput, error)
	UpdateVirtualRouterWithContext(ctx context.Context, input *appmesh.UpdateVirtualRouterInput, opts ...request.Option) (*appmesh.UpdateVirtualRouterOutput, error)
	UpdateVirtualServiceWithContext(ctx context.Context, input *appmesh.UpdateVirtualServiceInput, opts ...request.Option) (*appmesh.UpdateVirtualServiceOutput, error)
}

type Client struct {
	appmeshiface.AppMeshAPI
	Contexter awsctx.Contexter
}

func New(base appmeshiface.AppMeshAPI, ctxer awsctx.Contexter) AppMesh {
	return &Client{
		AppMeshAPI: base,
		Contexter: ctxer,
	}
}

var _ AppMesh = (*appmesh.AppMesh)(nil)
var _ AppMesh = (*Client)(nil)

func (c *Client) CreateGatewayRouteWithContext(ctx context.Context, input *appmesh.CreateGatewayRouteInput, opts ...request.Option) (*appmesh.CreateGatewayRouteOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "CreateGatewayRoute",
		Input:   input,
		Output:  (*appmesh.CreateGatewayRouteOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.CreateGatewayRouteWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.CreateGatewayRouteOutput), req.Error
}

func (c *Client) CreateMeshWithContext(ctx context.Context, input *appmesh.CreateMeshInput, opts ...request.Option) (*appmesh.CreateMeshOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "CreateMesh",
		Input:   input,
		Output:  (*appmesh.CreateMeshOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.CreateMeshWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.CreateMeshOutput), req.Error
}

func (c *Client) CreateRouteWithContext(ctx context.Context, input *appmesh.CreateRouteInput, opts ...request.Option) (*appmesh.CreateRouteOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "CreateRoute",
		Input:   input,
		Output:  (*appmesh.CreateRouteOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.CreateRouteWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.CreateRouteOutput), req.Error
}

func (c *Client) CreateVirtualGatewayWithContext(ctx context.Context, input *appmesh.CreateVirtualGatewayInput, opts ...request.Option) (*appmesh.CreateVirtualGatewayOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "CreateVirtualGateway",
		Input:   input,
		Output:  (*appmesh.CreateVirtualGatewayOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.CreateVirtualGatewayWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.CreateVirtualGatewayOutput), req.Error
}

func (c *Client) CreateVirtualNodeWithContext(ctx context.Context, input *appmesh.CreateVirtualNodeInput, opts ...request.Option) (*appmesh.CreateVirtualNodeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "CreateVirtualNode",
		Input:   input,
		Output:  (*appmesh.CreateVirtualNodeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.CreateVirtualNodeWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.CreateVirtualNodeOutput), req.Error
}

func (c *Client) CreateVirtualRouterWithContext(ctx context.Context, input *appmesh.CreateVirtualRouterInput, opts ...request.Option) (*appmesh.CreateVirtualRouterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "CreateVirtualRouter",
		Input:   input,
		Output:  (*appmesh.CreateVirtualRouterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.CreateVirtualRouterWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.CreateVirtualRouterOutput), req.Error
}

func (c *Client) CreateVirtualServiceWithContext(ctx context.Context, input *appmesh.CreateVirtualServiceInput, opts ...request.Option) (*appmesh.CreateVirtualServiceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "CreateVirtualService",
		Input:   input,
		Output:  (*appmesh.CreateVirtualServiceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.CreateVirtualServiceWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.CreateVirtualServiceOutput), req.Error
}

func (c *Client) DeleteGatewayRouteWithContext(ctx context.Context, input *appmesh.DeleteGatewayRouteInput, opts ...request.Option) (*appmesh.DeleteGatewayRouteOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DeleteGatewayRoute",
		Input:   input,
		Output:  (*appmesh.DeleteGatewayRouteOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DeleteGatewayRouteWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DeleteGatewayRouteOutput), req.Error
}

func (c *Client) DeleteMeshWithContext(ctx context.Context, input *appmesh.DeleteMeshInput, opts ...request.Option) (*appmesh.DeleteMeshOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DeleteMesh",
		Input:   input,
		Output:  (*appmesh.DeleteMeshOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DeleteMeshWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DeleteMeshOutput), req.Error
}

func (c *Client) DeleteRouteWithContext(ctx context.Context, input *appmesh.DeleteRouteInput, opts ...request.Option) (*appmesh.DeleteRouteOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DeleteRoute",
		Input:   input,
		Output:  (*appmesh.DeleteRouteOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DeleteRouteWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DeleteRouteOutput), req.Error
}

func (c *Client) DeleteVirtualGatewayWithContext(ctx context.Context, input *appmesh.DeleteVirtualGatewayInput, opts ...request.Option) (*appmesh.DeleteVirtualGatewayOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DeleteVirtualGateway",
		Input:   input,
		Output:  (*appmesh.DeleteVirtualGatewayOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DeleteVirtualGatewayWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DeleteVirtualGatewayOutput), req.Error
}

func (c *Client) DeleteVirtualNodeWithContext(ctx context.Context, input *appmesh.DeleteVirtualNodeInput, opts ...request.Option) (*appmesh.DeleteVirtualNodeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DeleteVirtualNode",
		Input:   input,
		Output:  (*appmesh.DeleteVirtualNodeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DeleteVirtualNodeWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DeleteVirtualNodeOutput), req.Error
}

func (c *Client) DeleteVirtualRouterWithContext(ctx context.Context, input *appmesh.DeleteVirtualRouterInput, opts ...request.Option) (*appmesh.DeleteVirtualRouterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DeleteVirtualRouter",
		Input:   input,
		Output:  (*appmesh.DeleteVirtualRouterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DeleteVirtualRouterWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DeleteVirtualRouterOutput), req.Error
}

func (c *Client) DeleteVirtualServiceWithContext(ctx context.Context, input *appmesh.DeleteVirtualServiceInput, opts ...request.Option) (*appmesh.DeleteVirtualServiceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DeleteVirtualService",
		Input:   input,
		Output:  (*appmesh.DeleteVirtualServiceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DeleteVirtualServiceWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DeleteVirtualServiceOutput), req.Error
}

func (c *Client) DescribeGatewayRouteWithContext(ctx context.Context, input *appmesh.DescribeGatewayRouteInput, opts ...request.Option) (*appmesh.DescribeGatewayRouteOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DescribeGatewayRoute",
		Input:   input,
		Output:  (*appmesh.DescribeGatewayRouteOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DescribeGatewayRouteWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DescribeGatewayRouteOutput), req.Error
}

func (c *Client) DescribeMeshWithContext(ctx context.Context, input *appmesh.DescribeMeshInput, opts ...request.Option) (*appmesh.DescribeMeshOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DescribeMesh",
		Input:   input,
		Output:  (*appmesh.DescribeMeshOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DescribeMeshWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DescribeMeshOutput), req.Error
}

func (c *Client) DescribeRouteWithContext(ctx context.Context, input *appmesh.DescribeRouteInput, opts ...request.Option) (*appmesh.DescribeRouteOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DescribeRoute",
		Input:   input,
		Output:  (*appmesh.DescribeRouteOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DescribeRouteWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DescribeRouteOutput), req.Error
}

func (c *Client) DescribeVirtualGatewayWithContext(ctx context.Context, input *appmesh.DescribeVirtualGatewayInput, opts ...request.Option) (*appmesh.DescribeVirtualGatewayOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DescribeVirtualGateway",
		Input:   input,
		Output:  (*appmesh.DescribeVirtualGatewayOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DescribeVirtualGatewayWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DescribeVirtualGatewayOutput), req.Error
}

func (c *Client) DescribeVirtualNodeWithContext(ctx context.Context, input *appmesh.DescribeVirtualNodeInput, opts ...request.Option) (*appmesh.DescribeVirtualNodeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DescribeVirtualNode",
		Input:   input,
		Output:  (*appmesh.DescribeVirtualNodeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DescribeVirtualNodeWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DescribeVirtualNodeOutput), req.Error
}

func (c *Client) DescribeVirtualRouterWithContext(ctx context.Context, input *appmesh.DescribeVirtualRouterInput, opts ...request.Option) (*appmesh.DescribeVirtualRouterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DescribeVirtualRouter",
		Input:   input,
		Output:  (*appmesh.DescribeVirtualRouterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DescribeVirtualRouterWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DescribeVirtualRouterOutput), req.Error
}

func (c *Client) DescribeVirtualServiceWithContext(ctx context.Context, input *appmesh.DescribeVirtualServiceInput, opts ...request.Option) (*appmesh.DescribeVirtualServiceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "DescribeVirtualService",
		Input:   input,
		Output:  (*appmesh.DescribeVirtualServiceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.DescribeVirtualServiceWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.DescribeVirtualServiceOutput), req.Error
}

func (c *Client) ListGatewayRoutesWithContext(ctx context.Context, input *appmesh.ListGatewayRoutesInput, opts ...request.Option) (*appmesh.ListGatewayRoutesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListGatewayRoutes",
		Input:   input,
		Output:  (*appmesh.ListGatewayRoutesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.ListGatewayRoutesWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.ListGatewayRoutesOutput), req.Error
}

func (c *Client) ListGatewayRoutesPagesWithContext(ctx context.Context, input *appmesh.ListGatewayRoutesInput, cb func(*appmesh.ListGatewayRoutesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListGatewayRoutes",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AppMeshAPI.ListGatewayRoutesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListMeshesWithContext(ctx context.Context, input *appmesh.ListMeshesInput, opts ...request.Option) (*appmesh.ListMeshesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListMeshes",
		Input:   input,
		Output:  (*appmesh.ListMeshesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.ListMeshesWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.ListMeshesOutput), req.Error
}

func (c *Client) ListMeshesPagesWithContext(ctx context.Context, input *appmesh.ListMeshesInput, cb func(*appmesh.ListMeshesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListMeshes",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AppMeshAPI.ListMeshesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListRoutesWithContext(ctx context.Context, input *appmesh.ListRoutesInput, opts ...request.Option) (*appmesh.ListRoutesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListRoutes",
		Input:   input,
		Output:  (*appmesh.ListRoutesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.ListRoutesWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.ListRoutesOutput), req.Error
}

func (c *Client) ListRoutesPagesWithContext(ctx context.Context, input *appmesh.ListRoutesInput, cb func(*appmesh.ListRoutesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListRoutes",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AppMeshAPI.ListRoutesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *appmesh.ListTagsForResourceInput, opts ...request.Option) (*appmesh.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*appmesh.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTagsForResourcePagesWithContext(ctx context.Context, input *appmesh.ListTagsForResourceInput, cb func(*appmesh.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListTagsForResource",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AppMeshAPI.ListTagsForResourcePagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListVirtualGatewaysWithContext(ctx context.Context, input *appmesh.ListVirtualGatewaysInput, opts ...request.Option) (*appmesh.ListVirtualGatewaysOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListVirtualGateways",
		Input:   input,
		Output:  (*appmesh.ListVirtualGatewaysOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.ListVirtualGatewaysWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.ListVirtualGatewaysOutput), req.Error
}

func (c *Client) ListVirtualGatewaysPagesWithContext(ctx context.Context, input *appmesh.ListVirtualGatewaysInput, cb func(*appmesh.ListVirtualGatewaysOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListVirtualGateways",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AppMeshAPI.ListVirtualGatewaysPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListVirtualNodesWithContext(ctx context.Context, input *appmesh.ListVirtualNodesInput, opts ...request.Option) (*appmesh.ListVirtualNodesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListVirtualNodes",
		Input:   input,
		Output:  (*appmesh.ListVirtualNodesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.ListVirtualNodesWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.ListVirtualNodesOutput), req.Error
}

func (c *Client) ListVirtualNodesPagesWithContext(ctx context.Context, input *appmesh.ListVirtualNodesInput, cb func(*appmesh.ListVirtualNodesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListVirtualNodes",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AppMeshAPI.ListVirtualNodesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListVirtualRoutersWithContext(ctx context.Context, input *appmesh.ListVirtualRoutersInput, opts ...request.Option) (*appmesh.ListVirtualRoutersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListVirtualRouters",
		Input:   input,
		Output:  (*appmesh.ListVirtualRoutersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.ListVirtualRoutersWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.ListVirtualRoutersOutput), req.Error
}

func (c *Client) ListVirtualRoutersPagesWithContext(ctx context.Context, input *appmesh.ListVirtualRoutersInput, cb func(*appmesh.ListVirtualRoutersOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListVirtualRouters",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AppMeshAPI.ListVirtualRoutersPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListVirtualServicesWithContext(ctx context.Context, input *appmesh.ListVirtualServicesInput, opts ...request.Option) (*appmesh.ListVirtualServicesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListVirtualServices",
		Input:   input,
		Output:  (*appmesh.ListVirtualServicesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.ListVirtualServicesWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.ListVirtualServicesOutput), req.Error
}

func (c *Client) ListVirtualServicesPagesWithContext(ctx context.Context, input *appmesh.ListVirtualServicesInput, cb func(*appmesh.ListVirtualServicesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "ListVirtualServices",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AppMeshAPI.ListVirtualServicesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *appmesh.TagResourceInput, opts ...request.Option) (*appmesh.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "TagResource",
		Input:   input,
		Output:  (*appmesh.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *appmesh.UntagResourceInput, opts ...request.Option) (*appmesh.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*appmesh.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.UntagResourceOutput), req.Error
}

func (c *Client) UpdateGatewayRouteWithContext(ctx context.Context, input *appmesh.UpdateGatewayRouteInput, opts ...request.Option) (*appmesh.UpdateGatewayRouteOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "UpdateGatewayRoute",
		Input:   input,
		Output:  (*appmesh.UpdateGatewayRouteOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.UpdateGatewayRouteWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.UpdateGatewayRouteOutput), req.Error
}

func (c *Client) UpdateMeshWithContext(ctx context.Context, input *appmesh.UpdateMeshInput, opts ...request.Option) (*appmesh.UpdateMeshOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "UpdateMesh",
		Input:   input,
		Output:  (*appmesh.UpdateMeshOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.UpdateMeshWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.UpdateMeshOutput), req.Error
}

func (c *Client) UpdateRouteWithContext(ctx context.Context, input *appmesh.UpdateRouteInput, opts ...request.Option) (*appmesh.UpdateRouteOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "UpdateRoute",
		Input:   input,
		Output:  (*appmesh.UpdateRouteOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.UpdateRouteWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.UpdateRouteOutput), req.Error
}

func (c *Client) UpdateVirtualGatewayWithContext(ctx context.Context, input *appmesh.UpdateVirtualGatewayInput, opts ...request.Option) (*appmesh.UpdateVirtualGatewayOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "UpdateVirtualGateway",
		Input:   input,
		Output:  (*appmesh.UpdateVirtualGatewayOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.UpdateVirtualGatewayWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.UpdateVirtualGatewayOutput), req.Error
}

func (c *Client) UpdateVirtualNodeWithContext(ctx context.Context, input *appmesh.UpdateVirtualNodeInput, opts ...request.Option) (*appmesh.UpdateVirtualNodeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "UpdateVirtualNode",
		Input:   input,
		Output:  (*appmesh.UpdateVirtualNodeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.UpdateVirtualNodeWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.UpdateVirtualNodeOutput), req.Error
}

func (c *Client) UpdateVirtualRouterWithContext(ctx context.Context, input *appmesh.UpdateVirtualRouterInput, opts ...request.Option) (*appmesh.UpdateVirtualRouterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "UpdateVirtualRouter",
		Input:   input,
		Output:  (*appmesh.UpdateVirtualRouterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.UpdateVirtualRouterWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.UpdateVirtualRouterOutput), req.Error
}

func (c *Client) UpdateVirtualServiceWithContext(ctx context.Context, input *appmesh.UpdateVirtualServiceInput, opts ...request.Option) (*appmesh.UpdateVirtualServiceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "appmesh",
		Action:  "UpdateVirtualService",
		Input:   input,
		Output:  (*appmesh.UpdateVirtualServiceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AppMeshAPI.UpdateVirtualServiceWithContext(ctx, input, opts...)
	})

	return req.Output.(*appmesh.UpdateVirtualServiceOutput), req.Error
}
