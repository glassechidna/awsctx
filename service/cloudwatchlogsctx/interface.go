// Code generated by internal/generate/main.go. DO NOT EDIT.

package cloudwatchlogsctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/glassechidna/awsctx"
)

type CloudWatchLogs interface {
	AssociateKmsKeyWithContext(ctx context.Context, input *cloudwatchlogs.AssociateKmsKeyInput, opts ...request.Option) (*cloudwatchlogs.AssociateKmsKeyOutput, error)
	CancelExportTaskWithContext(ctx context.Context, input *cloudwatchlogs.CancelExportTaskInput, opts ...request.Option) (*cloudwatchlogs.CancelExportTaskOutput, error)
	CreateExportTaskWithContext(ctx context.Context, input *cloudwatchlogs.CreateExportTaskInput, opts ...request.Option) (*cloudwatchlogs.CreateExportTaskOutput, error)
	CreateLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.CreateLogGroupInput, opts ...request.Option) (*cloudwatchlogs.CreateLogGroupOutput, error)
	CreateLogStreamWithContext(ctx context.Context, input *cloudwatchlogs.CreateLogStreamInput, opts ...request.Option) (*cloudwatchlogs.CreateLogStreamOutput, error)
	DeleteDestinationWithContext(ctx context.Context, input *cloudwatchlogs.DeleteDestinationInput, opts ...request.Option) (*cloudwatchlogs.DeleteDestinationOutput, error)
	DeleteLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.DeleteLogGroupInput, opts ...request.Option) (*cloudwatchlogs.DeleteLogGroupOutput, error)
	DeleteLogStreamWithContext(ctx context.Context, input *cloudwatchlogs.DeleteLogStreamInput, opts ...request.Option) (*cloudwatchlogs.DeleteLogStreamOutput, error)
	DeleteMetricFilterWithContext(ctx context.Context, input *cloudwatchlogs.DeleteMetricFilterInput, opts ...request.Option) (*cloudwatchlogs.DeleteMetricFilterOutput, error)
	DeleteResourcePolicyWithContext(ctx context.Context, input *cloudwatchlogs.DeleteResourcePolicyInput, opts ...request.Option) (*cloudwatchlogs.DeleteResourcePolicyOutput, error)
	DeleteRetentionPolicyWithContext(ctx context.Context, input *cloudwatchlogs.DeleteRetentionPolicyInput, opts ...request.Option) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error)
	DeleteSubscriptionFilterWithContext(ctx context.Context, input *cloudwatchlogs.DeleteSubscriptionFilterInput, opts ...request.Option) (*cloudwatchlogs.DeleteSubscriptionFilterOutput, error)
	DescribeDestinationsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeDestinationsInput, opts ...request.Option) (*cloudwatchlogs.DescribeDestinationsOutput, error)
	DescribeExportTasksWithContext(ctx context.Context, input *cloudwatchlogs.DescribeExportTasksInput, opts ...request.Option) (*cloudwatchlogs.DescribeExportTasksOutput, error)
	DescribeLogGroupsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeLogGroupsInput, opts ...request.Option) (*cloudwatchlogs.DescribeLogGroupsOutput, error)
	DescribeLogStreamsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeLogStreamsInput, opts ...request.Option) (*cloudwatchlogs.DescribeLogStreamsOutput, error)
	DescribeMetricFiltersWithContext(ctx context.Context, input *cloudwatchlogs.DescribeMetricFiltersInput, opts ...request.Option) (*cloudwatchlogs.DescribeMetricFiltersOutput, error)
	DescribeQueriesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeQueriesInput, opts ...request.Option) (*cloudwatchlogs.DescribeQueriesOutput, error)
	DescribeResourcePoliciesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeResourcePoliciesInput, opts ...request.Option) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error)
	DescribeSubscriptionFiltersWithContext(ctx context.Context, input *cloudwatchlogs.DescribeSubscriptionFiltersInput, opts ...request.Option) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error)
	DisassociateKmsKeyWithContext(ctx context.Context, input *cloudwatchlogs.DisassociateKmsKeyInput, opts ...request.Option) (*cloudwatchlogs.DisassociateKmsKeyOutput, error)
	FilterLogEventsWithContext(ctx context.Context, input *cloudwatchlogs.FilterLogEventsInput, opts ...request.Option) (*cloudwatchlogs.FilterLogEventsOutput, error)
	GetLogEventsWithContext(ctx context.Context, input *cloudwatchlogs.GetLogEventsInput, opts ...request.Option) (*cloudwatchlogs.GetLogEventsOutput, error)
	GetLogGroupFieldsWithContext(ctx context.Context, input *cloudwatchlogs.GetLogGroupFieldsInput, opts ...request.Option) (*cloudwatchlogs.GetLogGroupFieldsOutput, error)
	GetLogRecordWithContext(ctx context.Context, input *cloudwatchlogs.GetLogRecordInput, opts ...request.Option) (*cloudwatchlogs.GetLogRecordOutput, error)
	GetQueryResultsWithContext(ctx context.Context, input *cloudwatchlogs.GetQueryResultsInput, opts ...request.Option) (*cloudwatchlogs.GetQueryResultsOutput, error)
	ListTagsLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.ListTagsLogGroupInput, opts ...request.Option) (*cloudwatchlogs.ListTagsLogGroupOutput, error)
	PutDestinationWithContext(ctx context.Context, input *cloudwatchlogs.PutDestinationInput, opts ...request.Option) (*cloudwatchlogs.PutDestinationOutput, error)
	PutDestinationPolicyWithContext(ctx context.Context, input *cloudwatchlogs.PutDestinationPolicyInput, opts ...request.Option) (*cloudwatchlogs.PutDestinationPolicyOutput, error)
	PutLogEventsWithContext(ctx context.Context, input *cloudwatchlogs.PutLogEventsInput, opts ...request.Option) (*cloudwatchlogs.PutLogEventsOutput, error)
	PutMetricFilterWithContext(ctx context.Context, input *cloudwatchlogs.PutMetricFilterInput, opts ...request.Option) (*cloudwatchlogs.PutMetricFilterOutput, error)
	PutResourcePolicyWithContext(ctx context.Context, input *cloudwatchlogs.PutResourcePolicyInput, opts ...request.Option) (*cloudwatchlogs.PutResourcePolicyOutput, error)
	PutRetentionPolicyWithContext(ctx context.Context, input *cloudwatchlogs.PutRetentionPolicyInput, opts ...request.Option) (*cloudwatchlogs.PutRetentionPolicyOutput, error)
	PutSubscriptionFilterWithContext(ctx context.Context, input *cloudwatchlogs.PutSubscriptionFilterInput, opts ...request.Option) (*cloudwatchlogs.PutSubscriptionFilterOutput, error)
	StartQueryWithContext(ctx context.Context, input *cloudwatchlogs.StartQueryInput, opts ...request.Option) (*cloudwatchlogs.StartQueryOutput, error)
	StopQueryWithContext(ctx context.Context, input *cloudwatchlogs.StopQueryInput, opts ...request.Option) (*cloudwatchlogs.StopQueryOutput, error)
	TagLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.TagLogGroupInput, opts ...request.Option) (*cloudwatchlogs.TagLogGroupOutput, error)
	TestMetricFilterWithContext(ctx context.Context, input *cloudwatchlogs.TestMetricFilterInput, opts ...request.Option) (*cloudwatchlogs.TestMetricFilterOutput, error)
	UntagLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.UntagLogGroupInput, opts ...request.Option) (*cloudwatchlogs.UntagLogGroupOutput, error)
}

