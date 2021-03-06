// Code generated by internal/generate/main.go. DO NOT EDIT.

package codepipelinectx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/codepipeline/codepipelineiface"
	"github.com/glassechidna/awsctx"
)

type CodePipeline interface {
	AcknowledgeJobWithContext(ctx context.Context, input *codepipeline.AcknowledgeJobInput, opts ...request.Option) (*codepipeline.AcknowledgeJobOutput, error)
	AcknowledgeThirdPartyJobWithContext(ctx context.Context, input *codepipeline.AcknowledgeThirdPartyJobInput, opts ...request.Option) (*codepipeline.AcknowledgeThirdPartyJobOutput, error)
	CreateCustomActionTypeWithContext(ctx context.Context, input *codepipeline.CreateCustomActionTypeInput, opts ...request.Option) (*codepipeline.CreateCustomActionTypeOutput, error)
	CreatePipelineWithContext(ctx context.Context, input *codepipeline.CreatePipelineInput, opts ...request.Option) (*codepipeline.CreatePipelineOutput, error)
	DeleteCustomActionTypeWithContext(ctx context.Context, input *codepipeline.DeleteCustomActionTypeInput, opts ...request.Option) (*codepipeline.DeleteCustomActionTypeOutput, error)
	DeletePipelineWithContext(ctx context.Context, input *codepipeline.DeletePipelineInput, opts ...request.Option) (*codepipeline.DeletePipelineOutput, error)
	DeleteWebhookWithContext(ctx context.Context, input *codepipeline.DeleteWebhookInput, opts ...request.Option) (*codepipeline.DeleteWebhookOutput, error)
	DeregisterWebhookWithThirdPartyWithContext(ctx context.Context, input *codepipeline.DeregisterWebhookWithThirdPartyInput, opts ...request.Option) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error)
	DisableStageTransitionWithContext(ctx context.Context, input *codepipeline.DisableStageTransitionInput, opts ...request.Option) (*codepipeline.DisableStageTransitionOutput, error)
	EnableStageTransitionWithContext(ctx context.Context, input *codepipeline.EnableStageTransitionInput, opts ...request.Option) (*codepipeline.EnableStageTransitionOutput, error)
	GetActionTypeWithContext(ctx context.Context, input *codepipeline.GetActionTypeInput, opts ...request.Option) (*codepipeline.GetActionTypeOutput, error)
	GetJobDetailsWithContext(ctx context.Context, input *codepipeline.GetJobDetailsInput, opts ...request.Option) (*codepipeline.GetJobDetailsOutput, error)
	GetPipelineWithContext(ctx context.Context, input *codepipeline.GetPipelineInput, opts ...request.Option) (*codepipeline.GetPipelineOutput, error)
	GetPipelineExecutionWithContext(ctx context.Context, input *codepipeline.GetPipelineExecutionInput, opts ...request.Option) (*codepipeline.GetPipelineExecutionOutput, error)
	GetPipelineStateWithContext(ctx context.Context, input *codepipeline.GetPipelineStateInput, opts ...request.Option) (*codepipeline.GetPipelineStateOutput, error)
	GetThirdPartyJobDetailsWithContext(ctx context.Context, input *codepipeline.GetThirdPartyJobDetailsInput, opts ...request.Option) (*codepipeline.GetThirdPartyJobDetailsOutput, error)
	ListActionExecutionsWithContext(ctx context.Context, input *codepipeline.ListActionExecutionsInput, opts ...request.Option) (*codepipeline.ListActionExecutionsOutput, error)
	ListActionExecutionsPagesWithContext(ctx context.Context, input *codepipeline.ListActionExecutionsInput, cb func(*codepipeline.ListActionExecutionsOutput, bool) bool, opts ...request.Option) error
	ListActionTypesWithContext(ctx context.Context, input *codepipeline.ListActionTypesInput, opts ...request.Option) (*codepipeline.ListActionTypesOutput, error)
	ListActionTypesPagesWithContext(ctx context.Context, input *codepipeline.ListActionTypesInput, cb func(*codepipeline.ListActionTypesOutput, bool) bool, opts ...request.Option) error
	ListPipelineExecutionsWithContext(ctx context.Context, input *codepipeline.ListPipelineExecutionsInput, opts ...request.Option) (*codepipeline.ListPipelineExecutionsOutput, error)
	ListPipelineExecutionsPagesWithContext(ctx context.Context, input *codepipeline.ListPipelineExecutionsInput, cb func(*codepipeline.ListPipelineExecutionsOutput, bool) bool, opts ...request.Option) error
	ListPipelinesWithContext(ctx context.Context, input *codepipeline.ListPipelinesInput, opts ...request.Option) (*codepipeline.ListPipelinesOutput, error)
	ListPipelinesPagesWithContext(ctx context.Context, input *codepipeline.ListPipelinesInput, cb func(*codepipeline.ListPipelinesOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *codepipeline.ListTagsForResourceInput, opts ...request.Option) (*codepipeline.ListTagsForResourceOutput, error)
	ListTagsForResourcePagesWithContext(ctx context.Context, input *codepipeline.ListTagsForResourceInput, cb func(*codepipeline.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error
	ListWebhooksWithContext(ctx context.Context, input *codepipeline.ListWebhooksInput, opts ...request.Option) (*codepipeline.ListWebhooksOutput, error)
	ListWebhooksPagesWithContext(ctx context.Context, input *codepipeline.ListWebhooksInput, cb func(*codepipeline.ListWebhooksOutput, bool) bool, opts ...request.Option) error
	PollForJobsWithContext(ctx context.Context, input *codepipeline.PollForJobsInput, opts ...request.Option) (*codepipeline.PollForJobsOutput, error)
	PollForThirdPartyJobsWithContext(ctx context.Context, input *codepipeline.PollForThirdPartyJobsInput, opts ...request.Option) (*codepipeline.PollForThirdPartyJobsOutput, error)
	PutActionRevisionWithContext(ctx context.Context, input *codepipeline.PutActionRevisionInput, opts ...request.Option) (*codepipeline.PutActionRevisionOutput, error)
	PutApprovalResultWithContext(ctx context.Context, input *codepipeline.PutApprovalResultInput, opts ...request.Option) (*codepipeline.PutApprovalResultOutput, error)
	PutJobFailureResultWithContext(ctx context.Context, input *codepipeline.PutJobFailureResultInput, opts ...request.Option) (*codepipeline.PutJobFailureResultOutput, error)
	PutJobSuccessResultWithContext(ctx context.Context, input *codepipeline.PutJobSuccessResultInput, opts ...request.Option) (*codepipeline.PutJobSuccessResultOutput, error)
	PutThirdPartyJobFailureResultWithContext(ctx context.Context, input *codepipeline.PutThirdPartyJobFailureResultInput, opts ...request.Option) (*codepipeline.PutThirdPartyJobFailureResultOutput, error)
	PutThirdPartyJobSuccessResultWithContext(ctx context.Context, input *codepipeline.PutThirdPartyJobSuccessResultInput, opts ...request.Option) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error)
	PutWebhookWithContext(ctx context.Context, input *codepipeline.PutWebhookInput, opts ...request.Option) (*codepipeline.PutWebhookOutput, error)
	RegisterWebhookWithThirdPartyWithContext(ctx context.Context, input *codepipeline.RegisterWebhookWithThirdPartyInput, opts ...request.Option) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error)
	RetryStageExecutionWithContext(ctx context.Context, input *codepipeline.RetryStageExecutionInput, opts ...request.Option) (*codepipeline.RetryStageExecutionOutput, error)
	StartPipelineExecutionWithContext(ctx context.Context, input *codepipeline.StartPipelineExecutionInput, opts ...request.Option) (*codepipeline.StartPipelineExecutionOutput, error)
	StopPipelineExecutionWithContext(ctx context.Context, input *codepipeline.StopPipelineExecutionInput, opts ...request.Option) (*codepipeline.StopPipelineExecutionOutput, error)
	TagResourceWithContext(ctx context.Context, input *codepipeline.TagResourceInput, opts ...request.Option) (*codepipeline.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *codepipeline.UntagResourceInput, opts ...request.Option) (*codepipeline.UntagResourceOutput, error)
	UpdateActionTypeWithContext(ctx context.Context, input *codepipeline.UpdateActionTypeInput, opts ...request.Option) (*codepipeline.UpdateActionTypeOutput, error)
	UpdatePipelineWithContext(ctx context.Context, input *codepipeline.UpdatePipelineInput, opts ...request.Option) (*codepipeline.UpdatePipelineOutput, error)
}

