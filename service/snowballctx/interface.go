// Code generated by internal/generate/main.go. DO NOT EDIT.

package snowballctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/snowball"
	"github.com/aws/aws-sdk-go/service/snowball/snowballiface"
	"github.com/glassechidna/awsctx"
)

type Snowball interface {
	CancelClusterWithContext(ctx context.Context, input *snowball.CancelClusterInput, opts ...request.Option) (*snowball.CancelClusterOutput, error)
	CancelJobWithContext(ctx context.Context, input *snowball.CancelJobInput, opts ...request.Option) (*snowball.CancelJobOutput, error)
	CreateAddressWithContext(ctx context.Context, input *snowball.CreateAddressInput, opts ...request.Option) (*snowball.CreateAddressOutput, error)
	CreateClusterWithContext(ctx context.Context, input *snowball.CreateClusterInput, opts ...request.Option) (*snowball.CreateClusterOutput, error)
	CreateJobWithContext(ctx context.Context, input *snowball.CreateJobInput, opts ...request.Option) (*snowball.CreateJobOutput, error)
	CreateLongTermPricingWithContext(ctx context.Context, input *snowball.CreateLongTermPricingInput, opts ...request.Option) (*snowball.CreateLongTermPricingOutput, error)
	CreateReturnShippingLabelWithContext(ctx context.Context, input *snowball.CreateReturnShippingLabelInput, opts ...request.Option) (*snowball.CreateReturnShippingLabelOutput, error)
	DescribeAddressWithContext(ctx context.Context, input *snowball.DescribeAddressInput, opts ...request.Option) (*snowball.DescribeAddressOutput, error)
	DescribeAddressesWithContext(ctx context.Context, input *snowball.DescribeAddressesInput, opts ...request.Option) (*snowball.DescribeAddressesOutput, error)
	DescribeAddressesPagesWithContext(ctx context.Context, input *snowball.DescribeAddressesInput, cb func(*snowball.DescribeAddressesOutput, bool) bool, opts ...request.Option) error
	DescribeClusterWithContext(ctx context.Context, input *snowball.DescribeClusterInput, opts ...request.Option) (*snowball.DescribeClusterOutput, error)
	DescribeJobWithContext(ctx context.Context, input *snowball.DescribeJobInput, opts ...request.Option) (*snowball.DescribeJobOutput, error)
	DescribeReturnShippingLabelWithContext(ctx context.Context, input *snowball.DescribeReturnShippingLabelInput, opts ...request.Option) (*snowball.DescribeReturnShippingLabelOutput, error)
	GetJobManifestWithContext(ctx context.Context, input *snowball.GetJobManifestInput, opts ...request.Option) (*snowball.GetJobManifestOutput, error)
	GetJobUnlockCodeWithContext(ctx context.Context, input *snowball.GetJobUnlockCodeInput, opts ...request.Option) (*snowball.GetJobUnlockCodeOutput, error)
	GetSnowballUsageWithContext(ctx context.Context, input *snowball.GetSnowballUsageInput, opts ...request.Option) (*snowball.GetSnowballUsageOutput, error)
	GetSoftwareUpdatesWithContext(ctx context.Context, input *snowball.GetSoftwareUpdatesInput, opts ...request.Option) (*snowball.GetSoftwareUpdatesOutput, error)
	ListClusterJobsWithContext(ctx context.Context, input *snowball.ListClusterJobsInput, opts ...request.Option) (*snowball.ListClusterJobsOutput, error)
	ListClusterJobsPagesWithContext(ctx context.Context, input *snowball.ListClusterJobsInput, cb func(*snowball.ListClusterJobsOutput, bool) bool, opts ...request.Option) error
	ListClustersWithContext(ctx context.Context, input *snowball.ListClustersInput, opts ...request.Option) (*snowball.ListClustersOutput, error)
	ListClustersPagesWithContext(ctx context.Context, input *snowball.ListClustersInput, cb func(*snowball.ListClustersOutput, bool) bool, opts ...request.Option) error
	ListCompatibleImagesWithContext(ctx context.Context, input *snowball.ListCompatibleImagesInput, opts ...request.Option) (*snowball.ListCompatibleImagesOutput, error)
	ListCompatibleImagesPagesWithContext(ctx context.Context, input *snowball.ListCompatibleImagesInput, cb func(*snowball.ListCompatibleImagesOutput, bool) bool, opts ...request.Option) error
	ListJobsWithContext(ctx context.Context, input *snowball.ListJobsInput, opts ...request.Option) (*snowball.ListJobsOutput, error)
	ListJobsPagesWithContext(ctx context.Context, input *snowball.ListJobsInput, cb func(*snowball.ListJobsOutput, bool) bool, opts ...request.Option) error
	ListLongTermPricingWithContext(ctx context.Context, input *snowball.ListLongTermPricingInput, opts ...request.Option) (*snowball.ListLongTermPricingOutput, error)
	ListLongTermPricingPagesWithContext(ctx context.Context, input *snowball.ListLongTermPricingInput, cb func(*snowball.ListLongTermPricingOutput, bool) bool, opts ...request.Option) error
	ListServiceVersionsWithContext(ctx context.Context, input *snowball.ListServiceVersionsInput, opts ...request.Option) (*snowball.ListServiceVersionsOutput, error)
	UpdateClusterWithContext(ctx context.Context, input *snowball.UpdateClusterInput, opts ...request.Option) (*snowball.UpdateClusterOutput, error)
	UpdateJobWithContext(ctx context.Context, input *snowball.UpdateJobInput, opts ...request.Option) (*snowball.UpdateJobOutput, error)
	UpdateJobShipmentStateWithContext(ctx context.Context, input *snowball.UpdateJobShipmentStateInput, opts ...request.Option) (*snowball.UpdateJobShipmentStateOutput, error)
	UpdateLongTermPricingWithContext(ctx context.Context, input *snowball.UpdateLongTermPricingInput, opts ...request.Option) (*snowball.UpdateLongTermPricingOutput, error)
}