type Client struct {
	cloudwatchlogsiface.CloudWatchLogsAPI
	Contexter awsctx.Contexter
}

func New(base cloudwatchlogsiface.CloudWatchLogsAPI, ctxer awsctx.Contexter) CloudWatchLogs {
	return &Client{
		CloudWatchLogsAPI: base,
		Contexter: ctxer,
	}
}

var _ CloudWatchLogs = (*cloudwatchlogs.CloudWatchLogs)(nil)
var _ CloudWatchLogs = (*Client)(nil)

func (c *Client) AssociateKmsKeyWithContext(ctx context.Context, input *cloudwatchlogs.AssociateKmsKeyInput, opts ...request.Option) (*cloudwatchlogs.AssociateKmsKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "AssociateKmsKey",
		Input:   input,
		Output:  (*cloudwatchlogs.AssociateKmsKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.AssociateKmsKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.AssociateKmsKeyOutput), req.Error
}

func (c *Client) CancelExportTaskWithContext(ctx context.Context, input *cloudwatchlogs.CancelExportTaskInput, opts ...request.Option) (*cloudwatchlogs.CancelExportTaskOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "CancelExportTask",
		Input:   input,
		Output:  (*cloudwatchlogs.CancelExportTaskOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.CancelExportTaskWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.CancelExportTaskOutput), req.Error
}

