// Code generated by internal/generate/main.go. DO NOT EDIT.

package servicequotasctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/aws/aws-sdk-go/service/servicequotas/servicequotasiface"
	"github.com/glassechidna/awsctx"
)

type ServiceQuotas interface {
	AssociateServiceQuotaTemplateWithContext(ctx context.Context, input *servicequotas.AssociateServiceQuotaTemplateInput, opts ...request.Option) (*servicequotas.AssociateServiceQuotaTemplateOutput, error)
	DeleteServiceQuotaIncreaseRequestFromTemplateWithContext(ctx context.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput, opts ...request.Option) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error)
	DisassociateServiceQuotaTemplateWithContext(ctx context.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput, opts ...request.Option) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error)
	GetAWSDefaultServiceQuotaWithContext(ctx context.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput, opts ...request.Option) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error)
	GetAssociationForServiceQuotaTemplateWithContext(ctx context.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput, opts ...request.Option) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error)
	GetRequestedServiceQuotaChangeWithContext(ctx context.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput, opts ...request.Option) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error)
	GetServiceQuotaWithContext(ctx context.Context, input *servicequotas.GetServiceQuotaInput, opts ...request.Option) (*servicequotas.GetServiceQuotaOutput, error)
	GetServiceQuotaIncreaseRequestFromTemplateWithContext(ctx context.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput, opts ...request.Option) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error)
	ListAWSDefaultServiceQuotasWithContext(ctx context.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput, opts ...request.Option) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error)
	ListAWSDefaultServiceQuotasPagesWithContext(ctx context.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput, cb func(*servicequotas.ListAWSDefaultServiceQuotasOutput, bool) bool, opts ...request.Option) error
	ListRequestedServiceQuotaChangeHistoryWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, opts ...request.Option) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error)
	ListRequestedServiceQuotaChangeHistoryPagesWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, cb func(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, bool) bool, opts ...request.Option) error
	ListRequestedServiceQuotaChangeHistoryByQuotaWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, opts ...request.Option) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error)
	ListRequestedServiceQuotaChangeHistoryByQuotaPagesWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, cb func(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, bool) bool, opts ...request.Option) error
	ListServiceQuotaIncreaseRequestsInTemplateWithContext(ctx context.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, opts ...request.Option) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error)
	ListServiceQuotaIncreaseRequestsInTemplatePagesWithContext(ctx context.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, cb func(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, bool) bool, opts ...request.Option) error
	ListServiceQuotasWithContext(ctx context.Context, input *servicequotas.ListServiceQuotasInput, opts ...request.Option) (*servicequotas.ListServiceQuotasOutput, error)
	ListServiceQuotasPagesWithContext(ctx context.Context, input *servicequotas.ListServiceQuotasInput, cb func(*servicequotas.ListServiceQuotasOutput, bool) bool, opts ...request.Option) error
	ListServicesWithContext(ctx context.Context, input *servicequotas.ListServicesInput, opts ...request.Option) (*servicequotas.ListServicesOutput, error)
	ListServicesPagesWithContext(ctx context.Context, input *servicequotas.ListServicesInput, cb func(*servicequotas.ListServicesOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *servicequotas.ListTagsForResourceInput, opts ...request.Option) (*servicequotas.ListTagsForResourceOutput, error)
	PutServiceQuotaIncreaseRequestIntoTemplateWithContext(ctx context.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput, opts ...request.Option) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error)
	RequestServiceQuotaIncreaseWithContext(ctx context.Context, input *servicequotas.RequestServiceQuotaIncreaseInput, opts ...request.Option) (*servicequotas.RequestServiceQuotaIncreaseOutput, error)
	TagResourceWithContext(ctx context.Context, input *servicequotas.TagResourceInput, opts ...request.Option) (*servicequotas.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *servicequotas.UntagResourceInput, opts ...request.Option) (*servicequotas.UntagResourceOutput, error)
}

type Client struct {
	servicequotasiface.ServiceQuotasAPI
	Contexter awsctx.Contexter
}