type Client struct {
	snowballiface.SnowballAPI
	Contexter awsctx.Contexter
}

func New(base snowballiface.SnowballAPI, ctxer awsctx.Contexter) Snowball {
	return &Client{
		SnowballAPI: base,
		Contexter: ctxer,
	}
}

var _ Snowball = (*snowball.Snowball)(nil)
var _ Snowball = (*Client)(nil)

func (c *Client) CancelClusterWithContext(ctx context.Context, input *snowball.CancelClusterInput, opts ...request.Option) (*snowball.CancelClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "CancelCluster",
		Input:   input,
		Output:  (*snowball.CancelClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.CancelClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.CancelClusterOutput), req.Error
}

func (c *Client) CancelJobWithContext(ctx context.Context, input *snowball.CancelJobInput, opts ...request.Option) (*snowball.CancelJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "CancelJob",
		Input:   input,
		Output:  (*snowball.CancelJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.CancelJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.CancelJobOutput), req.Error
}

func (c *Client) CreateAddressWithContext(ctx context.Context, input *snowball.CreateAddressInput, opts ...request.Option) (*snowball.CreateAddressOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "CreateAddress",
		Input:   input,
		Output:  (*snowball.CreateAddressOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.CreateAddressWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.CreateAddressOutput), req.Error
}

func (c *Client) CreateClusterWithContext(ctx context.Context, input *snowball.CreateClusterInput, opts ...request.Option) (*snowball.CreateClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "CreateCluster",
		Input:   input,
		Output:  (*snowball.CreateClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.CreateClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.CreateClusterOutput), req.Error
}

func (c *Client) CreateJobWithContext(ctx context.Context, input *snowball.CreateJobInput, opts ...request.Option) (*snowball.CreateJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "CreateJob",
		Input:   input,
		Output:  (*snowball.CreateJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.CreateJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.CreateJobOutput), req.Error
}

func (c *Client) CreateLongTermPricingWithContext(ctx context.Context, input *snowball.CreateLongTermPricingInput, opts ...request.Option) (*snowball.CreateLongTermPricingOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "CreateLongTermPricing",
		Input:   input,
		Output:  (*snowball.CreateLongTermPricingOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.CreateLongTermPricingWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.CreateLongTermPricingOutput), req.Error
}