type Client struct {
	codepipelineiface.CodePipelineAPI
	Contexter awsctx.Contexter
}

func New(base codepipelineiface.CodePipelineAPI, ctxer awsctx.Contexter) CodePipeline {
	return &Client{
		CodePipelineAPI: base,
		Contexter: ctxer,
	}
}

var _ CodePipeline = (*codepipeline.CodePipeline)(nil)
var _ CodePipeline = (*Client)(nil)

func (c *Client) AcknowledgeJobWithContext(ctx context.Context, input *codepipeline.AcknowledgeJobInput, opts ...request.Option) (*codepipeline.AcknowledgeJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "AcknowledgeJob",
		Input:   input,
		Output:  (*codepipeline.AcknowledgeJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.AcknowledgeJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.AcknowledgeJobOutput), req.Error
}

func (c *Client) AcknowledgeThirdPartyJobWithContext(ctx context.Context, input *codepipeline.AcknowledgeThirdPartyJobInput, opts ...request.Option) (*codepipeline.AcknowledgeThirdPartyJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "AcknowledgeThirdPartyJob",
		Input:   input,
		Output:  (*codepipeline.AcknowledgeThirdPartyJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.AcknowledgeThirdPartyJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.AcknowledgeThirdPartyJobOutput), req.Error
}

