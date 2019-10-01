// Code generated by internal/generate/main.go. DO NOT EDIT.

package machinelearningctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"github.com/aws/aws-sdk-go/service/machinelearning/machinelearningiface"
	"github.com/glassechidna/awsctx"
)

type MachineLearning interface {
	AddTagsWithContext(ctx context.Context, input *machinelearning.AddTagsInput, opts ...request.Option) (*machinelearning.AddTagsOutput, error)
	CreateBatchPredictionWithContext(ctx context.Context, input *machinelearning.CreateBatchPredictionInput, opts ...request.Option) (*machinelearning.CreateBatchPredictionOutput, error)
	CreateDataSourceFromRDSWithContext(ctx context.Context, input *machinelearning.CreateDataSourceFromRDSInput, opts ...request.Option) (*machinelearning.CreateDataSourceFromRDSOutput, error)
	CreateDataSourceFromRedshiftWithContext(ctx context.Context, input *machinelearning.CreateDataSourceFromRedshiftInput, opts ...request.Option) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)
	CreateDataSourceFromS3WithContext(ctx context.Context, input *machinelearning.CreateDataSourceFromS3Input, opts ...request.Option) (*machinelearning.CreateDataSourceFromS3Output, error)
	CreateEvaluationWithContext(ctx context.Context, input *machinelearning.CreateEvaluationInput, opts ...request.Option) (*machinelearning.CreateEvaluationOutput, error)
	CreateMLModelWithContext(ctx context.Context, input *machinelearning.CreateMLModelInput, opts ...request.Option) (*machinelearning.CreateMLModelOutput, error)
	CreateRealtimeEndpointWithContext(ctx context.Context, input *machinelearning.CreateRealtimeEndpointInput, opts ...request.Option) (*machinelearning.CreateRealtimeEndpointOutput, error)
	DeleteBatchPredictionWithContext(ctx context.Context, input *machinelearning.DeleteBatchPredictionInput, opts ...request.Option) (*machinelearning.DeleteBatchPredictionOutput, error)
	DeleteDataSourceWithContext(ctx context.Context, input *machinelearning.DeleteDataSourceInput, opts ...request.Option) (*machinelearning.DeleteDataSourceOutput, error)
	DeleteEvaluationWithContext(ctx context.Context, input *machinelearning.DeleteEvaluationInput, opts ...request.Option) (*machinelearning.DeleteEvaluationOutput, error)
	DeleteMLModelWithContext(ctx context.Context, input *machinelearning.DeleteMLModelInput, opts ...request.Option) (*machinelearning.DeleteMLModelOutput, error)
	DeleteRealtimeEndpointWithContext(ctx context.Context, input *machinelearning.DeleteRealtimeEndpointInput, opts ...request.Option) (*machinelearning.DeleteRealtimeEndpointOutput, error)
	DeleteTagsWithContext(ctx context.Context, input *machinelearning.DeleteTagsInput, opts ...request.Option) (*machinelearning.DeleteTagsOutput, error)
	DescribeBatchPredictionsWithContext(ctx context.Context, input *machinelearning.DescribeBatchPredictionsInput, opts ...request.Option) (*machinelearning.DescribeBatchPredictionsOutput, error)
	DescribeDataSourcesWithContext(ctx context.Context, input *machinelearning.DescribeDataSourcesInput, opts ...request.Option) (*machinelearning.DescribeDataSourcesOutput, error)
	DescribeEvaluationsWithContext(ctx context.Context, input *machinelearning.DescribeEvaluationsInput, opts ...request.Option) (*machinelearning.DescribeEvaluationsOutput, error)
	DescribeMLModelsWithContext(ctx context.Context, input *machinelearning.DescribeMLModelsInput, opts ...request.Option) (*machinelearning.DescribeMLModelsOutput, error)
	DescribeTagsWithContext(ctx context.Context, input *machinelearning.DescribeTagsInput, opts ...request.Option) (*machinelearning.DescribeTagsOutput, error)
	GetBatchPredictionWithContext(ctx context.Context, input *machinelearning.GetBatchPredictionInput, opts ...request.Option) (*machinelearning.GetBatchPredictionOutput, error)
	GetDataSourceWithContext(ctx context.Context, input *machinelearning.GetDataSourceInput, opts ...request.Option) (*machinelearning.GetDataSourceOutput, error)
	GetEvaluationWithContext(ctx context.Context, input *machinelearning.GetEvaluationInput, opts ...request.Option) (*machinelearning.GetEvaluationOutput, error)
	GetMLModelWithContext(ctx context.Context, input *machinelearning.GetMLModelInput, opts ...request.Option) (*machinelearning.GetMLModelOutput, error)
	PredictWithContext(ctx context.Context, input *machinelearning.PredictInput, opts ...request.Option) (*machinelearning.PredictOutput, error)
	UpdateBatchPredictionWithContext(ctx context.Context, input *machinelearning.UpdateBatchPredictionInput, opts ...request.Option) (*machinelearning.UpdateBatchPredictionOutput, error)
	UpdateDataSourceWithContext(ctx context.Context, input *machinelearning.UpdateDataSourceInput, opts ...request.Option) (*machinelearning.UpdateDataSourceOutput, error)
	UpdateEvaluationWithContext(ctx context.Context, input *machinelearning.UpdateEvaluationInput, opts ...request.Option) (*machinelearning.UpdateEvaluationOutput, error)
	UpdateMLModelWithContext(ctx context.Context, input *machinelearning.UpdateMLModelInput, opts ...request.Option) (*machinelearning.UpdateMLModelOutput, error)
}