func (c *Client) CreateReturnShippingLabelWithContext(ctx context.Context, input *snowball.CreateReturnShippingLabelInput, opts ...request.Option) (*snowball.CreateReturnShippingLabelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "CreateReturnShippingLabel",
		Input:   input,
		Output:  (*snowball.CreateReturnShippingLabelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.CreateReturnShippingLabelWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.CreateReturnShippingLabelOutput), req.Error
}

func (c *Client) DescribeAddressWithContext(ctx context.Context, input *snowball.DescribeAddressInput, opts ...request.Option) (*snowball.DescribeAddressOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "DescribeAddress",
		Input:   input,
		Output:  (*snowball.DescribeAddressOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.DescribeAddressWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.DescribeAddressOutput), req.Error
}

func (c *Client) DescribeAddressesWithContext(ctx context.Context, input *snowball.DescribeAddressesInput, opts ...request.Option) (*snowball.DescribeAddressesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "DescribeAddresses",
		Input:   input,
		Output:  (*snowball.DescribeAddressesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.DescribeAddressesWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.DescribeAddressesOutput), req.Error
}

func (c *Client) DescribeAddressesPagesWithContext(ctx context.Context, input *snowball.DescribeAddressesInput, cb func(*snowball.DescribeAddressesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "DescribeAddresses",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SnowballAPI.DescribeAddressesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeClusterWithContext(ctx context.Context, input *snowball.DescribeClusterInput, opts ...request.Option) (*snowball.DescribeClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "DescribeCluster",
		Input:   input,
		Output:  (*snowball.DescribeClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.DescribeClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.DescribeClusterOutput), req.Error
}

func (c *Client) DescribeJobWithContext(ctx context.Context, input *snowball.DescribeJobInput, opts ...request.Option) (*snowball.DescribeJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "DescribeJob",
		Input:   input,
		Output:  (*snowball.DescribeJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.DescribeJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.DescribeJobOutput), req.Error
}

func (c *Client) DescribeReturnShippingLabelWithContext(ctx context.Context, input *snowball.DescribeReturnShippingLabelInput, opts ...request.Option) (*snowball.DescribeReturnShippingLabelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "DescribeReturnShippingLabel",
		Input:   input,
		Output:  (*snowball.DescribeReturnShippingLabelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.DescribeReturnShippingLabelWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.DescribeReturnShippingLabelOutput), req.Error
}

func (c *Client) GetJobManifestWithContext(ctx context.Context, input *snowball.GetJobManifestInput, opts ...request.Option) (*snowball.GetJobManifestOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "GetJobManifest",
		Input:   input,
		Output:  (*snowball.GetJobManifestOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.GetJobManifestWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.GetJobManifestOutput), req.Error
}

func (c *Client) GetJobUnlockCodeWithContext(ctx context.Context, input *snowball.GetJobUnlockCodeInput, opts ...request.Option) (*snowball.GetJobUnlockCodeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "GetJobUnlockCode",
		Input:   input,
		Output:  (*snowball.GetJobUnlockCodeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.GetJobUnlockCodeWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.GetJobUnlockCodeOutput), req.Error
}

func (c *Client) GetSnowballUsageWithContext(ctx context.Context, input *snowball.GetSnowballUsageInput, opts ...request.Option) (*snowball.GetSnowballUsageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "GetSnowballUsage",
		Input:   input,
		Output:  (*snowball.GetSnowballUsageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.GetSnowballUsageWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.GetSnowballUsageOutput), req.Error
}

func (c *Client) GetSoftwareUpdatesWithContext(ctx context.Context, input *snowball.GetSoftwareUpdatesInput, opts ...request.Option) (*snowball.GetSoftwareUpdatesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "GetSoftwareUpdates",
		Input:   input,
		Output:  (*snowball.GetSoftwareUpdatesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.GetSoftwareUpdatesWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.GetSoftwareUpdatesOutput), req.Error
}

func (c *Client) ListClusterJobsWithContext(ctx context.Context, input *snowball.ListClusterJobsInput, opts ...request.Option) (*snowball.ListClusterJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListClusterJobs",
		Input:   input,
		Output:  (*snowball.ListClusterJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.ListClusterJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.ListClusterJobsOutput), req.Error
}

func (c *Client) ListClusterJobsPagesWithContext(ctx context.Context, input *snowball.ListClusterJobsInput, cb func(*snowball.ListClusterJobsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListClusterJobs",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SnowballAPI.ListClusterJobsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListClustersWithContext(ctx context.Context, input *snowball.ListClustersInput, opts ...request.Option) (*snowball.ListClustersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListClusters",
		Input:   input,
		Output:  (*snowball.ListClustersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.ListClustersWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.ListClustersOutput), req.Error
}

func (c *Client) ListClustersPagesWithContext(ctx context.Context, input *snowball.ListClustersInput, cb func(*snowball.ListClustersOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListClusters",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SnowballAPI.ListClustersPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListCompatibleImagesWithContext(ctx context.Context, input *snowball.ListCompatibleImagesInput, opts ...request.Option) (*snowball.ListCompatibleImagesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListCompatibleImages",
		Input:   input,
		Output:  (*snowball.ListCompatibleImagesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.ListCompatibleImagesWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.ListCompatibleImagesOutput), req.Error
}

func (c *Client) ListCompatibleImagesPagesWithContext(ctx context.Context, input *snowball.ListCompatibleImagesInput, cb func(*snowball.ListCompatibleImagesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListCompatibleImages",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SnowballAPI.ListCompatibleImagesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListJobsWithContext(ctx context.Context, input *snowball.ListJobsInput, opts ...request.Option) (*snowball.ListJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListJobs",
		Input:   input,
		Output:  (*snowball.ListJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.ListJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.ListJobsOutput), req.Error
}

func (c *Client) ListJobsPagesWithContext(ctx context.Context, input *snowball.ListJobsInput, cb func(*snowball.ListJobsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListJobs",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SnowballAPI.ListJobsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListLongTermPricingWithContext(ctx context.Context, input *snowball.ListLongTermPricingInput, opts ...request.Option) (*snowball.ListLongTermPricingOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListLongTermPricing",
		Input:   input,
		Output:  (*snowball.ListLongTermPricingOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.ListLongTermPricingWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.ListLongTermPricingOutput), req.Error
}

func (c *Client) ListLongTermPricingPagesWithContext(ctx context.Context, input *snowball.ListLongTermPricingInput, cb func(*snowball.ListLongTermPricingOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListLongTermPricing",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SnowballAPI.ListLongTermPricingPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListServiceVersionsWithContext(ctx context.Context, input *snowball.ListServiceVersionsInput, opts ...request.Option) (*snowball.ListServiceVersionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "ListServiceVersions",
		Input:   input,
		Output:  (*snowball.ListServiceVersionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.ListServiceVersionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.ListServiceVersionsOutput), req.Error
}

func (c *Client) UpdateClusterWithContext(ctx context.Context, input *snowball.UpdateClusterInput, opts ...request.Option) (*snowball.UpdateClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "UpdateCluster",
		Input:   input,
		Output:  (*snowball.UpdateClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.UpdateClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.UpdateClusterOutput), req.Error
}

func (c *Client) UpdateJobWithContext(ctx context.Context, input *snowball.UpdateJobInput, opts ...request.Option) (*snowball.UpdateJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "UpdateJob",
		Input:   input,
		Output:  (*snowball.UpdateJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.UpdateJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.UpdateJobOutput), req.Error
}

func (c *Client) UpdateJobShipmentStateWithContext(ctx context.Context, input *snowball.UpdateJobShipmentStateInput, opts ...request.Option) (*snowball.UpdateJobShipmentStateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "UpdateJobShipmentState",
		Input:   input,
		Output:  (*snowball.UpdateJobShipmentStateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.UpdateJobShipmentStateWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.UpdateJobShipmentStateOutput), req.Error
}

func (c *Client) UpdateLongTermPricingWithContext(ctx context.Context, input *snowball.UpdateLongTermPricingInput, opts ...request.Option) (*snowball.UpdateLongTermPricingOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "snowball",
		Action:  "UpdateLongTermPricing",
		Input:   input,
		Output:  (*snowball.UpdateLongTermPricingOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SnowballAPI.UpdateLongTermPricingWithContext(ctx, input, opts...)
	})

	return req.Output.(*snowball.UpdateLongTermPricingOutput), req.Error
}
