// Code generated by internal/generate/main.go. DO NOT EDIT.

package cloudwatchctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/glassechidna/awsctx"
)

type CloudWatch interface {
	DeleteAlarmsWithContext(ctx context.Context, input *cloudwatch.DeleteAlarmsInput, opts ...request.Option) (*cloudwatch.DeleteAlarmsOutput, error)
	DeleteAnomalyDetectorWithContext(ctx context.Context, input *cloudwatch.DeleteAnomalyDetectorInput, opts ...request.Option) (*cloudwatch.DeleteAnomalyDetectorOutput, error)
	DeleteDashboardsWithContext(ctx context.Context, input *cloudwatch.DeleteDashboardsInput, opts ...request.Option) (*cloudwatch.DeleteDashboardsOutput, error)
	DeleteInsightRulesWithContext(ctx context.Context, input *cloudwatch.DeleteInsightRulesInput, opts ...request.Option) (*cloudwatch.DeleteInsightRulesOutput, error)
	DeleteMetricStreamWithContext(ctx context.Context, input *cloudwatch.DeleteMetricStreamInput, opts ...request.Option) (*cloudwatch.DeleteMetricStreamOutput, error)
	DescribeAlarmHistoryWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmHistoryInput, opts ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarmHistoryPagesWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmHistoryInput, cb func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool, opts ...request.Option) error
	DescribeAlarmsWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmsInput, opts ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsPagesWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmsInput, cb func(*cloudwatch.DescribeAlarmsOutput, bool) bool, opts ...request.Option) error
	DescribeAlarmsForMetricWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmsForMetricInput, opts ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAnomalyDetectorsWithContext(ctx context.Context, input *cloudwatch.DescribeAnomalyDetectorsInput, opts ...request.Option) (*cloudwatch.DescribeAnomalyDetectorsOutput, error)
	DescribeAnomalyDetectorsPagesWithContext(ctx context.Context, input *cloudwatch.DescribeAnomalyDetectorsInput, cb func(*cloudwatch.DescribeAnomalyDetectorsOutput, bool) bool, opts ...request.Option) error
	DescribeInsightRulesWithContext(ctx context.Context, input *cloudwatch.DescribeInsightRulesInput, opts ...request.Option) (*cloudwatch.DescribeInsightRulesOutput, error)
	DescribeInsightRulesPagesWithContext(ctx context.Context, input *cloudwatch.DescribeInsightRulesInput, cb func(*cloudwatch.DescribeInsightRulesOutput, bool) bool, opts ...request.Option) error
	DisableAlarmActionsWithContext(ctx context.Context, input *cloudwatch.DisableAlarmActionsInput, opts ...request.Option) (*cloudwatch.DisableAlarmActionsOutput, error)
	DisableInsightRulesWithContext(ctx context.Context, input *cloudwatch.DisableInsightRulesInput, opts ...request.Option) (*cloudwatch.DisableInsightRulesOutput, error)
	EnableAlarmActionsWithContext(ctx context.Context, input *cloudwatch.EnableAlarmActionsInput, opts ...request.Option) (*cloudwatch.EnableAlarmActionsOutput, error)
	EnableInsightRulesWithContext(ctx context.Context, input *cloudwatch.EnableInsightRulesInput, opts ...request.Option) (*cloudwatch.EnableInsightRulesOutput, error)
	GetDashboardWithContext(ctx context.Context, input *cloudwatch.GetDashboardInput, opts ...request.Option) (*cloudwatch.GetDashboardOutput, error)
	GetInsightRuleReportWithContext(ctx context.Context, input *cloudwatch.GetInsightRuleReportInput, opts ...request.Option) (*cloudwatch.GetInsightRuleReportOutput, error)
	GetMetricDataWithContext(ctx context.Context, input *cloudwatch.GetMetricDataInput, opts ...request.Option) (*cloudwatch.GetMetricDataOutput, error)
	GetMetricDataPagesWithContext(ctx context.Context, input *cloudwatch.GetMetricDataInput, cb func(*cloudwatch.GetMetricDataOutput, bool) bool, opts ...request.Option) error
	GetMetricStatisticsWithContext(ctx context.Context, input *cloudwatch.GetMetricStatisticsInput, opts ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStreamWithContext(ctx context.Context, input *cloudwatch.GetMetricStreamInput, opts ...request.Option) (*cloudwatch.GetMetricStreamOutput, error)
	GetMetricWidgetImageWithContext(ctx context.Context, input *cloudwatch.GetMetricWidgetImageInput, opts ...request.Option) (*cloudwatch.GetMetricWidgetImageOutput, error)
	ListDashboardsWithContext(ctx context.Context, input *cloudwatch.ListDashboardsInput, opts ...request.Option) (*cloudwatch.ListDashboardsOutput, error)
	ListDashboardsPagesWithContext(ctx context.Context, input *cloudwatch.ListDashboardsInput, cb func(*cloudwatch.ListDashboardsOutput, bool) bool, opts ...request.Option) error
	ListManagedInsightRulesWithContext(ctx context.Context, input *cloudwatch.ListManagedInsightRulesInput, opts ...request.Option) (*cloudwatch.ListManagedInsightRulesOutput, error)
	ListManagedInsightRulesPagesWithContext(ctx context.Context, input *cloudwatch.ListManagedInsightRulesInput, cb func(*cloudwatch.ListManagedInsightRulesOutput, bool) bool, opts ...request.Option) error
	ListMetricStreamsWithContext(ctx context.Context, input *cloudwatch.ListMetricStreamsInput, opts ...request.Option) (*cloudwatch.ListMetricStreamsOutput, error)
	ListMetricStreamsPagesWithContext(ctx context.Context, input *cloudwatch.ListMetricStreamsInput, cb func(*cloudwatch.ListMetricStreamsOutput, bool) bool, opts ...request.Option) error
	ListMetricsWithContext(ctx context.Context, input *cloudwatch.ListMetricsInput, opts ...request.Option) (*cloudwatch.ListMetricsOutput, error)
	ListMetricsPagesWithContext(ctx context.Context, input *cloudwatch.ListMetricsInput, cb func(*cloudwatch.ListMetricsOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *cloudwatch.ListTagsForResourceInput, opts ...request.Option) (*cloudwatch.ListTagsForResourceOutput, error)
	PutAnomalyDetectorWithContext(ctx context.Context, input *cloudwatch.PutAnomalyDetectorInput, opts ...request.Option) (*cloudwatch.PutAnomalyDetectorOutput, error)
	PutCompositeAlarmWithContext(ctx context.Context, input *cloudwatch.PutCompositeAlarmInput, opts ...request.Option) (*cloudwatch.PutCompositeAlarmOutput, error)
	PutDashboardWithContext(ctx context.Context, input *cloudwatch.PutDashboardInput, opts ...request.Option) (*cloudwatch.PutDashboardOutput, error)
	PutInsightRuleWithContext(ctx context.Context, input *cloudwatch.PutInsightRuleInput, opts ...request.Option) (*cloudwatch.PutInsightRuleOutput, error)
	PutManagedInsightRulesWithContext(ctx context.Context, input *cloudwatch.PutManagedInsightRulesInput, opts ...request.Option) (*cloudwatch.PutManagedInsightRulesOutput, error)
	PutMetricAlarmWithContext(ctx context.Context, input *cloudwatch.PutMetricAlarmInput, opts ...request.Option) (*cloudwatch.PutMetricAlarmOutput, error)
	PutMetricDataWithContext(ctx context.Context, input *cloudwatch.PutMetricDataInput, opts ...request.Option) (*cloudwatch.PutMetricDataOutput, error)
	PutMetricStreamWithContext(ctx context.Context, input *cloudwatch.PutMetricStreamInput, opts ...request.Option) (*cloudwatch.PutMetricStreamOutput, error)
	SetAlarmStateWithContext(ctx context.Context, input *cloudwatch.SetAlarmStateInput, opts ...request.Option) (*cloudwatch.SetAlarmStateOutput, error)
	StartMetricStreamsWithContext(ctx context.Context, input *cloudwatch.StartMetricStreamsInput, opts ...request.Option) (*cloudwatch.StartMetricStreamsOutput, error)
	StopMetricStreamsWithContext(ctx context.Context, input *cloudwatch.StopMetricStreamsInput, opts ...request.Option) (*cloudwatch.StopMetricStreamsOutput, error)
	TagResourceWithContext(ctx context.Context, input *cloudwatch.TagResourceInput, opts ...request.Option) (*cloudwatch.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *cloudwatch.UntagResourceInput, opts ...request.Option) (*cloudwatch.UntagResourceOutput, error)
}

type Client struct {
	cloudwatchiface.CloudWatchAPI
	Contexter awsctx.Contexter
}

func New(base cloudwatchiface.CloudWatchAPI, ctxer awsctx.Contexter) CloudWatch {
	return &Client{
		CloudWatchAPI: base,
		Contexter: ctxer,
	}
}

var _ CloudWatch = (*cloudwatch.CloudWatch)(nil)
var _ CloudWatch = (*Client)(nil)

func (c *Client) DeleteAlarmsWithContext(ctx context.Context, input *cloudwatch.DeleteAlarmsInput, opts ...request.Option) (*cloudwatch.DeleteAlarmsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DeleteAlarms",
		Input:   input,
		Output:  (*cloudwatch.DeleteAlarmsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DeleteAlarmsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DeleteAlarmsOutput), req.Error
}

func (c *Client) DeleteAnomalyDetectorWithContext(ctx context.Context, input *cloudwatch.DeleteAnomalyDetectorInput, opts ...request.Option) (*cloudwatch.DeleteAnomalyDetectorOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DeleteAnomalyDetector",
		Input:   input,
		Output:  (*cloudwatch.DeleteAnomalyDetectorOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DeleteAnomalyDetectorWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DeleteAnomalyDetectorOutput), req.Error
}

func (c *Client) DeleteDashboardsWithContext(ctx context.Context, input *cloudwatch.DeleteDashboardsInput, opts ...request.Option) (*cloudwatch.DeleteDashboardsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DeleteDashboards",
		Input:   input,
		Output:  (*cloudwatch.DeleteDashboardsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DeleteDashboardsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DeleteDashboardsOutput), req.Error
}

func (c *Client) DeleteInsightRulesWithContext(ctx context.Context, input *cloudwatch.DeleteInsightRulesInput, opts ...request.Option) (*cloudwatch.DeleteInsightRulesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DeleteInsightRules",
		Input:   input,
		Output:  (*cloudwatch.DeleteInsightRulesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DeleteInsightRulesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DeleteInsightRulesOutput), req.Error
}

func (c *Client) DeleteMetricStreamWithContext(ctx context.Context, input *cloudwatch.DeleteMetricStreamInput, opts ...request.Option) (*cloudwatch.DeleteMetricStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DeleteMetricStream",
		Input:   input,
		Output:  (*cloudwatch.DeleteMetricStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DeleteMetricStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DeleteMetricStreamOutput), req.Error
}

func (c *Client) DescribeAlarmHistoryWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmHistoryInput, opts ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DescribeAlarmHistory",
		Input:   input,
		Output:  (*cloudwatch.DescribeAlarmHistoryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DescribeAlarmHistoryWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DescribeAlarmHistoryOutput), req.Error
}

func (c *Client) DescribeAlarmHistoryPagesWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmHistoryInput, cb func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DescribeAlarmHistory",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CloudWatchAPI.DescribeAlarmHistoryPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeAlarmsWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmsInput, opts ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DescribeAlarms",
		Input:   input,
		Output:  (*cloudwatch.DescribeAlarmsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DescribeAlarmsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DescribeAlarmsOutput), req.Error
}

func (c *Client) DescribeAlarmsPagesWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmsInput, cb func(*cloudwatch.DescribeAlarmsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DescribeAlarms",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CloudWatchAPI.DescribeAlarmsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeAlarmsForMetricWithContext(ctx context.Context, input *cloudwatch.DescribeAlarmsForMetricInput, opts ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DescribeAlarmsForMetric",
		Input:   input,
		Output:  (*cloudwatch.DescribeAlarmsForMetricOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DescribeAlarmsForMetricWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DescribeAlarmsForMetricOutput), req.Error
}

func (c *Client) DescribeAnomalyDetectorsWithContext(ctx context.Context, input *cloudwatch.DescribeAnomalyDetectorsInput, opts ...request.Option) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DescribeAnomalyDetectors",
		Input:   input,
		Output:  (*cloudwatch.DescribeAnomalyDetectorsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DescribeAnomalyDetectorsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DescribeAnomalyDetectorsOutput), req.Error
}

