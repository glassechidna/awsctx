// Code generated by internal/generate/main.go. DO NOT EDIT.

package comprehendmedicalctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"github.com/aws/aws-sdk-go/service/comprehendmedical/comprehendmedicaliface"
	"github.com/glassechidna/awsctx"
)

type ComprehendMedical interface {
	DescribeEntitiesDetectionV2JobWithContext(ctx context.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput, opts ...request.Option) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error)
	DescribePHIDetectionJobWithContext(ctx context.Context, input *comprehendmedical.DescribePHIDetectionJobInput, opts ...request.Option) (*comprehendmedical.DescribePHIDetectionJobOutput, error)
	DetectEntitiesWithContext(ctx context.Context, input *comprehendmedical.DetectEntitiesInput, opts ...request.Option) (*comprehendmedical.DetectEntitiesOutput, error)
	DetectEntitiesV2WithContext(ctx context.Context, input *comprehendmedical.DetectEntitiesV2Input, opts ...request.Option) (*comprehendmedical.DetectEntitiesV2Output, error)
	DetectPHIWithContext(ctx context.Context, input *comprehendmedical.DetectPHIInput, opts ...request.Option) (*comprehendmedical.DetectPHIOutput, error)
	InferICD10CMWithContext(ctx context.Context, input *comprehendmedical.InferICD10CMInput, opts ...request.Option) (*comprehendmedical.InferICD10CMOutput, error)
	InferRxNormWithContext(ctx context.Context, input *comprehendmedical.InferRxNormInput, opts ...request.Option) (*comprehendmedical.InferRxNormOutput, error)
	ListEntitiesDetectionV2JobsWithContext(ctx context.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput, opts ...request.Option) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error)
	ListPHIDetectionJobsWithContext(ctx context.Context, input *comprehendmedical.ListPHIDetectionJobsInput, opts ...request.Option) (*comprehendmedical.ListPHIDetectionJobsOutput, error)
	StartEntitiesDetectionV2JobWithContext(ctx context.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput, opts ...request.Option) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error)
	StartPHIDetectionJobWithContext(ctx context.Context, input *comprehendmedical.StartPHIDetectionJobInput, opts ...request.Option) (*comprehendmedical.StartPHIDetectionJobOutput, error)
	StopEntitiesDetectionV2JobWithContext(ctx context.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput, opts ...request.Option) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error)
	StopPHIDetectionJobWithContext(ctx context.Context, input *comprehendmedical.StopPHIDetectionJobInput, opts ...request.Option) (*comprehendmedical.StopPHIDetectionJobOutput, error)
}

type Client struct {
	comprehendmedicaliface.ComprehendMedicalAPI
	Contexter awsctx.Contexter
}

func New(base comprehendmedicaliface.ComprehendMedicalAPI, ctxer awsctx.Contexter) ComprehendMedical {
	return &Client{
		ComprehendMedicalAPI: base,
		Contexter: ctxer,
	}
}

var _ ComprehendMedical = (*comprehendmedical.ComprehendMedical)(nil)
var _ ComprehendMedical = (*Client)(nil)

func (c *Client) DescribeEntitiesDetectionV2JobWithContext(ctx context.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput, opts ...request.Option) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "DescribeEntitiesDetectionV2Job",
		Input:   input,
		Output:  (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.DescribeEntitiesDetectionV2JobWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput), req.Error
}

func (c *Client) DescribePHIDetectionJobWithContext(ctx context.Context, input *comprehendmedical.DescribePHIDetectionJobInput, opts ...request.Option) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "DescribePHIDetectionJob",
		Input:   input,
		Output:  (*comprehendmedical.DescribePHIDetectionJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.DescribePHIDetectionJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.DescribePHIDetectionJobOutput), req.Error
}

func (c *Client) DetectEntitiesWithContext(ctx context.Context, input *comprehendmedical.DetectEntitiesInput, opts ...request.Option) (*comprehendmedical.DetectEntitiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "DetectEntities",
		Input:   input,
		Output:  (*comprehendmedical.DetectEntitiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.DetectEntitiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.DetectEntitiesOutput), req.Error
}

