// Code generated by internal/generate/main.go. DO NOT EDIT.

package costandusagereportservicectx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice/costandusagereportserviceiface"
	"github.com/glassechidna/awsctx"
)

type CostandUsageReportService interface {
	DeleteReportDefinitionWithContext(ctx context.Context, input *costandusagereportservice.DeleteReportDefinitionInput, opts ...request.Option) (*costandusagereportservice.DeleteReportDefinitionOutput, error)
	DescribeReportDefinitionsWithContext(ctx context.Context, input *costandusagereportservice.DescribeReportDefinitionsInput, opts ...request.Option) (*costandusagereportservice.DescribeReportDefinitionsOutput, error)
	PutReportDefinitionWithContext(ctx context.Context, input *costandusagereportservice.PutReportDefinitionInput, opts ...request.Option) (*costandusagereportservice.PutReportDefinitionOutput, error)
}

type Client struct {
	costandusagereportserviceiface.CostandUsageReportServiceAPI
	Contexter awsctx.Contexter
}

func New(base costandusagereportserviceiface.CostandUsageReportServiceAPI, ctxer awsctx.Contexter) CostandUsageReportService {
	return &Client{
		CostandUsageReportServiceAPI: base,
		Contexter: ctxer,
	}
}

var _ CostandUsageReportService = (*costandusagereportservice.CostandUsageReportService)(nil)
var _ CostandUsageReportService = (*Client)(nil)

func (c *Client) DeleteReportDefinitionWithContext(ctx context.Context, input *costandusagereportservice.DeleteReportDefinitionInput, opts ...request.Option) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costandusagereportservice",
		Action:  "DeleteReportDefinitionWithContext",
		Input:   input,
		Output:  (*costandusagereportservice.DeleteReportDefinitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostandUsageReportServiceAPI.DeleteReportDefinitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*costandusagereportservice.DeleteReportDefinitionOutput), req.Error
}

func (c *Client) DescribeReportDefinitionsWithContext(ctx context.Context, input *costandusagereportservice.DescribeReportDefinitionsInput, opts ...request.Option) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costandusagereportservice",
		Action:  "DescribeReportDefinitionsWithContext",
		Input:   input,
		Output:  (*costandusagereportservice.DescribeReportDefinitionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostandUsageReportServiceAPI.DescribeReportDefinitionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*costandusagereportservice.DescribeReportDefinitionsOutput), req.Error
}

func (c *Client) PutReportDefinitionWithContext(ctx context.Context, input *costandusagereportservice.PutReportDefinitionInput, opts ...request.Option) (*costandusagereportservice.PutReportDefinitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costandusagereportservice",
		Action:  "PutReportDefinitionWithContext",
		Input:   input,
		Output:  (*costandusagereportservice.PutReportDefinitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostandUsageReportServiceAPI.PutReportDefinitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*costandusagereportservice.PutReportDefinitionOutput), req.Error
}