func New(base servicequotasiface.ServiceQuotasAPI, ctxer awsctx.Contexter) ServiceQuotas {
	return &Client{
		ServiceQuotasAPI: base,
		Contexter: ctxer,
	}
}

var _ ServiceQuotas = (*servicequotas.ServiceQuotas)(nil)
var _ ServiceQuotas = (*Client)(nil)

func (c *Client) AssociateServiceQuotaTemplateWithContext(ctx context.Context, input *servicequotas.AssociateServiceQuotaTemplateInput, opts ...request.Option) (*servicequotas.AssociateServiceQuotaTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "AssociateServiceQuotaTemplate",
		Input:   input,
		Output:  (*servicequotas.AssociateServiceQuotaTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.AssociateServiceQuotaTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.AssociateServiceQuotaTemplateOutput), req.Error
}

func (c *Client) DeleteServiceQuotaIncreaseRequestFromTemplateWithContext(ctx context.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput, opts ...request.Option) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "DeleteServiceQuotaIncreaseRequestFromTemplate",
		Input:   input,
		Output:  (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.DeleteServiceQuotaIncreaseRequestFromTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput), req.Error
}

func (c *Client) DisassociateServiceQuotaTemplateWithContext(ctx context.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput, opts ...request.Option) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "DisassociateServiceQuotaTemplate",
		Input:   input,
		Output:  (*servicequotas.DisassociateServiceQuotaTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.DisassociateServiceQuotaTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.DisassociateServiceQuotaTemplateOutput), req.Error
}

func (c *Client) GetAWSDefaultServiceQuotaWithContext(ctx context.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput, opts ...request.Option) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "GetAWSDefaultServiceQuota",
		Input:   input,
		Output:  (*servicequotas.GetAWSDefaultServiceQuotaOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.GetAWSDefaultServiceQuotaWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.GetAWSDefaultServiceQuotaOutput), req.Error
}

func (c *Client) GetAssociationForServiceQuotaTemplateWithContext(ctx context.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput, opts ...request.Option) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "GetAssociationForServiceQuotaTemplate",
		Input:   input,
		Output:  (*servicequotas.GetAssociationForServiceQuotaTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.GetAssociationForServiceQuotaTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.GetAssociationForServiceQuotaTemplateOutput), req.Error
}

func (c *Client) GetRequestedServiceQuotaChangeWithContext(ctx context.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput, opts ...request.Option) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "GetRequestedServiceQuotaChange",
		Input:   input,
		Output:  (*servicequotas.GetRequestedServiceQuotaChangeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.GetRequestedServiceQuotaChangeWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.GetRequestedServiceQuotaChangeOutput), req.Error
}

func (c *Client) GetServiceQuotaWithContext(ctx context.Context, input *servicequotas.GetServiceQuotaInput, opts ...request.Option) (*servicequotas.GetServiceQuotaOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "GetServiceQuota",
		Input:   input,
		Output:  (*servicequotas.GetServiceQuotaOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.GetServiceQuotaWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.GetServiceQuotaOutput), req.Error
}

func (c *Client) GetServiceQuotaIncreaseRequestFromTemplateWithContext(ctx context.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput, opts ...request.Option) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "GetServiceQuotaIncreaseRequestFromTemplate",
		Input:   input,
		Output:  (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.GetServiceQuotaIncreaseRequestFromTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput), req.Error
}

func (c *Client) ListAWSDefaultServiceQuotasWithContext(ctx context.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput, opts ...request.Option) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListAWSDefaultServiceQuotas",
		Input:   input,
		Output:  (*servicequotas.ListAWSDefaultServiceQuotasOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.ListAWSDefaultServiceQuotasWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.ListAWSDefaultServiceQuotasOutput), req.Error
}