type Client struct {
	machinelearningiface.MachineLearningAPI
	Contexter awsctx.Contexter
}

func New(base machinelearningiface.MachineLearningAPI, ctxer awsctx.Contexter) MachineLearning {
	return &Client{
		MachineLearningAPI: base,
		Contexter: ctxer,
	}
}

var _ MachineLearning = (*machinelearning.MachineLearning)(nil)
var _ MachineLearning = (*Client)(nil)

func (c *Client) AddTagsWithContext(ctx context.Context, input *machinelearning.AddTagsInput, opts ...request.Option) (*machinelearning.AddTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "AddTagsWithContext",
		Input:   input,
		Output:  (*machinelearning.AddTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.AddTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.AddTagsOutput), req.Error
}

func (c *Client) CreateBatchPredictionWithContext(ctx context.Context, input *machinelearning.CreateBatchPredictionInput, opts ...request.Option) (*machinelearning.CreateBatchPredictionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "CreateBatchPredictionWithContext",
		Input:   input,
		Output:  (*machinelearning.CreateBatchPredictionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.CreateBatchPredictionWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.CreateBatchPredictionOutput), req.Error
}

func (c *Client) CreateDataSourceFromRDSWithContext(ctx context.Context, input *machinelearning.CreateDataSourceFromRDSInput, opts ...request.Option) (*machinelearning.CreateDataSourceFromRDSOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "CreateDataSourceFromRDSWithContext",
		Input:   input,
		Output:  (*machinelearning.CreateDataSourceFromRDSOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.CreateDataSourceFromRDSWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.CreateDataSourceFromRDSOutput), req.Error
}

func (c *Client) CreateDataSourceFromRedshiftWithContext(ctx context.Context, input *machinelearning.CreateDataSourceFromRedshiftInput, opts ...request.Option) (*machinelearning.CreateDataSourceFromRedshiftOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "CreateDataSourceFromRedshiftWithContext",
		Input:   input,
		Output:  (*machinelearning.CreateDataSourceFromRedshiftOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.CreateDataSourceFromRedshiftWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.CreateDataSourceFromRedshiftOutput), req.Error
}

func (c *Client) CreateDataSourceFromS3WithContext(ctx context.Context, input *machinelearning.CreateDataSourceFromS3Input, opts ...request.Option) (*machinelearning.CreateDataSourceFromS3Output, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "CreateDataSourceFromS3WithContext",
		Input:   input,
		Output:  (*machinelearning.CreateDataSourceFromS3Output)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.CreateDataSourceFromS3WithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.CreateDataSourceFromS3Output), req.Error
}

func (c *Client) CreateEvaluationWithContext(ctx context.Context, input *machinelearning.CreateEvaluationInput, opts ...request.Option) (*machinelearning.CreateEvaluationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "CreateEvaluationWithContext",
		Input:   input,
		Output:  (*machinelearning.CreateEvaluationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.CreateEvaluationWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.CreateEvaluationOutput), req.Error
}