func (c *Client) CreateCustomActionTypeWithContext(ctx context.Context, input *codepipeline.CreateCustomActionTypeInput, opts ...request.Option) (*codepipeline.CreateCustomActionTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "CreateCustomActionType",
		Input:   input,
		Output:  (*codepipeline.CreateCustomActionTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.CreateCustomActionTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.CreateCustomActionTypeOutput), req.Error
}

func (c *Client) CreatePipelineWithContext(ctx context.Context, input *codepipeline.CreatePipelineInput, opts ...request.Option) (*codepipeline.CreatePipelineOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "CreatePipeline",
		Input:   input,
		Output:  (*codepipeline.CreatePipelineOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.CreatePipelineWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.CreatePipelineOutput), req.Error
}

func (c *Client) DeleteCustomActionTypeWithContext(ctx context.Context, input *codepipeline.DeleteCustomActionTypeInput, opts ...request.Option) (*codepipeline.DeleteCustomActionTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "DeleteCustomActionType",
		Input:   input,
		Output:  (*codepipeline.DeleteCustomActionTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.DeleteCustomActionTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.DeleteCustomActionTypeOutput), req.Error
}

func (c *Client) DeletePipelineWithContext(ctx context.Context, input *codepipeline.DeletePipelineInput, opts ...request.Option) (*codepipeline.DeletePipelineOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "DeletePipeline",
		Input:   input,
		Output:  (*codepipeline.DeletePipelineOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.DeletePipelineWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.DeletePipelineOutput), req.Error
}

func (c *Client) DeleteWebhookWithContext(ctx context.Context, input *codepipeline.DeleteWebhookInput, opts ...request.Option) (*codepipeline.DeleteWebhookOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "DeleteWebhook",
		Input:   input,
		Output:  (*codepipeline.DeleteWebhookOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.DeleteWebhookWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.DeleteWebhookOutput), req.Error
}

func (c *Client) DeregisterWebhookWithThirdPartyWithContext(ctx context.Context, input *codepipeline.DeregisterWebhookWithThirdPartyInput, opts ...request.Option) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "DeregisterWebhookWithThirdParty",
		Input:   input,
		Output:  (*codepipeline.DeregisterWebhookWithThirdPartyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.DeregisterWebhookWithThirdPartyWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.DeregisterWebhookWithThirdPartyOutput), req.Error
}