func (c *Client) DescribeAnomalyDetectorsPagesWithContext(ctx context.Context, input *cloudwatch.DescribeAnomalyDetectorsInput, cb func(*cloudwatch.DescribeAnomalyDetectorsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DescribeAnomalyDetectors",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CloudWatchAPI.DescribeAnomalyDetectorsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeInsightRulesWithContext(ctx context.Context, input *cloudwatch.DescribeInsightRulesInput, opts ...request.Option) (*cloudwatch.DescribeInsightRulesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DescribeInsightRules",
		Input:   input,
		Output:  (*cloudwatch.DescribeInsightRulesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DescribeInsightRulesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DescribeInsightRulesOutput), req.Error
}

func (c *Client) DescribeInsightRulesPagesWithContext(ctx context.Context, input *cloudwatch.DescribeInsightRulesInput, cb func(*cloudwatch.DescribeInsightRulesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DescribeInsightRules",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CloudWatchAPI.DescribeInsightRulesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DisableAlarmActionsWithContext(ctx context.Context, input *cloudwatch.DisableAlarmActionsInput, opts ...request.Option) (*cloudwatch.DisableAlarmActionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DisableAlarmActions",
		Input:   input,
		Output:  (*cloudwatch.DisableAlarmActionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DisableAlarmActionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DisableAlarmActionsOutput), req.Error
}

func (c *Client) DisableInsightRulesWithContext(ctx context.Context, input *cloudwatch.DisableInsightRulesInput, opts ...request.Option) (*cloudwatch.DisableInsightRulesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "DisableInsightRules",
		Input:   input,
		Output:  (*cloudwatch.DisableInsightRulesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.DisableInsightRulesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.DisableInsightRulesOutput), req.Error
}

func (c *Client) EnableAlarmActionsWithContext(ctx context.Context, input *cloudwatch.EnableAlarmActionsInput, opts ...request.Option) (*cloudwatch.EnableAlarmActionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "EnableAlarmActions",
		Input:   input,
		Output:  (*cloudwatch.EnableAlarmActionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.EnableAlarmActionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.EnableAlarmActionsOutput), req.Error
}

func (c *Client) EnableInsightRulesWithContext(ctx context.Context, input *cloudwatch.EnableInsightRulesInput, opts ...request.Option) (*cloudwatch.EnableInsightRulesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "EnableInsightRules",
		Input:   input,
		Output:  (*cloudwatch.EnableInsightRulesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.EnableInsightRulesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.EnableInsightRulesOutput), req.Error
}

func (c *Client) GetDashboardWithContext(ctx context.Context, input *cloudwatch.GetDashboardInput, opts ...request.Option) (*cloudwatch.GetDashboardOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "GetDashboard",
		Input:   input,
		Output:  (*cloudwatch.GetDashboardOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.GetDashboardWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.GetDashboardOutput), req.Error
}

func (c *Client) GetInsightRuleReportWithContext(ctx context.Context, input *cloudwatch.GetInsightRuleReportInput, opts ...request.Option) (*cloudwatch.GetInsightRuleReportOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "GetInsightRuleReport",
		Input:   input,
		Output:  (*cloudwatch.GetInsightRuleReportOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.GetInsightRuleReportWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.GetInsightRuleReportOutput), req.Error
}

func (c *Client) GetMetricDataWithContext(ctx context.Context, input *cloudwatch.GetMetricDataInput, opts ...request.Option) (*cloudwatch.GetMetricDataOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "GetMetricData",
		Input:   input,
		Output:  (*cloudwatch.GetMetricDataOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.GetMetricDataWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.GetMetricDataOutput), req.Error
}

func (c *Client) GetMetricDataPagesWithContext(ctx context.Context, input *cloudwatch.GetMetricDataInput, cb func(*cloudwatch.GetMetricDataOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "GetMetricData",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CloudWatchAPI.GetMetricDataPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetMetricStatisticsWithContext(ctx context.Context, input *cloudwatch.GetMetricStatisticsInput, opts ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "GetMetricStatistics",
		Input:   input,
		Output:  (*cloudwatch.GetMetricStatisticsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.GetMetricStatisticsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.GetMetricStatisticsOutput), req.Error
}

func (c *Client) GetMetricStreamWithContext(ctx context.Context, input *cloudwatch.GetMetricStreamInput, opts ...request.Option) (*cloudwatch.GetMetricStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "GetMetricStream",
		Input:   input,
		Output:  (*cloudwatch.GetMetricStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.GetMetricStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.GetMetricStreamOutput), req.Error
}

func (c *Client) GetMetricWidgetImageWithContext(ctx context.Context, input *cloudwatch.GetMetricWidgetImageInput, opts ...request.Option) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "GetMetricWidgetImage",
		Input:   input,
		Output:  (*cloudwatch.GetMetricWidgetImageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.GetMetricWidgetImageWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.GetMetricWidgetImageOutput), req.Error
}

func (c *Client) ListDashboardsWithContext(ctx context.Context, input *cloudwatch.ListDashboardsInput, opts ...request.Option) (*cloudwatch.ListDashboardsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "ListDashboards",
		Input:   input,
		Output:  (*cloudwatch.ListDashboardsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.ListDashboardsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.ListDashboardsOutput), req.Error
}

func (c *Client) ListDashboardsPagesWithContext(ctx context.Context, input *cloudwatch.ListDashboardsInput, cb func(*cloudwatch.ListDashboardsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "ListDashboards",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CloudWatchAPI.ListDashboardsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListManagedInsightRulesWithContext(ctx context.Context, input *cloudwatch.ListManagedInsightRulesInput, opts ...request.Option) (*cloudwatch.ListManagedInsightRulesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "ListManagedInsightRules",
		Input:   input,
		Output:  (*cloudwatch.ListManagedInsightRulesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.ListManagedInsightRulesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.ListManagedInsightRulesOutput), req.Error
}

func (c *Client) ListManagedInsightRulesPagesWithContext(ctx context.Context, input *cloudwatch.ListManagedInsightRulesInput, cb func(*cloudwatch.ListManagedInsightRulesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "ListManagedInsightRules",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CloudWatchAPI.ListManagedInsightRulesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListMetricStreamsWithContext(ctx context.Context, input *cloudwatch.ListMetricStreamsInput, opts ...request.Option) (*cloudwatch.ListMetricStreamsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "ListMetricStreams",
		Input:   input,
		Output:  (*cloudwatch.ListMetricStreamsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.ListMetricStreamsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.ListMetricStreamsOutput), req.Error
}

func (c *Client) ListMetricStreamsPagesWithContext(ctx context.Context, input *cloudwatch.ListMetricStreamsInput, cb func(*cloudwatch.ListMetricStreamsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "ListMetricStreams",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CloudWatchAPI.ListMetricStreamsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListMetricsWithContext(ctx context.Context, input *cloudwatch.ListMetricsInput, opts ...request.Option) (*cloudwatch.ListMetricsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "ListMetrics",
		Input:   input,
		Output:  (*cloudwatch.ListMetricsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.ListMetricsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.ListMetricsOutput), req.Error
}

func (c *Client) ListMetricsPagesWithContext(ctx context.Context, input *cloudwatch.ListMetricsInput, cb func(*cloudwatch.ListMetricsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "ListMetrics",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CloudWatchAPI.ListMetricsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *cloudwatch.ListTagsForResourceInput, opts ...request.Option) (*cloudwatch.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*cloudwatch.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.ListTagsForResourceOutput), req.Error
}

func (c *Client) PutAnomalyDetectorWithContext(ctx context.Context, input *cloudwatch.PutAnomalyDetectorInput, opts ...request.Option) (*cloudwatch.PutAnomalyDetectorOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "PutAnomalyDetector",
		Input:   input,
		Output:  (*cloudwatch.PutAnomalyDetectorOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.PutAnomalyDetectorWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.PutAnomalyDetectorOutput), req.Error
}

func (c *Client) PutCompositeAlarmWithContext(ctx context.Context, input *cloudwatch.PutCompositeAlarmInput, opts ...request.Option) (*cloudwatch.PutCompositeAlarmOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "PutCompositeAlarm",
		Input:   input,
		Output:  (*cloudwatch.PutCompositeAlarmOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.PutCompositeAlarmWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.PutCompositeAlarmOutput), req.Error
}

func (c *Client) PutDashboardWithContext(ctx context.Context, input *cloudwatch.PutDashboardInput, opts ...request.Option) (*cloudwatch.PutDashboardOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "PutDashboard",
		Input:   input,
		Output:  (*cloudwatch.PutDashboardOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.PutDashboardWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.PutDashboardOutput), req.Error
}

func (c *Client) PutInsightRuleWithContext(ctx context.Context, input *cloudwatch.PutInsightRuleInput, opts ...request.Option) (*cloudwatch.PutInsightRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "PutInsightRule",
		Input:   input,
		Output:  (*cloudwatch.PutInsightRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.PutInsightRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.PutInsightRuleOutput), req.Error
}

func (c *Client) PutManagedInsightRulesWithContext(ctx context.Context, input *cloudwatch.PutManagedInsightRulesInput, opts ...request.Option) (*cloudwatch.PutManagedInsightRulesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "PutManagedInsightRules",
		Input:   input,
		Output:  (*cloudwatch.PutManagedInsightRulesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.PutManagedInsightRulesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.PutManagedInsightRulesOutput), req.Error
}

func (c *Client) PutMetricAlarmWithContext(ctx context.Context, input *cloudwatch.PutMetricAlarmInput, opts ...request.Option) (*cloudwatch.PutMetricAlarmOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "PutMetricAlarm",
		Input:   input,
		Output:  (*cloudwatch.PutMetricAlarmOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.PutMetricAlarmWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.PutMetricAlarmOutput), req.Error
}

func (c *Client) PutMetricDataWithContext(ctx context.Context, input *cloudwatch.PutMetricDataInput, opts ...request.Option) (*cloudwatch.PutMetricDataOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "PutMetricData",
		Input:   input,
		Output:  (*cloudwatch.PutMetricDataOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.PutMetricDataWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.PutMetricDataOutput), req.Error
}

func (c *Client) PutMetricStreamWithContext(ctx context.Context, input *cloudwatch.PutMetricStreamInput, opts ...request.Option) (*cloudwatch.PutMetricStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "PutMetricStream",
		Input:   input,
		Output:  (*cloudwatch.PutMetricStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.PutMetricStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.PutMetricStreamOutput), req.Error
}

func (c *Client) SetAlarmStateWithContext(ctx context.Context, input *cloudwatch.SetAlarmStateInput, opts ...request.Option) (*cloudwatch.SetAlarmStateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "SetAlarmState",
		Input:   input,
		Output:  (*cloudwatch.SetAlarmStateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.SetAlarmStateWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.SetAlarmStateOutput), req.Error
}

func (c *Client) StartMetricStreamsWithContext(ctx context.Context, input *cloudwatch.StartMetricStreamsInput, opts ...request.Option) (*cloudwatch.StartMetricStreamsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "StartMetricStreams",
		Input:   input,
		Output:  (*cloudwatch.StartMetricStreamsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.StartMetricStreamsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.StartMetricStreamsOutput), req.Error
}

func (c *Client) StopMetricStreamsWithContext(ctx context.Context, input *cloudwatch.StopMetricStreamsInput, opts ...request.Option) (*cloudwatch.StopMetricStreamsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "StopMetricStreams",
		Input:   input,
		Output:  (*cloudwatch.StopMetricStreamsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.StopMetricStreamsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.StopMetricStreamsOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *cloudwatch.TagResourceInput, opts ...request.Option) (*cloudwatch.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "TagResource",
		Input:   input,
		Output:  (*cloudwatch.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *cloudwatch.UntagResourceInput, opts ...request.Option) (*cloudwatch.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloudwatch",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*cloudwatch.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CloudWatchAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloudwatch.UntagResourceOutput), req.Error
}