func (c *Client) CreateMLModelWithContext(ctx context.Context, input *machinelearning.CreateMLModelInput, opts ...request.Option) (*machinelearning.CreateMLModelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "CreateMLModelWithContext",
		Input:   input,
		Output:  (*machinelearning.CreateMLModelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.CreateMLModelWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.CreateMLModelOutput), req.Error
}

func (c *Client) CreateRealtimeEndpointWithContext(ctx context.Context, input *machinelearning.CreateRealtimeEndpointInput, opts ...request.Option) (*machinelearning.CreateRealtimeEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "CreateRealtimeEndpointWithContext",
		Input:   input,
		Output:  (*machinelearning.CreateRealtimeEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.CreateRealtimeEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.CreateRealtimeEndpointOutput), req.Error
}

func (c *Client) DeleteBatchPredictionWithContext(ctx context.Context, input *machinelearning.DeleteBatchPredictionInput, opts ...request.Option) (*machinelearning.DeleteBatchPredictionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DeleteBatchPredictionWithContext",
		Input:   input,
		Output:  (*machinelearning.DeleteBatchPredictionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DeleteBatchPredictionWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DeleteBatchPredictionOutput), req.Error
}

func (c *Client) DeleteDataSourceWithContext(ctx context.Context, input *machinelearning.DeleteDataSourceInput, opts ...request.Option) (*machinelearning.DeleteDataSourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DeleteDataSourceWithContext",
		Input:   input,
		Output:  (*machinelearning.DeleteDataSourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DeleteDataSourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DeleteDataSourceOutput), req.Error
}

func (c *Client) DeleteEvaluationWithContext(ctx context.Context, input *machinelearning.DeleteEvaluationInput, opts ...request.Option) (*machinelearning.DeleteEvaluationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DeleteEvaluationWithContext",
		Input:   input,
		Output:  (*machinelearning.DeleteEvaluationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DeleteEvaluationWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DeleteEvaluationOutput), req.Error
}

func (c *Client) DeleteMLModelWithContext(ctx context.Context, input *machinelearning.DeleteMLModelInput, opts ...request.Option) (*machinelearning.DeleteMLModelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DeleteMLModelWithContext",
		Input:   input,
		Output:  (*machinelearning.DeleteMLModelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DeleteMLModelWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DeleteMLModelOutput), req.Error
}

func (c *Client) DeleteRealtimeEndpointWithContext(ctx context.Context, input *machinelearning.DeleteRealtimeEndpointInput, opts ...request.Option) (*machinelearning.DeleteRealtimeEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DeleteRealtimeEndpointWithContext",
		Input:   input,
		Output:  (*machinelearning.DeleteRealtimeEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DeleteRealtimeEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DeleteRealtimeEndpointOutput), req.Error
}

func (c *Client) DeleteTagsWithContext(ctx context.Context, input *machinelearning.DeleteTagsInput, opts ...request.Option) (*machinelearning.DeleteTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DeleteTagsWithContext",
		Input:   input,
		Output:  (*machinelearning.DeleteTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DeleteTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DeleteTagsOutput), req.Error
}

func (c *Client) DescribeBatchPredictionsWithContext(ctx context.Context, input *machinelearning.DescribeBatchPredictionsInput, opts ...request.Option) (*machinelearning.DescribeBatchPredictionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DescribeBatchPredictionsWithContext",
		Input:   input,
		Output:  (*machinelearning.DescribeBatchPredictionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DescribeBatchPredictionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DescribeBatchPredictionsOutput), req.Error
}

func (c *Client) DescribeDataSourcesWithContext(ctx context.Context, input *machinelearning.DescribeDataSourcesInput, opts ...request.Option) (*machinelearning.DescribeDataSourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DescribeDataSourcesWithContext",
		Input:   input,
		Output:  (*machinelearning.DescribeDataSourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DescribeDataSourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DescribeDataSourcesOutput), req.Error
}

func (c *Client) DescribeEvaluationsWithContext(ctx context.Context, input *machinelearning.DescribeEvaluationsInput, opts ...request.Option) (*machinelearning.DescribeEvaluationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DescribeEvaluationsWithContext",
		Input:   input,
		Output:  (*machinelearning.DescribeEvaluationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DescribeEvaluationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DescribeEvaluationsOutput), req.Error
}