func (c *Client) DisableStageTransitionWithContext(ctx context.Context, input *codepipeline.DisableStageTransitionInput, opts ...request.Option) (*codepipeline.DisableStageTransitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "DisableStageTransition",
		Input:   input,
		Output:  (*codepipeline.DisableStageTransitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.DisableStageTransitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.DisableStageTransitionOutput), req.Error
}

func (c *Client) EnableStageTransitionWithContext(ctx context.Context, input *codepipeline.EnableStageTransitionInput, opts ...request.Option) (*codepipeline.EnableStageTransitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "EnableStageTransition",
		Input:   input,
		Output:  (*codepipeline.EnableStageTransitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.EnableStageTransitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.EnableStageTransitionOutput), req.Error
}

func (c *Client) GetActionTypeWithContext(ctx context.Context, input *codepipeline.GetActionTypeInput, opts ...request.Option) (*codepipeline.GetActionTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "GetActionType",
		Input:   input,
		Output:  (*codepipeline.GetActionTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.GetActionTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.GetActionTypeOutput), req.Error
}

func (c *Client) GetJobDetailsWithContext(ctx context.Context, input *codepipeline.GetJobDetailsInput, opts ...request.Option) (*codepipeline.GetJobDetailsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "GetJobDetails",
		Input:   input,
		Output:  (*codepipeline.GetJobDetailsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.GetJobDetailsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.GetJobDetailsOutput), req.Error
}

func (c *Client) GetPipelineWithContext(ctx context.Context, input *codepipeline.GetPipelineInput, opts ...request.Option) (*codepipeline.GetPipelineOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "GetPipeline",
		Input:   input,
		Output:  (*codepipeline.GetPipelineOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.GetPipelineWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.GetPipelineOutput), req.Error
}

func (c *Client) GetPipelineExecutionWithContext(ctx context.Context, input *codepipeline.GetPipelineExecutionInput, opts ...request.Option) (*codepipeline.GetPipelineExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "GetPipelineExecution",
		Input:   input,
		Output:  (*codepipeline.GetPipelineExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.GetPipelineExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.GetPipelineExecutionOutput), req.Error
}

func (c *Client) GetPipelineStateWithContext(ctx context.Context, input *codepipeline.GetPipelineStateInput, opts ...request.Option) (*codepipeline.GetPipelineStateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "GetPipelineState",
		Input:   input,
		Output:  (*codepipeline.GetPipelineStateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.GetPipelineStateWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.GetPipelineStateOutput), req.Error
}

func (c *Client) GetThirdPartyJobDetailsWithContext(ctx context.Context, input *codepipeline.GetThirdPartyJobDetailsInput, opts ...request.Option) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "GetThirdPartyJobDetails",
		Input:   input,
		Output:  (*codepipeline.GetThirdPartyJobDetailsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.GetThirdPartyJobDetailsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.GetThirdPartyJobDetailsOutput), req.Error
}

func (c *Client) ListActionExecutionsWithContext(ctx context.Context, input *codepipeline.ListActionExecutionsInput, opts ...request.Option) (*codepipeline.ListActionExecutionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListActionExecutions",
		Input:   input,
		Output:  (*codepipeline.ListActionExecutionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.ListActionExecutionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.ListActionExecutionsOutput), req.Error
}

