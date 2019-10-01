// Code generated by internal/generate/main.go. DO NOT EDIT.

package acmpcactx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/aws/aws-sdk-go/service/acmpca/acmpcaiface"
	"github.com/glassechidna/awsctx"
)

type ACMPCA interface {
	CreateCertificateAuthorityWithContext(ctx context.Context, input *acmpca.CreateCertificateAuthorityInput, opts ...request.Option) (*acmpca.CreateCertificateAuthorityOutput, error)
	CreateCertificateAuthorityAuditReportWithContext(ctx context.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput, opts ...request.Option) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error)
	CreatePermissionWithContext(ctx context.Context, input *acmpca.CreatePermissionInput, opts ...request.Option) (*acmpca.CreatePermissionOutput, error)
	DeleteCertificateAuthorityWithContext(ctx context.Context, input *acmpca.DeleteCertificateAuthorityInput, opts ...request.Option) (*acmpca.DeleteCertificateAuthorityOutput, error)
	DeletePermissionWithContext(ctx context.Context, input *acmpca.DeletePermissionInput, opts ...request.Option) (*acmpca.DeletePermissionOutput, error)
	DescribeCertificateAuthorityWithContext(ctx context.Context, input *acmpca.DescribeCertificateAuthorityInput, opts ...request.Option) (*acmpca.DescribeCertificateAuthorityOutput, error)
	DescribeCertificateAuthorityAuditReportWithContext(ctx context.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput, opts ...request.Option) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error)
	GetCertificateWithContext(ctx context.Context, input *acmpca.GetCertificateInput, opts ...request.Option) (*acmpca.GetCertificateOutput, error)
	GetCertificateAuthorityCertificateWithContext(ctx context.Context, input *acmpca.GetCertificateAuthorityCertificateInput, opts ...request.Option) (*acmpca.GetCertificateAuthorityCertificateOutput, error)
	GetCertificateAuthorityCsrWithContext(ctx context.Context, input *acmpca.GetCertificateAuthorityCsrInput, opts ...request.Option) (*acmpca.GetCertificateAuthorityCsrOutput, error)
	ImportCertificateAuthorityCertificateWithContext(ctx context.Context, input *acmpca.ImportCertificateAuthorityCertificateInput, opts ...request.Option) (*acmpca.ImportCertificateAuthorityCertificateOutput, error)
	IssueCertificateWithContext(ctx context.Context, input *acmpca.IssueCertificateInput, opts ...request.Option) (*acmpca.IssueCertificateOutput, error)
	ListCertificateAuthoritiesWithContext(ctx context.Context, input *acmpca.ListCertificateAuthoritiesInput, opts ...request.Option) (*acmpca.ListCertificateAuthoritiesOutput, error)
	ListPermissionsWithContext(ctx context.Context, input *acmpca.ListPermissionsInput, opts ...request.Option) (*acmpca.ListPermissionsOutput, error)
	ListTagsWithContext(ctx context.Context, input *acmpca.ListTagsInput, opts ...request.Option) (*acmpca.ListTagsOutput, error)
	RestoreCertificateAuthorityWithContext(ctx context.Context, input *acmpca.RestoreCertificateAuthorityInput, opts ...request.Option) (*acmpca.RestoreCertificateAuthorityOutput, error)
	RevokeCertificateWithContext(ctx context.Context, input *acmpca.RevokeCertificateInput, opts ...request.Option) (*acmpca.RevokeCertificateOutput, error)
	TagCertificateAuthorityWithContext(ctx context.Context, input *acmpca.TagCertificateAuthorityInput, opts ...request.Option) (*acmpca.TagCertificateAuthorityOutput, error)
	UntagCertificateAuthorityWithContext(ctx context.Context, input *acmpca.UntagCertificateAuthorityInput, opts ...request.Option) (*acmpca.UntagCertificateAuthorityOutput, error)
	UpdateCertificateAuthorityWithContext(ctx context.Context, input *acmpca.UpdateCertificateAuthorityInput, opts ...request.Option) (*acmpca.UpdateCertificateAuthorityOutput, error)
}

type Client struct {
	acmpcaiface.ACMPCAAPI
	Contexter awsctx.Contexter
}

func New(base acmpcaiface.ACMPCAAPI, ctxer awsctx.Contexter) ACMPCA {
	return &Client{
		ACMPCAAPI: base,
		Contexter: ctxer,
	}
}

var _ ACMPCA = (*acmpca.ACMPCA)(nil)
var _ ACMPCA = (*Client)(nil)

func (c *Client) CreateCertificateAuthorityWithContext(ctx context.Context, input *acmpca.CreateCertificateAuthorityInput, opts ...request.Option) (*acmpca.CreateCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "CreateCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*acmpca.CreateCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.CreateCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.CreateCertificateAuthorityOutput), req.Error
}

func (c *Client) CreateCertificateAuthorityAuditReportWithContext(ctx context.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput, opts ...request.Option) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "CreateCertificateAuthorityAuditReportWithContext",
		Input:   input,
		Output:  (*acmpca.CreateCertificateAuthorityAuditReportOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.CreateCertificateAuthorityAuditReportWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.CreateCertificateAuthorityAuditReportOutput), req.Error
}

func (c *Client) CreatePermissionWithContext(ctx context.Context, input *acmpca.CreatePermissionInput, opts ...request.Option) (*acmpca.CreatePermissionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "CreatePermissionWithContext",
		Input:   input,
		Output:  (*acmpca.CreatePermissionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.CreatePermissionWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.CreatePermissionOutput), req.Error
}

func (c *Client) DeleteCertificateAuthorityWithContext(ctx context.Context, input *acmpca.DeleteCertificateAuthorityInput, opts ...request.Option) (*acmpca.DeleteCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "DeleteCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*acmpca.DeleteCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.DeleteCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.DeleteCertificateAuthorityOutput), req.Error
}