func (c *Client) ListAWSDefaultServiceQuotasPagesWithContext(ctx context.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput, cb func(*servicequotas.ListAWSDefaultServiceQuotasOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListAWSDefaultServiceQuotas",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ServiceQuotasAPI.ListAWSDefaultServiceQuotasPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListRequestedServiceQuotaChangeHistoryWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, opts ...request.Option) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListRequestedServiceQuotaChangeHistory",
		Input:   input,
		Output:  (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput), req.Error
}

func (c *Client) ListRequestedServiceQuotaChangeHistoryPagesWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, cb func(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListRequestedServiceQuotaChangeHistory",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListRequestedServiceQuotaChangeHistoryByQuotaWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, opts ...request.Option) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListRequestedServiceQuotaChangeHistoryByQuota",
		Input:   input,
		Output:  (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryByQuotaWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput), req.Error
}

func (c *Client) ListRequestedServiceQuotaChangeHistoryByQuotaPagesWithContext(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, cb func(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListRequestedServiceQuotaChangeHistoryByQuota",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ServiceQuotasAPI.ListRequestedServiceQuotaChangeHistoryByQuotaPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListServiceQuotaIncreaseRequestsInTemplateWithContext(ctx context.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, opts ...request.Option) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListServiceQuotaIncreaseRequestsInTemplate",
		Input:   input,
		Output:  (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.ListServiceQuotaIncreaseRequestsInTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput), req.Error
}

func (c *Client) ListServiceQuotaIncreaseRequestsInTemplatePagesWithContext(ctx context.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, cb func(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListServiceQuotaIncreaseRequestsInTemplate",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ServiceQuotasAPI.ListServiceQuotaIncreaseRequestsInTemplatePagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListServiceQuotasWithContext(ctx context.Context, input *servicequotas.ListServiceQuotasInput, opts ...request.Option) (*servicequotas.ListServiceQuotasOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListServiceQuotas",
		Input:   input,
		Output:  (*servicequotas.ListServiceQuotasOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.ListServiceQuotasWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.ListServiceQuotasOutput), req.Error
}

func (c *Client) ListServiceQuotasPagesWithContext(ctx context.Context, input *servicequotas.ListServiceQuotasInput, cb func(*servicequotas.ListServiceQuotasOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListServiceQuotas",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ServiceQuotasAPI.ListServiceQuotasPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListServicesWithContext(ctx context.Context, input *servicequotas.ListServicesInput, opts ...request.Option) (*servicequotas.ListServicesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListServices",
		Input:   input,
		Output:  (*servicequotas.ListServicesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.ListServicesWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.ListServicesOutput), req.Error
}

func (c *Client) ListServicesPagesWithContext(ctx context.Context, input *servicequotas.ListServicesInput, cb func(*servicequotas.ListServicesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListServices",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ServiceQuotasAPI.ListServicesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *servicequotas.ListTagsForResourceInput, opts ...request.Option) (*servicequotas.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*servicequotas.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.ListTagsForResourceOutput), req.Error
}

func (c *Client) PutServiceQuotaIncreaseRequestIntoTemplateWithContext(ctx context.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput, opts ...request.Option) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "PutServiceQuotaIncreaseRequestIntoTemplate",
		Input:   input,
		Output:  (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.PutServiceQuotaIncreaseRequestIntoTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput), req.Error
}

func (c *Client) RequestServiceQuotaIncreaseWithContext(ctx context.Context, input *servicequotas.RequestServiceQuotaIncreaseInput, opts ...request.Option) (*servicequotas.RequestServiceQuotaIncreaseOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "RequestServiceQuotaIncrease",
		Input:   input,
		Output:  (*servicequotas.RequestServiceQuotaIncreaseOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.RequestServiceQuotaIncreaseWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.RequestServiceQuotaIncreaseOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *servicequotas.TagResourceInput, opts ...request.Option) (*servicequotas.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "TagResource",
		Input:   input,
		Output:  (*servicequotas.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *servicequotas.UntagResourceInput, opts ...request.Option) (*servicequotas.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicequotas",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*servicequotas.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceQuotasAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicequotas.UntagResourceOutput), req.Error
}