func (c *Client) ListActionExecutionsPagesWithContext(ctx context.Context, input *codepipeline.ListActionExecutionsInput, cb func(*codepipeline.ListActionExecutionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListActionExecutions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CodePipelineAPI.ListActionExecutionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListActionTypesWithContext(ctx context.Context, input *codepipeline.ListActionTypesInput, opts ...request.Option) (*codepipeline.ListActionTypesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListActionTypes",
		Input:   input,
		Output:  (*codepipeline.ListActionTypesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.ListActionTypesWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.ListActionTypesOutput), req.Error
}

func (c *Client) ListActionTypesPagesWithContext(ctx context.Context, input *codepipeline.ListActionTypesInput, cb func(*codepipeline.ListActionTypesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListActionTypes",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CodePipelineAPI.ListActionTypesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListPipelineExecutionsWithContext(ctx context.Context, input *codepipeline.ListPipelineExecutionsInput, opts ...request.Option) (*codepipeline.ListPipelineExecutionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListPipelineExecutions",
		Input:   input,
		Output:  (*codepipeline.ListPipelineExecutionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.ListPipelineExecutionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.ListPipelineExecutionsOutput), req.Error
}

func (c *Client) ListPipelineExecutionsPagesWithContext(ctx context.Context, input *codepipeline.ListPipelineExecutionsInput, cb func(*codepipeline.ListPipelineExecutionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListPipelineExecutions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CodePipelineAPI.ListPipelineExecutionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListPipelinesWithContext(ctx context.Context, input *codepipeline.ListPipelinesInput, opts ...request.Option) (*codepipeline.ListPipelinesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListPipelines",
		Input:   input,
		Output:  (*codepipeline.ListPipelinesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.ListPipelinesWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.ListPipelinesOutput), req.Error
}

func (c *Client) ListPipelinesPagesWithContext(ctx context.Context, input *codepipeline.ListPipelinesInput, cb func(*codepipeline.ListPipelinesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListPipelines",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CodePipelineAPI.ListPipelinesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *codepipeline.ListTagsForResourceInput, opts ...request.Option) (*codepipeline.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*codepipeline.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTagsForResourcePagesWithContext(ctx context.Context, input *codepipeline.ListTagsForResourceInput, cb func(*codepipeline.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListTagsForResource",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CodePipelineAPI.ListTagsForResourcePagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListWebhooksWithContext(ctx context.Context, input *codepipeline.ListWebhooksInput, opts ...request.Option) (*codepipeline.ListWebhooksOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListWebhooks",
		Input:   input,
		Output:  (*codepipeline.ListWebhooksOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.ListWebhooksWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.ListWebhooksOutput), req.Error
}

func (c *Client) ListWebhooksPagesWithContext(ctx context.Context, input *codepipeline.ListWebhooksInput, cb func(*codepipeline.ListWebhooksOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "ListWebhooks",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CodePipelineAPI.ListWebhooksPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) PollForJobsWithContext(ctx context.Context, input *codepipeline.PollForJobsInput, opts ...request.Option) (*codepipeline.PollForJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "PollForJobs",
		Input:   input,
		Output:  (*codepipeline.PollForJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.PollForJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.PollForJobsOutput), req.Error
}

func (c *Client) PollForThirdPartyJobsWithContext(ctx context.Context, input *codepipeline.PollForThirdPartyJobsInput, opts ...request.Option) (*codepipeline.PollForThirdPartyJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "PollForThirdPartyJobs",
		Input:   input,
		Output:  (*codepipeline.PollForThirdPartyJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.PollForThirdPartyJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.PollForThirdPartyJobsOutput), req.Error
}

func (c *Client) PutActionRevisionWithContext(ctx context.Context, input *codepipeline.PutActionRevisionInput, opts ...request.Option) (*codepipeline.PutActionRevisionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "PutActionRevision",
		Input:   input,
		Output:  (*codepipeline.PutActionRevisionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.PutActionRevisionWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.PutActionRevisionOutput), req.Error
}

func (c *Client) PutApprovalResultWithContext(ctx context.Context, input *codepipeline.PutApprovalResultInput, opts ...request.Option) (*codepipeline.PutApprovalResultOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "PutApprovalResult",
		Input:   input,
		Output:  (*codepipeline.PutApprovalResultOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.PutApprovalResultWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.PutApprovalResultOutput), req.Error
}

func (c *Client) PutJobFailureResultWithContext(ctx context.Context, input *codepipeline.PutJobFailureResultInput, opts ...request.Option) (*codepipeline.PutJobFailureResultOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "PutJobFailureResult",
		Input:   input,
		Output:  (*codepipeline.PutJobFailureResultOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.PutJobFailureResultWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.PutJobFailureResultOutput), req.Error
}

func (c *Client) PutJobSuccessResultWithContext(ctx context.Context, input *codepipeline.PutJobSuccessResultInput, opts ...request.Option) (*codepipeline.PutJobSuccessResultOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "PutJobSuccessResult",
		Input:   input,
		Output:  (*codepipeline.PutJobSuccessResultOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.PutJobSuccessResultWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.PutJobSuccessResultOutput), req.Error
}

func (c *Client) PutThirdPartyJobFailureResultWithContext(ctx context.Context, input *codepipeline.PutThirdPartyJobFailureResultInput, opts ...request.Option) (*codepipeline.PutThirdPartyJobFailureResultOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "PutThirdPartyJobFailureResult",
		Input:   input,
		Output:  (*codepipeline.PutThirdPartyJobFailureResultOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.PutThirdPartyJobFailureResultWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.PutThirdPartyJobFailureResultOutput), req.Error
}

func (c *Client) PutThirdPartyJobSuccessResultWithContext(ctx context.Context, input *codepipeline.PutThirdPartyJobSuccessResultInput, opts ...request.Option) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "PutThirdPartyJobSuccessResult",
		Input:   input,
		Output:  (*codepipeline.PutThirdPartyJobSuccessResultOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.PutThirdPartyJobSuccessResultWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.PutThirdPartyJobSuccessResultOutput), req.Error
}

func (c *Client) PutWebhookWithContext(ctx context.Context, input *codepipeline.PutWebhookInput, opts ...request.Option) (*codepipeline.PutWebhookOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "PutWebhook",
		Input:   input,
		Output:  (*codepipeline.PutWebhookOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.PutWebhookWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.PutWebhookOutput), req.Error
}

func (c *Client) RegisterWebhookWithThirdPartyWithContext(ctx context.Context, input *codepipeline.RegisterWebhookWithThirdPartyInput, opts ...request.Option) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "RegisterWebhookWithThirdParty",
		Input:   input,
		Output:  (*codepipeline.RegisterWebhookWithThirdPartyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.RegisterWebhookWithThirdPartyWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.RegisterWebhookWithThirdPartyOutput), req.Error
}

func (c *Client) RetryStageExecutionWithContext(ctx context.Context, input *codepipeline.RetryStageExecutionInput, opts ...request.Option) (*codepipeline.RetryStageExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "RetryStageExecution",
		Input:   input,
		Output:  (*codepipeline.RetryStageExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.RetryStageExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.RetryStageExecutionOutput), req.Error
}

func (c *Client) StartPipelineExecutionWithContext(ctx context.Context, input *codepipeline.StartPipelineExecutionInput, opts ...request.Option) (*codepipeline.StartPipelineExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "StartPipelineExecution",
		Input:   input,
		Output:  (*codepipeline.StartPipelineExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.StartPipelineExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.StartPipelineExecutionOutput), req.Error
}

func (c *Client) StopPipelineExecutionWithContext(ctx context.Context, input *codepipeline.StopPipelineExecutionInput, opts ...request.Option) (*codepipeline.StopPipelineExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "StopPipelineExecution",
		Input:   input,
		Output:  (*codepipeline.StopPipelineExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.StopPipelineExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.StopPipelineExecutionOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *codepipeline.TagResourceInput, opts ...request.Option) (*codepipeline.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "TagResource",
		Input:   input,
		Output:  (*codepipeline.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *codepipeline.UntagResourceInput, opts ...request.Option) (*codepipeline.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*codepipeline.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.UntagResourceOutput), req.Error
}

func (c *Client) UpdateActionTypeWithContext(ctx context.Context, input *codepipeline.UpdateActionTypeInput, opts ...request.Option) (*codepipeline.UpdateActionTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "UpdateActionType",
		Input:   input,
		Output:  (*codepipeline.UpdateActionTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.UpdateActionTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.UpdateActionTypeOutput), req.Error
}

func (c *Client) UpdatePipelineWithContext(ctx context.Context, input *codepipeline.UpdatePipelineInput, opts ...request.Option) (*codepipeline.UpdatePipelineOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codepipeline",
		Action:  "UpdatePipeline",
		Input:   input,
		Output:  (*codepipeline.UpdatePipelineOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodePipelineAPI.UpdatePipelineWithContext(ctx, input, opts...)
	})

	return req.Output.(*codepipeline.UpdatePipelineOutput), req.Error
}