func (c *Client) CreateExportTaskWithContext(ctx context.Context, input *cloudwatchlogs.CreateExportTaskInput, opts ...request.Option) (*cloudwatchlogs.CreateExportTaskOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "CreateExportTask",
		Input:   input,
		Output:  (*cloudwatchlogs.CreateExportTaskOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.CreateExportTaskWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.CreateExportTaskOutput), req.Error
}

func (c *Client) CreateLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.CreateLogGroupInput, opts ...request.Option) (*cloudwatchlogs.CreateLogGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "CreateLogGroup",
		Input:   input,
		Output:  (*cloudwatchlogs.CreateLogGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.CreateLogGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.CreateLogGroupOutput), req.Error
}

func (c *Client) CreateLogStreamWithContext(ctx context.Context, input *cloudwatchlogs.CreateLogStreamInput, opts ...request.Option) (*cloudwatchlogs.CreateLogStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "CreateLogStream",
		Input:   input,
		Output:  (*cloudwatchlogs.CreateLogStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.CreateLogStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.CreateLogStreamOutput), req.Error
}

func (c *Client) DeleteDestinationWithContext(ctx context.Context, input *cloudwatchlogs.DeleteDestinationInput, opts ...request.Option) (*cloudwatchlogs.DeleteDestinationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DeleteDestination",
		Input:   input,
		Output:  (*cloudwatchlogs.DeleteDestinationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DeleteDestinationWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DeleteDestinationOutput), req.Error
}

func (c *Client) DeleteLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.DeleteLogGroupInput, opts ...request.Option) (*cloudwatchlogs.DeleteLogGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DeleteLogGroup",
		Input:   input,
		Output:  (*cloudwatchlogs.DeleteLogGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DeleteLogGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DeleteLogGroupOutput), req.Error
}

func (c *Client) DeleteLogStreamWithContext(ctx context.Context, input *cloudwatchlogs.DeleteLogStreamInput, opts ...request.Option) (*cloudwatchlogs.DeleteLogStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DeleteLogStream",
		Input:   input,
		Output:  (*cloudwatchlogs.DeleteLogStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DeleteLogStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DeleteLogStreamOutput), req.Error
}

func (c *Client) DeleteMetricFilterWithContext(ctx context.Context, input *cloudwatchlogs.DeleteMetricFilterInput, opts ...request.Option) (*cloudwatchlogs.DeleteMetricFilterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DeleteMetricFilter",
		Input:   input,
		Output:  (*cloudwatchlogs.DeleteMetricFilterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DeleteMetricFilterWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DeleteMetricFilterOutput), req.Error
}

func (c *Client) DeleteResourcePolicyWithContext(ctx context.Context, input *cloudwatchlogs.DeleteResourcePolicyInput, opts ...request.Option) (*cloudwatchlogs.DeleteResourcePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DeleteResourcePolicy",
		Input:   input,
		Output:  (*cloudwatchlogs.DeleteResourcePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DeleteResourcePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DeleteResourcePolicyOutput), req.Error
}

func (c *Client) DeleteRetentionPolicyWithContext(ctx context.Context, input *cloudwatchlogs.DeleteRetentionPolicyInput, opts ...request.Option) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DeleteRetentionPolicy",
		Input:   input,
		Output:  (*cloudwatchlogs.DeleteRetentionPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DeleteRetentionPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DeleteRetentionPolicyOutput), req.Error
}

func (c *Client) DeleteSubscriptionFilterWithContext(ctx context.Context, input *cloudwatchlogs.DeleteSubscriptionFilterInput, opts ...request.Option) (*cloudwatchlogs.DeleteSubscriptionFilterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DeleteSubscriptionFilter",
		Input:   input,
		Output:  (*cloudwatchlogs.DeleteSubscriptionFilterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DeleteSubscriptionFilterWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DeleteSubscriptionFilterOutput), req.Error
}

func (c *Client) DescribeDestinationsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeDestinationsInput, opts ...request.Option) (*cloudwatchlogs.DescribeDestinationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DescribeDestinations",
		Input:   input,
		Output:  (*cloudwatchlogs.DescribeDestinationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DescribeDestinationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DescribeDestinationsOutput), req.Error
}

func (c *Client) DescribeExportTasksWithContext(ctx context.Context, input *cloudwatchlogs.DescribeExportTasksInput, opts ...request.Option) (*cloudwatchlogs.DescribeExportTasksOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DescribeExportTasks",
		Input:   input,
		Output:  (*cloudwatchlogs.DescribeExportTasksOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DescribeExportTasksWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DescribeExportTasksOutput), req.Error
}

