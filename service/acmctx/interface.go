// Code generated by internal/generate/main.go. DO NOT EDIT.

package acmctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acm/acmiface"
	"github.com/glassechidna/awsctx"
)

type ACM interface {
	AddTagsToCertificateWithContext(ctx context.Context, input *acm.AddTagsToCertificateInput, opts ...request.Option) (*acm.AddTagsToCertificateOutput, error)
	DeleteCertificateWithContext(ctx context.Context, input *acm.DeleteCertificateInput, opts ...request.Option) (*acm.DeleteCertificateOutput, error)
	DescribeCertificateWithContext(ctx context.Context, input *acm.DescribeCertificateInput, opts ...request.Option) (*acm.DescribeCertificateOutput, error)
	ExportCertificateWithContext(ctx context.Context, input *acm.ExportCertificateInput, opts ...request.Option) (*acm.ExportCertificateOutput, error)
	GetCertificateWithContext(ctx context.Context, input *acm.GetCertificateInput, opts ...request.Option) (*acm.GetCertificateOutput, error)
	ImportCertificateWithContext(ctx context.Context, input *acm.ImportCertificateInput, opts ...request.Option) (*acm.ImportCertificateOutput, error)
	ListCertificatesWithContext(ctx context.Context, input *acm.ListCertificatesInput, opts ...request.Option) (*acm.ListCertificatesOutput, error)
	ListCertificatesPagesWithContext(ctx context.Context, input *acm.ListCertificatesInput, cb func(*acm.ListCertificatesOutput, bool) bool, opts ...request.Option) error
	ListTagsForCertificateWithContext(ctx context.Context, input *acm.ListTagsForCertificateInput, opts ...request.Option) (*acm.ListTagsForCertificateOutput, error)
	RemoveTagsFromCertificateWithContext(ctx context.Context, input *acm.RemoveTagsFromCertificateInput, opts ...request.Option) (*acm.RemoveTagsFromCertificateOutput, error)
	RenewCertificateWithContext(ctx context.Context, input *acm.RenewCertificateInput, opts ...request.Option) (*acm.RenewCertificateOutput, error)
	RequestCertificateWithContext(ctx context.Context, input *acm.RequestCertificateInput, opts ...request.Option) (*acm.RequestCertificateOutput, error)
	ResendValidationEmailWithContext(ctx context.Context, input *acm.ResendValidationEmailInput, opts ...request.Option) (*acm.ResendValidationEmailOutput, error)
	UpdateCertificateOptionsWithContext(ctx context.Context, input *acm.UpdateCertificateOptionsInput, opts ...request.Option) (*acm.UpdateCertificateOptionsOutput, error)
}

type Client struct {
	acmiface.ACMAPI
	Contexter awsctx.Contexter
}

func New(base acmiface.ACMAPI, ctxer awsctx.Contexter) ACM {
	return &Client{
		ACMAPI: base,
		Contexter: ctxer,
	}
}

var _ ACM = (*acm.ACM)(nil)
var _ ACM = (*Client)(nil)

func (c *Client) AddTagsToCertificateWithContext(ctx context.Context, input *acm.AddTagsToCertificateInput, opts ...request.Option) (*acm.AddTagsToCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "AddTagsToCertificate",
		Input:   input,
		Output:  (*acm.AddTagsToCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.AddTagsToCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.AddTagsToCertificateOutput), req.Error
}

func (c *Client) DeleteCertificateWithContext(ctx context.Context, input *acm.DeleteCertificateInput, opts ...request.Option) (*acm.DeleteCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "DeleteCertificate",
		Input:   input,
		Output:  (*acm.DeleteCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.DeleteCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.DeleteCertificateOutput), req.Error
}

func (c *Client) DescribeCertificateWithContext(ctx context.Context, input *acm.DescribeCertificateInput, opts ...request.Option) (*acm.DescribeCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "DescribeCertificate",
		Input:   input,
		Output:  (*acm.DescribeCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.DescribeCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.DescribeCertificateOutput), req.Error
}

func (c *Client) ExportCertificateWithContext(ctx context.Context, input *acm.ExportCertificateInput, opts ...request.Option) (*acm.ExportCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "ExportCertificate",
		Input:   input,
		Output:  (*acm.ExportCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.ExportCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.ExportCertificateOutput), req.Error
}

func (c *Client) GetCertificateWithContext(ctx context.Context, input *acm.GetCertificateInput, opts ...request.Option) (*acm.GetCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "GetCertificate",
		Input:   input,
		Output:  (*acm.GetCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.GetCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.GetCertificateOutput), req.Error
}

func (c *Client) ImportCertificateWithContext(ctx context.Context, input *acm.ImportCertificateInput, opts ...request.Option) (*acm.ImportCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "ImportCertificate",
		Input:   input,
		Output:  (*acm.ImportCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.ImportCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.ImportCertificateOutput), req.Error
}

func (c *Client) ListCertificatesWithContext(ctx context.Context, input *acm.ListCertificatesInput, opts ...request.Option) (*acm.ListCertificatesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "ListCertificates",
		Input:   input,
		Output:  (*acm.ListCertificatesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.ListCertificatesWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.ListCertificatesOutput), req.Error
}

func (c *Client) ListCertificatesPagesWithContext(ctx context.Context, input *acm.ListCertificatesInput, cb func(*acm.ListCertificatesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "ListCertificates",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ACMAPI.ListCertificatesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForCertificateWithContext(ctx context.Context, input *acm.ListTagsForCertificateInput, opts ...request.Option) (*acm.ListTagsForCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "ListTagsForCertificate",
		Input:   input,
		Output:  (*acm.ListTagsForCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.ListTagsForCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.ListTagsForCertificateOutput), req.Error
}

func (c *Client) RemoveTagsFromCertificateWithContext(ctx context.Context, input *acm.RemoveTagsFromCertificateInput, opts ...request.Option) (*acm.RemoveTagsFromCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "RemoveTagsFromCertificate",
		Input:   input,
		Output:  (*acm.RemoveTagsFromCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.RemoveTagsFromCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.RemoveTagsFromCertificateOutput), req.Error
}

func (c *Client) RenewCertificateWithContext(ctx context.Context, input *acm.RenewCertificateInput, opts ...request.Option) (*acm.RenewCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "RenewCertificate",
		Input:   input,
		Output:  (*acm.RenewCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.RenewCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.RenewCertificateOutput), req.Error
}

func (c *Client) RequestCertificateWithContext(ctx context.Context, input *acm.RequestCertificateInput, opts ...request.Option) (*acm.RequestCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "RequestCertificate",
		Input:   input,
		Output:  (*acm.RequestCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.RequestCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.RequestCertificateOutput), req.Error
}

func (c *Client) ResendValidationEmailWithContext(ctx context.Context, input *acm.ResendValidationEmailInput, opts ...request.Option) (*acm.ResendValidationEmailOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "ResendValidationEmail",
		Input:   input,
		Output:  (*acm.ResendValidationEmailOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.ResendValidationEmailWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.ResendValidationEmailOutput), req.Error
}

func (c *Client) UpdateCertificateOptionsWithContext(ctx context.Context, input *acm.UpdateCertificateOptionsInput, opts ...request.Option) (*acm.UpdateCertificateOptionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acm",
		Action:  "UpdateCertificateOptions",
		Input:   input,
		Output:  (*acm.UpdateCertificateOptionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMAPI.UpdateCertificateOptionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*acm.UpdateCertificateOptionsOutput), req.Error
}
