// Code generated by internal/generate/main.go. DO NOT EDIT.

package dynamodbctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/glassechidna/awsctx"
)

type DynamoDB interface {
	BatchGetItemWithContext(ctx context.Context, input *dynamodb.BatchGetItemInput, opts ...request.Option) (*dynamodb.BatchGetItemOutput, error)
	BatchWriteItemWithContext(ctx context.Context, input *dynamodb.BatchWriteItemInput, opts ...request.Option) (*dynamodb.BatchWriteItemOutput, error)
	CreateBackupWithContext(ctx context.Context, input *dynamodb.CreateBackupInput, opts ...request.Option) (*dynamodb.CreateBackupOutput, error)
	CreateGlobalTableWithContext(ctx context.Context, input *dynamodb.CreateGlobalTableInput, opts ...request.Option) (*dynamodb.CreateGlobalTableOutput, error)
	CreateTableWithContext(ctx context.Context, input *dynamodb.CreateTableInput, opts ...request.Option) (*dynamodb.CreateTableOutput, error)
	DeleteBackupWithContext(ctx context.Context, input *dynamodb.DeleteBackupInput, opts ...request.Option) (*dynamodb.DeleteBackupOutput, error)
	DeleteItemWithContext(ctx context.Context, input *dynamodb.DeleteItemInput, opts ...request.Option) (*dynamodb.DeleteItemOutput, error)
	DeleteTableWithContext(ctx context.Context, input *dynamodb.DeleteTableInput, opts ...request.Option) (*dynamodb.DeleteTableOutput, error)
	DescribeBackupWithContext(ctx context.Context, input *dynamodb.DescribeBackupInput, opts ...request.Option) (*dynamodb.DescribeBackupOutput, error)
	DescribeContinuousBackupsWithContext(ctx context.Context, input *dynamodb.DescribeContinuousBackupsInput, opts ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error)
	DescribeEndpointsWithContext(ctx context.Context, input *dynamodb.DescribeEndpointsInput, opts ...request.Option) (*dynamodb.DescribeEndpointsOutput, error)
	DescribeGlobalTableWithContext(ctx context.Context, input *dynamodb.DescribeGlobalTableInput, opts ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error)
	DescribeGlobalTableSettingsWithContext(ctx context.Context, input *dynamodb.DescribeGlobalTableSettingsInput, opts ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	DescribeLimitsWithContext(ctx context.Context, input *dynamodb.DescribeLimitsInput, opts ...request.Option) (*dynamodb.DescribeLimitsOutput, error)
	DescribeTableWithContext(ctx context.Context, input *dynamodb.DescribeTableInput, opts ...request.Option) (*dynamodb.DescribeTableOutput, error)
	DescribeTimeToLiveWithContext(ctx context.Context, input *dynamodb.DescribeTimeToLiveInput, opts ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error)
	GetItemWithContext(ctx context.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error)
	ListBackupsWithContext(ctx context.Context, input *dynamodb.ListBackupsInput, opts ...request.Option) (*dynamodb.ListBackupsOutput, error)
	ListGlobalTablesWithContext(ctx context.Context, input *dynamodb.ListGlobalTablesInput, opts ...request.Option) (*dynamodb.ListGlobalTablesOutput, error)
	ListTablesWithContext(ctx context.Context, input *dynamodb.ListTablesInput, opts ...request.Option) (*dynamodb.ListTablesOutput, error)
	ListTagsOfResourceWithContext(ctx context.Context, input *dynamodb.ListTagsOfResourceInput, opts ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error)
	PutItemWithContext(ctx context.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error)
	QueryWithContext(ctx context.Context, input *dynamodb.QueryInput, opts ...request.Option) (*dynamodb.QueryOutput, error)
	RestoreTableFromBackupWithContext(ctx context.Context, input *dynamodb.RestoreTableFromBackupInput, opts ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error)
	RestoreTableToPointInTimeWithContext(ctx context.Context, input *dynamodb.RestoreTableToPointInTimeInput, opts ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error)
	ScanWithContext(ctx context.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error)
	TagResourceWithContext(ctx context.Context, input *dynamodb.TagResourceInput, opts ...request.Option) (*dynamodb.TagResourceOutput, error)
	TransactGetItemsWithContext(ctx context.Context, input *dynamodb.TransactGetItemsInput, opts ...request.Option) (*dynamodb.TransactGetItemsOutput, error)
	TransactWriteItemsWithContext(ctx context.Context, input *dynamodb.TransactWriteItemsInput, opts ...request.Option) (*dynamodb.TransactWriteItemsOutput, error)
	UntagResourceWithContext(ctx context.Context, input *dynamodb.UntagResourceInput, opts ...request.Option) (*dynamodb.UntagResourceOutput, error)
	UpdateContinuousBackupsWithContext(ctx context.Context, input *dynamodb.UpdateContinuousBackupsInput, opts ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error)
	UpdateGlobalTableWithContext(ctx context.Context, input *dynamodb.UpdateGlobalTableInput, opts ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error)
	UpdateGlobalTableSettingsWithContext(ctx context.Context, input *dynamodb.UpdateGlobalTableSettingsInput, opts ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
	UpdateItemWithContext(ctx context.Context, input *dynamodb.UpdateItemInput, opts ...request.Option) (*dynamodb.UpdateItemOutput, error)
	UpdateTableWithContext(ctx context.Context, input *dynamodb.UpdateTableInput, opts ...request.Option) (*dynamodb.UpdateTableOutput, error)
	UpdateTimeToLiveWithContext(ctx context.Context, input *dynamodb.UpdateTimeToLiveInput, opts ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error)
}

type Client struct {
	dynamodbiface.DynamoDBAPI
	Contexter awsctx.Contexter
}

func New(base dynamodbiface.DynamoDBAPI, ctxer awsctx.Contexter) DynamoDB {
	return &Client{
		DynamoDBAPI: base,
		Contexter: ctxer,
	}
}

var _ DynamoDB = (*dynamodb.DynamoDB)(nil)
var _ DynamoDB = (*Client)(nil)

func (c *Client) BatchGetItemWithContext(ctx context.Context, input *dynamodb.BatchGetItemInput, opts ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "BatchGetItemWithContext",
		Input:   input,
		Output:  (*dynamodb.BatchGetItemOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.BatchGetItemWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.BatchGetItemOutput), req.Error
}

func (c *Client) BatchWriteItemWithContext(ctx context.Context, input *dynamodb.BatchWriteItemInput, opts ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "BatchWriteItemWithContext",
		Input:   input,
		Output:  (*dynamodb.BatchWriteItemOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.BatchWriteItemWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.BatchWriteItemOutput), req.Error
}

func (c *Client) CreateBackupWithContext(ctx context.Context, input *dynamodb.CreateBackupInput, opts ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "CreateBackupWithContext",
		Input:   input,
		Output:  (*dynamodb.CreateBackupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.CreateBackupWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.CreateBackupOutput), req.Error
}

func (c *Client) CreateGlobalTableWithContext(ctx context.Context, input *dynamodb.CreateGlobalTableInput, opts ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "CreateGlobalTableWithContext",
		Input:   input,
		Output:  (*dynamodb.CreateGlobalTableOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.CreateGlobalTableWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.CreateGlobalTableOutput), req.Error
}

func (c *Client) CreateTableWithContext(ctx context.Context, input *dynamodb.CreateTableInput, opts ...request.Option) (*dynamodb.CreateTableOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "CreateTableWithContext",
		Input:   input,
		Output:  (*dynamodb.CreateTableOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.CreateTableWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.CreateTableOutput), req.Error
}

func (c *Client) DeleteBackupWithContext(ctx context.Context, input *dynamodb.DeleteBackupInput, opts ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DeleteBackupWithContext",
		Input:   input,
		Output:  (*dynamodb.DeleteBackupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DeleteBackupWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DeleteBackupOutput), req.Error
}

func (c *Client) DeleteItemWithContext(ctx context.Context, input *dynamodb.DeleteItemInput, opts ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DeleteItemWithContext",
		Input:   input,
		Output:  (*dynamodb.DeleteItemOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DeleteItemWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DeleteItemOutput), req.Error
}

func (c *Client) DeleteTableWithContext(ctx context.Context, input *dynamodb.DeleteTableInput, opts ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DeleteTableWithContext",
		Input:   input,
		Output:  (*dynamodb.DeleteTableOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DeleteTableWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DeleteTableOutput), req.Error
}

func (c *Client) DescribeBackupWithContext(ctx context.Context, input *dynamodb.DescribeBackupInput, opts ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DescribeBackupWithContext",
		Input:   input,
		Output:  (*dynamodb.DescribeBackupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DescribeBackupWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DescribeBackupOutput), req.Error
}

func (c *Client) DescribeContinuousBackupsWithContext(ctx context.Context, input *dynamodb.DescribeContinuousBackupsInput, opts ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DescribeContinuousBackupsWithContext",
		Input:   input,
		Output:  (*dynamodb.DescribeContinuousBackupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DescribeContinuousBackupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DescribeContinuousBackupsOutput), req.Error
}

func (c *Client) DescribeEndpointsWithContext(ctx context.Context, input *dynamodb.DescribeEndpointsInput, opts ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DescribeEndpointsWithContext",
		Input:   input,
		Output:  (*dynamodb.DescribeEndpointsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DescribeEndpointsWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DescribeEndpointsOutput), req.Error
}

func (c *Client) DescribeGlobalTableWithContext(ctx context.Context, input *dynamodb.DescribeGlobalTableInput, opts ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DescribeGlobalTableWithContext",
		Input:   input,
		Output:  (*dynamodb.DescribeGlobalTableOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DescribeGlobalTableWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DescribeGlobalTableOutput), req.Error
}

func (c *Client) DescribeGlobalTableSettingsWithContext(ctx context.Context, input *dynamodb.DescribeGlobalTableSettingsInput, opts ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DescribeGlobalTableSettingsWithContext",
		Input:   input,
		Output:  (*dynamodb.DescribeGlobalTableSettingsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DescribeGlobalTableSettingsWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DescribeGlobalTableSettingsOutput), req.Error
}

func (c *Client) DescribeLimitsWithContext(ctx context.Context, input *dynamodb.DescribeLimitsInput, opts ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DescribeLimitsWithContext",
		Input:   input,
		Output:  (*dynamodb.DescribeLimitsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DescribeLimitsWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DescribeLimitsOutput), req.Error
}

func (c *Client) DescribeTableWithContext(ctx context.Context, input *dynamodb.DescribeTableInput, opts ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DescribeTableWithContext",
		Input:   input,
		Output:  (*dynamodb.DescribeTableOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DescribeTableWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DescribeTableOutput), req.Error
}

func (c *Client) DescribeTimeToLiveWithContext(ctx context.Context, input *dynamodb.DescribeTimeToLiveInput, opts ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "DescribeTimeToLiveWithContext",
		Input:   input,
		Output:  (*dynamodb.DescribeTimeToLiveOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.DescribeTimeToLiveWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.DescribeTimeToLiveOutput), req.Error
}

func (c *Client) GetItemWithContext(ctx context.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "GetItemWithContext",
		Input:   input,
		Output:  (*dynamodb.GetItemOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.GetItemWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.GetItemOutput), req.Error
}

func (c *Client) ListBackupsWithContext(ctx context.Context, input *dynamodb.ListBackupsInput, opts ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "ListBackupsWithContext",
		Input:   input,
		Output:  (*dynamodb.ListBackupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.ListBackupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.ListBackupsOutput), req.Error
}

func (c *Client) ListGlobalTablesWithContext(ctx context.Context, input *dynamodb.ListGlobalTablesInput, opts ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "ListGlobalTablesWithContext",
		Input:   input,
		Output:  (*dynamodb.ListGlobalTablesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.ListGlobalTablesWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.ListGlobalTablesOutput), req.Error
}

func (c *Client) ListTablesWithContext(ctx context.Context, input *dynamodb.ListTablesInput, opts ...request.Option) (*dynamodb.ListTablesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "ListTablesWithContext",
		Input:   input,
		Output:  (*dynamodb.ListTablesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.ListTablesWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.ListTablesOutput), req.Error
}

func (c *Client) ListTagsOfResourceWithContext(ctx context.Context, input *dynamodb.ListTagsOfResourceInput, opts ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "ListTagsOfResourceWithContext",
		Input:   input,
		Output:  (*dynamodb.ListTagsOfResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.ListTagsOfResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.ListTagsOfResourceOutput), req.Error
}

func (c *Client) PutItemWithContext(ctx context.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "PutItemWithContext",
		Input:   input,
		Output:  (*dynamodb.PutItemOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.PutItemWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.PutItemOutput), req.Error
}

func (c *Client) QueryWithContext(ctx context.Context, input *dynamodb.QueryInput, opts ...request.Option) (*dynamodb.QueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "QueryWithContext",
		Input:   input,
		Output:  (*dynamodb.QueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.QueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.QueryOutput), req.Error
}

func (c *Client) RestoreTableFromBackupWithContext(ctx context.Context, input *dynamodb.RestoreTableFromBackupInput, opts ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "RestoreTableFromBackupWithContext",
		Input:   input,
		Output:  (*dynamodb.RestoreTableFromBackupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.RestoreTableFromBackupWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.RestoreTableFromBackupOutput), req.Error
}

func (c *Client) RestoreTableToPointInTimeWithContext(ctx context.Context, input *dynamodb.RestoreTableToPointInTimeInput, opts ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "RestoreTableToPointInTimeWithContext",
		Input:   input,
		Output:  (*dynamodb.RestoreTableToPointInTimeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.RestoreTableToPointInTimeWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.RestoreTableToPointInTimeOutput), req.Error
}

func (c *Client) ScanWithContext(ctx context.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "ScanWithContext",
		Input:   input,
		Output:  (*dynamodb.ScanOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.ScanWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.ScanOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *dynamodb.TagResourceInput, opts ...request.Option) (*dynamodb.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "TagResourceWithContext",
		Input:   input,
		Output:  (*dynamodb.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.TagResourceOutput), req.Error
}

func (c *Client) TransactGetItemsWithContext(ctx context.Context, input *dynamodb.TransactGetItemsInput, opts ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "TransactGetItemsWithContext",
		Input:   input,
		Output:  (*dynamodb.TransactGetItemsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.TransactGetItemsWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.TransactGetItemsOutput), req.Error
}

func (c *Client) TransactWriteItemsWithContext(ctx context.Context, input *dynamodb.TransactWriteItemsInput, opts ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "TransactWriteItemsWithContext",
		Input:   input,
		Output:  (*dynamodb.TransactWriteItemsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.TransactWriteItemsWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.TransactWriteItemsOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *dynamodb.UntagResourceInput, opts ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "UntagResourceWithContext",
		Input:   input,
		Output:  (*dynamodb.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.UntagResourceOutput), req.Error
}

func (c *Client) UpdateContinuousBackupsWithContext(ctx context.Context, input *dynamodb.UpdateContinuousBackupsInput, opts ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "UpdateContinuousBackupsWithContext",
		Input:   input,
		Output:  (*dynamodb.UpdateContinuousBackupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.UpdateContinuousBackupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.UpdateContinuousBackupsOutput), req.Error
}

func (c *Client) UpdateGlobalTableWithContext(ctx context.Context, input *dynamodb.UpdateGlobalTableInput, opts ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "UpdateGlobalTableWithContext",
		Input:   input,
		Output:  (*dynamodb.UpdateGlobalTableOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.UpdateGlobalTableWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.UpdateGlobalTableOutput), req.Error
}

func (c *Client) UpdateGlobalTableSettingsWithContext(ctx context.Context, input *dynamodb.UpdateGlobalTableSettingsInput, opts ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "UpdateGlobalTableSettingsWithContext",
		Input:   input,
		Output:  (*dynamodb.UpdateGlobalTableSettingsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.UpdateGlobalTableSettingsWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.UpdateGlobalTableSettingsOutput), req.Error
}

func (c *Client) UpdateItemWithContext(ctx context.Context, input *dynamodb.UpdateItemInput, opts ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "UpdateItemWithContext",
		Input:   input,
		Output:  (*dynamodb.UpdateItemOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.UpdateItemWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.UpdateItemOutput), req.Error
}

func (c *Client) UpdateTableWithContext(ctx context.Context, input *dynamodb.UpdateTableInput, opts ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "UpdateTableWithContext",
		Input:   input,
		Output:  (*dynamodb.UpdateTableOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.UpdateTableWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.UpdateTableOutput), req.Error
}

func (c *Client) UpdateTimeToLiveWithContext(ctx context.Context, input *dynamodb.UpdateTimeToLiveInput, opts ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dynamodb",
		Action:  "UpdateTimeToLiveWithContext",
		Input:   input,
		Output:  (*dynamodb.UpdateTimeToLiveOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DynamoDBAPI.UpdateTimeToLiveWithContext(ctx, input, opts...)
	})

	return req.Output.(*dynamodb.UpdateTimeToLiveOutput), req.Error
}