func (c *Client) DescribeLogGroupsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeLogGroupsInput, opts ...request.Option) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DescribeLogGroups",
		Input:   input,
		Output:  (*cloudwatchlogs.DescribeLogGroupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DescribeLogGroupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DescribeLogGroupsOutput), req.Error
}

func (c *Client) DescribeLogStreamsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeLogStreamsInput, opts ...request.Option) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DescribeLogStreams",
		Input:   input,
		Output:  (*cloudwatchlogs.DescribeLogStreamsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DescribeLogStreamsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DescribeLogStreamsOutput), req.Error
}

func (c *Client) DescribeMetricFiltersWithContext(ctx context.Context, input *cloudwatchlogs.DescribeMetricFiltersInput, opts ...request.Option) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DescribeMetricFilters",
		Input:   input,
		Output:  (*cloudwatchlogs.DescribeMetricFiltersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DescribeMetricFiltersWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DescribeMetricFiltersOutput), req.Error
}

func (c *Client) DescribeQueriesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeQueriesInput, opts ...request.Option) (*cloudwatchlogs.DescribeQueriesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DescribeQueries",
		Input:   input,
		Output:  (*cloudwatchlogs.DescribeQueriesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DescribeQueriesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DescribeQueriesOutput), req.Error
}

func (c *Client) DescribeResourcePoliciesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeResourcePoliciesInput, opts ...request.Option) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DescribeResourcePolicies",
		Input:   input,
		Output:  (*cloudwatchlogs.DescribeResourcePoliciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DescribeResourcePoliciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DescribeResourcePoliciesOutput), req.Error
}

func (c *Client) DescribeSubscriptionFiltersWithContext(ctx context.Context, input *cloudwatchlogs.DescribeSubscriptionFiltersInput, opts ...request.Option) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DescribeSubscriptionFilters",
		Input:   input,
		Output:  (*cloudwatchlogs.DescribeSubscriptionFiltersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DescribeSubscriptionFiltersWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DescribeSubscriptionFiltersOutput), req.Error
}

func (c *Client) DisassociateKmsKeyWithContext(ctx context.Context, input *cloudwatchlogs.DisassociateKmsKeyInput, opts ...request.Option) (*cloudwatchlogs.DisassociateKmsKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "DisassociateKmsKey",
		Input:   input,
		Output:  (*cloudwatchlogs.DisassociateKmsKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.DisassociateKmsKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.DisassociateKmsKeyOutput), req.Error
}

func (c *Client) FilterLogEventsWithContext(ctx context.Context, input *cloudwatchlogs.FilterLogEventsInput, opts ...request.Option) (*cloudwatchlogs.FilterLogEventsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "FilterLogEvents",
		Input:   input,
		Output:  (*cloudwatchlogs.FilterLogEventsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.FilterLogEventsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.FilterLogEventsOutput), req.Error
}

func (c *Client) GetLogEventsWithContext(ctx context.Context, input *cloudwatchlogs.GetLogEventsInput, opts ...request.Option) (*cloudwatchlogs.GetLogEventsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "GetLogEvents",
		Input:   input,
		Output:  (*cloudwatchlogs.GetLogEventsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.GetLogEventsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.GetLogEventsOutput), req.Error
}

func (c *Client) GetLogGroupFieldsWithContext(ctx context.Context, input *cloudwatchlogs.GetLogGroupFieldsInput, opts ...request.Option) (*cloudwatchlogs.GetLogGroupFieldsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "GetLogGroupFields",
		Input:   input,
		Output:  (*cloudwatchlogs.GetLogGroupFieldsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.GetLogGroupFieldsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.GetLogGroupFieldsOutput), req.Error
}

func (c *Client) GetLogRecordWithContext(ctx context.Context, input *cloudwatchlogs.GetLogRecordInput, opts ...request.Option) (*cloudwatchlogs.GetLogRecordOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "GetLogRecord",
		Input:   input,
		Output:  (*cloudwatchlogs.GetLogRecordOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.GetLogRecordWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.GetLogRecordOutput), req.Error
}

func (c *Client) GetQueryResultsWithContext(ctx context.Context, input *cloudwatchlogs.GetQueryResultsInput, opts ...request.Option) (*cloudwatchlogs.GetQueryResultsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "GetQueryResults",
		Input:   input,
		Output:  (*cloudwatchlogs.GetQueryResultsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.GetQueryResultsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.GetQueryResultsOutput), req.Error
}

func (c *Client) ListTagsLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.ListTagsLogGroupInput, opts ...request.Option) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "ListTagsLogGroup",
		Input:   input,
		Output:  (*cloudwatchlogs.ListTagsLogGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.ListTagsLogGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.ListTagsLogGroupOutput), req.Error
}