func (c *Client) DeletePermissionWithContext(ctx context.Context, input *acmpca.DeletePermissionInput, opts ...request.Option) (*acmpca.DeletePermissionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "DeletePermissionWithContext",
		Input:   input,
		Output:  (*acmpca.DeletePermissionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.DeletePermissionWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.DeletePermissionOutput), req.Error
}

func (c *Client) DescribeCertificateAuthorityWithContext(ctx context.Context, input *acmpca.DescribeCertificateAuthorityInput, opts ...request.Option) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "DescribeCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*acmpca.DescribeCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.DescribeCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.DescribeCertificateAuthorityOutput), req.Error
}

func (c *Client) DescribeCertificateAuthorityAuditReportWithContext(ctx context.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput, opts ...request.Option) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "DescribeCertificateAuthorityAuditReportWithContext",
		Input:   input,
		Output:  (*acmpca.DescribeCertificateAuthorityAuditReportOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.DescribeCertificateAuthorityAuditReportWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.DescribeCertificateAuthorityAuditReportOutput), req.Error
}

func (c *Client) GetCertificateWithContext(ctx context.Context, input *acmpca.GetCertificateInput, opts ...request.Option) (*acmpca.GetCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "GetCertificateWithContext",
		Input:   input,
		Output:  (*acmpca.GetCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.GetCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.GetCertificateOutput), req.Error
}

func (c *Client) GetCertificateAuthorityCertificateWithContext(ctx context.Context, input *acmpca.GetCertificateAuthorityCertificateInput, opts ...request.Option) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "GetCertificateAuthorityCertificateWithContext",
		Input:   input,
		Output:  (*acmpca.GetCertificateAuthorityCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.GetCertificateAuthorityCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.GetCertificateAuthorityCertificateOutput), req.Error
}

func (c *Client) GetCertificateAuthorityCsrWithContext(ctx context.Context, input *acmpca.GetCertificateAuthorityCsrInput, opts ...request.Option) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "GetCertificateAuthorityCsrWithContext",
		Input:   input,
		Output:  (*acmpca.GetCertificateAuthorityCsrOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.GetCertificateAuthorityCsrWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.GetCertificateAuthorityCsrOutput), req.Error
}

func (c *Client) ImportCertificateAuthorityCertificateWithContext(ctx context.Context, input *acmpca.ImportCertificateAuthorityCertificateInput, opts ...request.Option) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "ImportCertificateAuthorityCertificateWithContext",
		Input:   input,
		Output:  (*acmpca.ImportCertificateAuthorityCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.ImportCertificateAuthorityCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.ImportCertificateAuthorityCertificateOutput), req.Error
}

func (c *Client) IssueCertificateWithContext(ctx context.Context, input *acmpca.IssueCertificateInput, opts ...request.Option) (*acmpca.IssueCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "IssueCertificateWithContext",
		Input:   input,
		Output:  (*acmpca.IssueCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.IssueCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.IssueCertificateOutput), req.Error
}

func (c *Client) ListCertificateAuthoritiesWithContext(ctx context.Context, input *acmpca.ListCertificateAuthoritiesInput, opts ...request.Option) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "ListCertificateAuthoritiesWithContext",
		Input:   input,
		Output:  (*acmpca.ListCertificateAuthoritiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.ListCertificateAuthoritiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.ListCertificateAuthoritiesOutput), req.Error
}

func (c *Client) ListPermissionsWithContext(ctx context.Context, input *acmpca.ListPermissionsInput, opts ...request.Option) (*acmpca.ListPermissionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "ListPermissionsWithContext",
		Input:   input,
		Output:  (*acmpca.ListPermissionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.ListPermissionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.ListPermissionsOutput), req.Error
}

func (c *Client) ListTagsWithContext(ctx context.Context, input *acmpca.ListTagsInput, opts ...request.Option) (*acmpca.ListTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "ListTagsWithContext",
		Input:   input,
		Output:  (*acmpca.ListTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.ListTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.ListTagsOutput), req.Error
}

func (c *Client) RestoreCertificateAuthorityWithContext(ctx context.Context, input *acmpca.RestoreCertificateAuthorityInput, opts ...request.Option) (*acmpca.RestoreCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "RestoreCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*acmpca.RestoreCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.RestoreCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.RestoreCertificateAuthorityOutput), req.Error
}

func (c *Client) RevokeCertificateWithContext(ctx context.Context, input *acmpca.RevokeCertificateInput, opts ...request.Option) (*acmpca.RevokeCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "RevokeCertificateWithContext",
		Input:   input,
		Output:  (*acmpca.RevokeCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.RevokeCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.RevokeCertificateOutput), req.Error
}

func (c *Client) TagCertificateAuthorityWithContext(ctx context.Context, input *acmpca.TagCertificateAuthorityInput, opts ...request.Option) (*acmpca.TagCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "TagCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*acmpca.TagCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.TagCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.TagCertificateAuthorityOutput), req.Error
}

func (c *Client) UntagCertificateAuthorityWithContext(ctx context.Context, input *acmpca.UntagCertificateAuthorityInput, opts ...request.Option) (*acmpca.UntagCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "UntagCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*acmpca.UntagCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.UntagCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.UntagCertificateAuthorityOutput), req.Error
}

func (c *Client) UpdateCertificateAuthorityWithContext(ctx context.Context, input *acmpca.UpdateCertificateAuthorityInput, opts ...request.Option) (*acmpca.UpdateCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "acmpca",
		Action:  "UpdateCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*acmpca.UpdateCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ACMPCAAPI.UpdateCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*acmpca.UpdateCertificateAuthorityOutput), req.Error
}