func (c *Client) DescribeMLModelsWithContext(ctx context.Context, input *machinelearning.DescribeMLModelsInput, opts ...request.Option) (*machinelearning.DescribeMLModelsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DescribeMLModelsWithContext",
		Input:   input,
		Output:  (*machinelearning.DescribeMLModelsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DescribeMLModelsWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DescribeMLModelsOutput), req.Error
}

func (c *Client) DescribeTagsWithContext(ctx context.Context, input *machinelearning.DescribeTagsInput, opts ...request.Option) (*machinelearning.DescribeTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "DescribeTagsWithContext",
		Input:   input,
		Output:  (*machinelearning.DescribeTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.DescribeTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.DescribeTagsOutput), req.Error
}

func (c *Client) GetBatchPredictionWithContext(ctx context.Context, input *machinelearning.GetBatchPredictionInput, opts ...request.Option) (*machinelearning.GetBatchPredictionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "GetBatchPredictionWithContext",
		Input:   input,
		Output:  (*machinelearning.GetBatchPredictionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.GetBatchPredictionWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.GetBatchPredictionOutput), req.Error
}

func (c *Client) GetDataSourceWithContext(ctx context.Context, input *machinelearning.GetDataSourceInput, opts ...request.Option) (*machinelearning.GetDataSourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "GetDataSourceWithContext",
		Input:   input,
		Output:  (*machinelearning.GetDataSourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.GetDataSourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.GetDataSourceOutput), req.Error
}

func (c *Client) GetEvaluationWithContext(ctx context.Context, input *machinelearning.GetEvaluationInput, opts ...request.Option) (*machinelearning.GetEvaluationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "GetEvaluationWithContext",
		Input:   input,
		Output:  (*machinelearning.GetEvaluationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.GetEvaluationWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.GetEvaluationOutput), req.Error
}

func (c *Client) GetMLModelWithContext(ctx context.Context, input *machinelearning.GetMLModelInput, opts ...request.Option) (*machinelearning.GetMLModelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "GetMLModelWithContext",
		Input:   input,
		Output:  (*machinelearning.GetMLModelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.GetMLModelWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.GetMLModelOutput), req.Error
}

func (c *Client) PredictWithContext(ctx context.Context, input *machinelearning.PredictInput, opts ...request.Option) (*machinelearning.PredictOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "PredictWithContext",
		Input:   input,
		Output:  (*machinelearning.PredictOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.PredictWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.PredictOutput), req.Error
}

func (c *Client) UpdateBatchPredictionWithContext(ctx context.Context, input *machinelearning.UpdateBatchPredictionInput, opts ...request.Option) (*machinelearning.UpdateBatchPredictionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "UpdateBatchPredictionWithContext",
		Input:   input,
		Output:  (*machinelearning.UpdateBatchPredictionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.UpdateBatchPredictionWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.UpdateBatchPredictionOutput), req.Error
}

func (c *Client) UpdateDataSourceWithContext(ctx context.Context, input *machinelearning.UpdateDataSourceInput, opts ...request.Option) (*machinelearning.UpdateDataSourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "UpdateDataSourceWithContext",
		Input:   input,
		Output:  (*machinelearning.UpdateDataSourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.UpdateDataSourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.UpdateDataSourceOutput), req.Error
}

func (c *Client) UpdateEvaluationWithContext(ctx context.Context, input *machinelearning.UpdateEvaluationInput, opts ...request.Option) (*machinelearning.UpdateEvaluationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "UpdateEvaluationWithContext",
		Input:   input,
		Output:  (*machinelearning.UpdateEvaluationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.UpdateEvaluationWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.UpdateEvaluationOutput), req.Error
}

func (c *Client) UpdateMLModelWithContext(ctx context.Context, input *machinelearning.UpdateMLModelInput, opts ...request.Option) (*machinelearning.UpdateMLModelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "machinelearning",
		Action:  "UpdateMLModelWithContext",
		Input:   input,
		Output:  (*machinelearning.UpdateMLModelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MachineLearningAPI.UpdateMLModelWithContext(ctx, input, opts...)
	})

	return req.Output.(*machinelearning.UpdateMLModelOutput), req.Error
}