func (c *Client) PutDestinationWithContext(ctx context.Context, input *cloudwatchlogs.PutDestinationInput, opts ...request.Option) (*cloudwatchlogs.PutDestinationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "PutDestination",
		Input:   input,
		Output:  (*cloudwatchlogs.PutDestinationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.PutDestinationWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.PutDestinationOutput), req.Error
}

func (c *Client) PutDestinationPolicyWithContext(ctx context.Context, input *cloudwatchlogs.PutDestinationPolicyInput, opts ...request.Option) (*cloudwatchlogs.PutDestinationPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "PutDestinationPolicy",
		Input:   input,
		Output:  (*cloudwatchlogs.PutDestinationPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.PutDestinationPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.PutDestinationPolicyOutput), req.Error
}

func (c *Client) PutLogEventsWithContext(ctx context.Context, input *cloudwatchlogs.PutLogEventsInput, opts ...request.Option) (*cloudwatchlogs.PutLogEventsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "PutLogEvents",
		Input:   input,
		Output:  (*cloudwatchlogs.PutLogEventsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.PutLogEventsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.PutLogEventsOutput), req.Error
}

func (c *Client) PutMetricFilterWithContext(ctx context.Context, input *cloudwatchlogs.PutMetricFilterInput, opts ...request.Option) (*cloudwatchlogs.PutMetricFilterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "PutMetricFilter",
		Input:   input,
		Output:  (*cloudwatchlogs.PutMetricFilterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.PutMetricFilterWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.PutMetricFilterOutput), req.Error
}

func (c *Client) PutResourcePolicyWithContext(ctx context.Context, input *cloudwatchlogs.PutResourcePolicyInput, opts ...request.Option) (*cloudwatchlogs.PutResourcePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "PutResourcePolicy",
		Input:   input,
		Output:  (*cloudwatchlogs.PutResourcePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.PutResourcePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.PutResourcePolicyOutput), req.Error
}

func (c *Client) PutRetentionPolicyWithContext(ctx context.Context, input *cloudwatchlogs.PutRetentionPolicyInput, opts ...request.Option) (*cloudwatchlogs.PutRetentionPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "PutRetentionPolicy",
		Input:   input,
		Output:  (*cloudwatchlogs.PutRetentionPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.PutRetentionPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.PutRetentionPolicyOutput), req.Error
}

func (c *Client) PutSubscriptionFilterWithContext(ctx context.Context, input *cloudwatchlogs.PutSubscriptionFilterInput, opts ...request.Option) (*cloudwatchlogs.PutSubscriptionFilterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "PutSubscriptionFilter",
		Input:   input,
		Output:  (*cloudwatchlogs.PutSubscriptionFilterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.PutSubscriptionFilterWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.PutSubscriptionFilterOutput), req.Error
}

func (c *Client) StartQueryWithContext(ctx context.Context, input *cloudwatchlogs.StartQueryInput, opts ...request.Option) (*cloudwatchlogs.StartQueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "StartQuery",
		Input:   input,
		Output:  (*cloudwatchlogs.StartQueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.StartQueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.StartQueryOutput), req.Error
}

func (c *Client) StopQueryWithContext(ctx context.Context, input *cloudwatchlogs.StopQueryInput, opts ...request.Option) (*cloudwatchlogs.StopQueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "StopQuery",
		Input:   input,
		Output:  (*cloudwatchlogs.StopQueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.StopQueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.StopQueryOutput), req.Error
}

func (c *Client) TagLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.TagLogGroupInput, opts ...request.Option) (*cloudwatchlogs.TagLogGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "TagLogGroup",
		Input:   input,
		Output:  (*cloudwatchlogs.TagLogGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.TagLogGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.TagLogGroupOutput), req.Error
}

func (c *Client) TestMetricFilterWithContext(ctx context.Context, input *cloudwatchlogs.TestMetricFilterInput, opts ...request.Option) (*cloudwatchlogs.TestMetricFilterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "TestMetricFilter",
		Input:   input,
		Output:  (*cloudwatchlogs.TestMetricFilterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.TestMetricFilterWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.TestMetricFilterOutput), req.Error
}

func (c *Client) UntagLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.UntagLogGroupInput, opts ...request.Option) (*cloudwatchlogs.UntagLogGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatchlogs",
		Action:  "UntagLogGroup",
		Input:   input,
		Output:  (*cloudwatchlogs.UntagLogGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchLogsAPI.UntagLogGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatchlogs.UntagLogGroupOutput), req.Error
}
