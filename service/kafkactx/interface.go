// Code generated by internal/generate/main.go. DO NOT EDIT.

package kafkactx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/aws/aws-sdk-go/service/kafka/kafkaiface"
	"github.com/glassechidna/awsctx"
)

type Kafka interface {
	CreateClusterWithContext(ctx context.Context, input *kafka.CreateClusterInput, opts ...request.Option) (*kafka.CreateClusterOutput, error)
	CreateConfigurationWithContext(ctx context.Context, input *kafka.CreateConfigurationInput, opts ...request.Option) (*kafka.CreateConfigurationOutput, error)
	DeleteClusterWithContext(ctx context.Context, input *kafka.DeleteClusterInput, opts ...request.Option) (*kafka.DeleteClusterOutput, error)
	DescribeClusterWithContext(ctx context.Context, input *kafka.DescribeClusterInput, opts ...request.Option) (*kafka.DescribeClusterOutput, error)
	DescribeClusterOperationWithContext(ctx context.Context, input *kafka.DescribeClusterOperationInput, opts ...request.Option) (*kafka.DescribeClusterOperationOutput, error)
	DescribeConfigurationWithContext(ctx context.Context, input *kafka.DescribeConfigurationInput, opts ...request.Option) (*kafka.DescribeConfigurationOutput, error)
	DescribeConfigurationRevisionWithContext(ctx context.Context, input *kafka.DescribeConfigurationRevisionInput, opts ...request.Option) (*kafka.DescribeConfigurationRevisionOutput, error)
	GetBootstrapBrokersWithContext(ctx context.Context, input *kafka.GetBootstrapBrokersInput, opts ...request.Option) (*kafka.GetBootstrapBrokersOutput, error)
	GetCompatibleKafkaVersionsWithContext(ctx context.Context, input *kafka.GetCompatibleKafkaVersionsInput, opts ...request.Option) (*kafka.GetCompatibleKafkaVersionsOutput, error)
	ListClusterOperationsWithContext(ctx context.Context, input *kafka.ListClusterOperationsInput, opts ...request.Option) (*kafka.ListClusterOperationsOutput, error)
	ListClusterOperationsPagesWithContext(ctx context.Context, input *kafka.ListClusterOperationsInput, cb func(*kafka.ListClusterOperationsOutput, bool) bool, opts ...request.Option) error
	ListClustersWithContext(ctx context.Context, input *kafka.ListClustersInput, opts ...request.Option) (*kafka.ListClustersOutput, error)
	ListClustersPagesWithContext(ctx context.Context, input *kafka.ListClustersInput, cb func(*kafka.ListClustersOutput, bool) bool, opts ...request.Option) error
	ListConfigurationRevisionsWithContext(ctx context.Context, input *kafka.ListConfigurationRevisionsInput, opts ...request.Option) (*kafka.ListConfigurationRevisionsOutput, error)
	ListConfigurationRevisionsPagesWithContext(ctx context.Context, input *kafka.ListConfigurationRevisionsInput, cb func(*kafka.ListConfigurationRevisionsOutput, bool) bool, opts ...request.Option) error
	ListConfigurationsWithContext(ctx context.Context, input *kafka.ListConfigurationsInput, opts ...request.Option) (*kafka.ListConfigurationsOutput, error)
	ListConfigurationsPagesWithContext(ctx context.Context, input *kafka.ListConfigurationsInput, cb func(*kafka.ListConfigurationsOutput, bool) bool, opts ...request.Option) error
	ListKafkaVersionsWithContext(ctx context.Context, input *kafka.ListKafkaVersionsInput, opts ...request.Option) (*kafka.ListKafkaVersionsOutput, error)
	ListKafkaVersionsPagesWithContext(ctx context.Context, input *kafka.ListKafkaVersionsInput, cb func(*kafka.ListKafkaVersionsOutput, bool) bool, opts ...request.Option) error
	ListNodesWithContext(ctx context.Context, input *kafka.ListNodesInput, opts ...request.Option) (*kafka.ListNodesOutput, error)
	ListNodesPagesWithContext(ctx context.Context, input *kafka.ListNodesInput, cb func(*kafka.ListNodesOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *kafka.ListTagsForResourceInput, opts ...request.Option) (*kafka.ListTagsForResourceOutput, error)
	RebootBrokerWithContext(ctx context.Context, input *kafka.RebootBrokerInput, opts ...request.Option) (*kafka.RebootBrokerOutput, error)
	TagResourceWithContext(ctx context.Context, input *kafka.TagResourceInput, opts ...request.Option) (*kafka.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *kafka.UntagResourceInput, opts ...request.Option) (*kafka.UntagResourceOutput, error)
	UpdateBrokerCountWithContext(ctx context.Context, input *kafka.UpdateBrokerCountInput, opts ...request.Option) (*kafka.UpdateBrokerCountOutput, error)
	UpdateBrokerStorageWithContext(ctx context.Context, input *kafka.UpdateBrokerStorageInput, opts ...request.Option) (*kafka.UpdateBrokerStorageOutput, error)
	UpdateClusterConfigurationWithContext(ctx context.Context, input *kafka.UpdateClusterConfigurationInput, opts ...request.Option) (*kafka.UpdateClusterConfigurationOutput, error)
	UpdateClusterKafkaVersionWithContext(ctx context.Context, input *kafka.UpdateClusterKafkaVersionInput, opts ...request.Option) (*kafka.UpdateClusterKafkaVersionOutput, error)
	UpdateMonitoringWithContext(ctx context.Context, input *kafka.UpdateMonitoringInput, opts ...request.Option) (*kafka.UpdateMonitoringOutput, error)
}

type Client struct {
	kafkaiface.KafkaAPI
	Contexter awsctx.Contexter
}

func New(base kafkaiface.KafkaAPI, ctxer awsctx.Contexter) Kafka {
	return &Client{
		KafkaAPI: base,
		Contexter: ctxer,
	}
}

var _ Kafka = (*kafka.Kafka)(nil)
var _ Kafka = (*Client)(nil)

func (c *Client) CreateClusterWithContext(ctx context.Context, input *kafka.CreateClusterInput, opts ...request.Option) (*kafka.CreateClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "CreateCluster",
		Input:   input,
		Output:  (*kafka.CreateClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.CreateClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.CreateClusterOutput), req.Error
}

func (c *Client) CreateConfigurationWithContext(ctx context.Context, input *kafka.CreateConfigurationInput, opts ...request.Option) (*kafka.CreateConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "CreateConfiguration",
		Input:   input,
		Output:  (*kafka.CreateConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.CreateConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.CreateConfigurationOutput), req.Error
}

func (c *Client) DeleteClusterWithContext(ctx context.Context, input *kafka.DeleteClusterInput, opts ...request.Option) (*kafka.DeleteClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "DeleteCluster",
		Input:   input,
		Output:  (*kafka.DeleteClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.DeleteClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.DeleteClusterOutput), req.Error
}

func (c *Client) DescribeClusterWithContext(ctx context.Context, input *kafka.DescribeClusterInput, opts ...request.Option) (*kafka.DescribeClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "DescribeCluster",
		Input:   input,
		Output:  (*kafka.DescribeClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.DescribeClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.DescribeClusterOutput), req.Error
}

func (c *Client) DescribeClusterOperationWithContext(ctx context.Context, input *kafka.DescribeClusterOperationInput, opts ...request.Option) (*kafka.DescribeClusterOperationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "DescribeClusterOperation",
		Input:   input,
		Output:  (*kafka.DescribeClusterOperationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.DescribeClusterOperationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.DescribeClusterOperationOutput), req.Error
}

func (c *Client) DescribeConfigurationWithContext(ctx context.Context, input *kafka.DescribeConfigurationInput, opts ...request.Option) (*kafka.DescribeConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "DescribeConfiguration",
		Input:   input,
		Output:  (*kafka.DescribeConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.DescribeConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.DescribeConfigurationOutput), req.Error
}

func (c *Client) DescribeConfigurationRevisionWithContext(ctx context.Context, input *kafka.DescribeConfigurationRevisionInput, opts ...request.Option) (*kafka.DescribeConfigurationRevisionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "DescribeConfigurationRevision",
		Input:   input,
		Output:  (*kafka.DescribeConfigurationRevisionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.DescribeConfigurationRevisionWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.DescribeConfigurationRevisionOutput), req.Error
}

func (c *Client) GetBootstrapBrokersWithContext(ctx context.Context, input *kafka.GetBootstrapBrokersInput, opts ...request.Option) (*kafka.GetBootstrapBrokersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "GetBootstrapBrokers",
		Input:   input,
		Output:  (*kafka.GetBootstrapBrokersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.GetBootstrapBrokersWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.GetBootstrapBrokersOutput), req.Error
}

func (c *Client) GetCompatibleKafkaVersionsWithContext(ctx context.Context, input *kafka.GetCompatibleKafkaVersionsInput, opts ...request.Option) (*kafka.GetCompatibleKafkaVersionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "GetCompatibleKafkaVersions",
		Input:   input,
		Output:  (*kafka.GetCompatibleKafkaVersionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.GetCompatibleKafkaVersionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.GetCompatibleKafkaVersionsOutput), req.Error
}

func (c *Client) ListClusterOperationsWithContext(ctx context.Context, input *kafka.ListClusterOperationsInput, opts ...request.Option) (*kafka.ListClusterOperationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListClusterOperations",
		Input:   input,
		Output:  (*kafka.ListClusterOperationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.ListClusterOperationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.ListClusterOperationsOutput), req.Error
}

func (c *Client) ListClusterOperationsPagesWithContext(ctx context.Context, input *kafka.ListClusterOperationsInput, cb func(*kafka.ListClusterOperationsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListClusterOperations",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.KafkaAPI.ListClusterOperationsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListClustersWithContext(ctx context.Context, input *kafka.ListClustersInput, opts ...request.Option) (*kafka.ListClustersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListClusters",
		Input:   input,
		Output:  (*kafka.ListClustersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.ListClustersWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.ListClustersOutput), req.Error
}

func (c *Client) ListClustersPagesWithContext(ctx context.Context, input *kafka.ListClustersInput, cb func(*kafka.ListClustersOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListClusters",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.KafkaAPI.ListClustersPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListConfigurationRevisionsWithContext(ctx context.Context, input *kafka.ListConfigurationRevisionsInput, opts ...request.Option) (*kafka.ListConfigurationRevisionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListConfigurationRevisions",
		Input:   input,
		Output:  (*kafka.ListConfigurationRevisionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.ListConfigurationRevisionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.ListConfigurationRevisionsOutput), req.Error
}

func (c *Client) ListConfigurationRevisionsPagesWithContext(ctx context.Context, input *kafka.ListConfigurationRevisionsInput, cb func(*kafka.ListConfigurationRevisionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListConfigurationRevisions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.KafkaAPI.ListConfigurationRevisionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListConfigurationsWithContext(ctx context.Context, input *kafka.ListConfigurationsInput, opts ...request.Option) (*kafka.ListConfigurationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListConfigurations",
		Input:   input,
		Output:  (*kafka.ListConfigurationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.ListConfigurationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.ListConfigurationsOutput), req.Error
}

func (c *Client) ListConfigurationsPagesWithContext(ctx context.Context, input *kafka.ListConfigurationsInput, cb func(*kafka.ListConfigurationsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListConfigurations",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.KafkaAPI.ListConfigurationsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListKafkaVersionsWithContext(ctx context.Context, input *kafka.ListKafkaVersionsInput, opts ...request.Option) (*kafka.ListKafkaVersionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListKafkaVersions",
		Input:   input,
		Output:  (*kafka.ListKafkaVersionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.ListKafkaVersionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.ListKafkaVersionsOutput), req.Error
}

func (c *Client) ListKafkaVersionsPagesWithContext(ctx context.Context, input *kafka.ListKafkaVersionsInput, cb func(*kafka.ListKafkaVersionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListKafkaVersions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.KafkaAPI.ListKafkaVersionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListNodesWithContext(ctx context.Context, input *kafka.ListNodesInput, opts ...request.Option) (*kafka.ListNodesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListNodes",
		Input:   input,
		Output:  (*kafka.ListNodesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.ListNodesWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.ListNodesOutput), req.Error
}

func (c *Client) ListNodesPagesWithContext(ctx context.Context, input *kafka.ListNodesInput, cb func(*kafka.ListNodesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListNodes",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.KafkaAPI.ListNodesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *kafka.ListTagsForResourceInput, opts ...request.Option) (*kafka.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*kafka.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.ListTagsForResourceOutput), req.Error
}

func (c *Client) RebootBrokerWithContext(ctx context.Context, input *kafka.RebootBrokerInput, opts ...request.Option) (*kafka.RebootBrokerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "RebootBroker",
		Input:   input,
		Output:  (*kafka.RebootBrokerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.RebootBrokerWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.RebootBrokerOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *kafka.TagResourceInput, opts ...request.Option) (*kafka.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "TagResource",
		Input:   input,
		Output:  (*kafka.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *kafka.UntagResourceInput, opts ...request.Option) (*kafka.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*kafka.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.UntagResourceOutput), req.Error
}

func (c *Client) UpdateBrokerCountWithContext(ctx context.Context, input *kafka.UpdateBrokerCountInput, opts ...request.Option) (*kafka.UpdateBrokerCountOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "UpdateBrokerCount",
		Input:   input,
		Output:  (*kafka.UpdateBrokerCountOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.UpdateBrokerCountWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.UpdateBrokerCountOutput), req.Error
}

func (c *Client) UpdateBrokerStorageWithContext(ctx context.Context, input *kafka.UpdateBrokerStorageInput, opts ...request.Option) (*kafka.UpdateBrokerStorageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "UpdateBrokerStorage",
		Input:   input,
		Output:  (*kafka.UpdateBrokerStorageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.UpdateBrokerStorageWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.UpdateBrokerStorageOutput), req.Error
}

func (c *Client) UpdateClusterConfigurationWithContext(ctx context.Context, input *kafka.UpdateClusterConfigurationInput, opts ...request.Option) (*kafka.UpdateClusterConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "UpdateClusterConfiguration",
		Input:   input,
		Output:  (*kafka.UpdateClusterConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.UpdateClusterConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.UpdateClusterConfigurationOutput), req.Error
}

func (c *Client) UpdateClusterKafkaVersionWithContext(ctx context.Context, input *kafka.UpdateClusterKafkaVersionInput, opts ...request.Option) (*kafka.UpdateClusterKafkaVersionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "UpdateClusterKafkaVersion",
		Input:   input,
		Output:  (*kafka.UpdateClusterKafkaVersionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.UpdateClusterKafkaVersionWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.UpdateClusterKafkaVersionOutput), req.Error
}

func (c *Client) UpdateMonitoringWithContext(ctx context.Context, input *kafka.UpdateMonitoringInput, opts ...request.Option) (*kafka.UpdateMonitoringOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kafka",
		Action:  "UpdateMonitoring",
		Input:   input,
		Output:  (*kafka.UpdateMonitoringOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KafkaAPI.UpdateMonitoringWithContext(ctx, input, opts...)
	})

	return req.Output.(*kafka.UpdateMonitoringOutput), req.Error
}