func (c *Client) DetectEntitiesV2WithContext(ctx context.Context, input *comprehendmedical.DetectEntitiesV2Input, opts ...request.Option) (*comprehendmedical.DetectEntitiesV2Output, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "DetectEntitiesV2",
		Input:   input,
		Output:  (*comprehendmedical.DetectEntitiesV2Output)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.DetectEntitiesV2WithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.DetectEntitiesV2Output), req.Error
}

func (c *Client) DetectPHIWithContext(ctx context.Context, input *comprehendmedical.DetectPHIInput, opts ...request.Option) (*comprehendmedical.DetectPHIOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "DetectPHI",
		Input:   input,
		Output:  (*comprehendmedical.DetectPHIOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.DetectPHIWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.DetectPHIOutput), req.Error
}

func (c *Client) InferICD10CMWithContext(ctx context.Context, input *comprehendmedical.InferICD10CMInput, opts ...request.Option) (*comprehendmedical.InferICD10CMOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "InferICD10CM",
		Input:   input,
		Output:  (*comprehendmedical.InferICD10CMOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.InferICD10CMWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.InferICD10CMOutput), req.Error
}

func (c *Client) InferRxNormWithContext(ctx context.Context, input *comprehendmedical.InferRxNormInput, opts ...request.Option) (*comprehendmedical.InferRxNormOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "InferRxNorm",
		Input:   input,
		Output:  (*comprehendmedical.InferRxNormOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.InferRxNormWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.InferRxNormOutput), req.Error
}

func (c *Client) ListEntitiesDetectionV2JobsWithContext(ctx context.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput, opts ...request.Option) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "ListEntitiesDetectionV2Jobs",
		Input:   input,
		Output:  (*comprehendmedical.ListEntitiesDetectionV2JobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.ListEntitiesDetectionV2JobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.ListEntitiesDetectionV2JobsOutput), req.Error
}

func (c *Client) ListPHIDetectionJobsWithContext(ctx context.Context, input *comprehendmedical.ListPHIDetectionJobsInput, opts ...request.Option) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "ListPHIDetectionJobs",
		Input:   input,
		Output:  (*comprehendmedical.ListPHIDetectionJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.ListPHIDetectionJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.ListPHIDetectionJobsOutput), req.Error
}

func (c *Client) StartEntitiesDetectionV2JobWithContext(ctx context.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput, opts ...request.Option) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "StartEntitiesDetectionV2Job",
		Input:   input,
		Output:  (*comprehendmedical.StartEntitiesDetectionV2JobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.StartEntitiesDetectionV2JobWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.StartEntitiesDetectionV2JobOutput), req.Error
}

func (c *Client) StartPHIDetectionJobWithContext(ctx context.Context, input *comprehendmedical.StartPHIDetectionJobInput, opts ...request.Option) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "StartPHIDetectionJob",
		Input:   input,
		Output:  (*comprehendmedical.StartPHIDetectionJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.StartPHIDetectionJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.StartPHIDetectionJobOutput), req.Error
}

func (c *Client) StopEntitiesDetectionV2JobWithContext(ctx context.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput, opts ...request.Option) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "StopEntitiesDetectionV2Job",
		Input:   input,
		Output:  (*comprehendmedical.StopEntitiesDetectionV2JobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.StopEntitiesDetectionV2JobWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.StopEntitiesDetectionV2JobOutput), req.Error
}

func (c *Client) StopPHIDetectionJobWithContext(ctx context.Context, input *comprehendmedical.StopPHIDetectionJobInput, opts ...request.Option) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "comprehendmedical",
		Action:  "StopPHIDetectionJob",
		Input:   input,
		Output:  (*comprehendmedical.StopPHIDetectionJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ComprehendMedicalAPI.StopPHIDetectionJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*comprehendmedical.StopPHIDetectionJobOutput), req.Error
}
