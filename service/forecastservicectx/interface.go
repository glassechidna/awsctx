// Code generated by internal/generate/main.go. DO NOT EDIT.

package forecastservicectx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"github.com/aws/aws-sdk-go/service/forecastservice/forecastserviceiface"
	"github.com/glassechidna/awsctx"
)

type ForecastService interface {
	CreateDatasetWithContext(ctx context.Context, input *forecastservice.CreateDatasetInput, opts ...request.Option) (*forecastservice.CreateDatasetOutput, error)
	CreateDatasetGroupWithContext(ctx context.Context, input *forecastservice.CreateDatasetGroupInput, opts ...request.Option) (*forecastservice.CreateDatasetGroupOutput, error)
	CreateDatasetImportJobWithContext(ctx context.Context, input *forecastservice.CreateDatasetImportJobInput, opts ...request.Option) (*forecastservice.CreateDatasetImportJobOutput, error)
	CreateForecastWithContext(ctx context.Context, input *forecastservice.CreateForecastInput, opts ...request.Option) (*forecastservice.CreateForecastOutput, error)
	CreateForecastExportJobWithContext(ctx context.Context, input *forecastservice.CreateForecastExportJobInput, opts ...request.Option) (*forecastservice.CreateForecastExportJobOutput, error)
	CreatePredictorWithContext(ctx context.Context, input *forecastservice.CreatePredictorInput, opts ...request.Option) (*forecastservice.CreatePredictorOutput, error)
	DeleteDatasetWithContext(ctx context.Context, input *forecastservice.DeleteDatasetInput, opts ...request.Option) (*forecastservice.DeleteDatasetOutput, error)
	DeleteDatasetGroupWithContext(ctx context.Context, input *forecastservice.DeleteDatasetGroupInput, opts ...request.Option) (*forecastservice.DeleteDatasetGroupOutput, error)
	DeleteDatasetImportJobWithContext(ctx context.Context, input *forecastservice.DeleteDatasetImportJobInput, opts ...request.Option) (*forecastservice.DeleteDatasetImportJobOutput, error)
	DeleteForecastWithContext(ctx context.Context, input *forecastservice.DeleteForecastInput, opts ...request.Option) (*forecastservice.DeleteForecastOutput, error)
	DeleteForecastExportJobWithContext(ctx context.Context, input *forecastservice.DeleteForecastExportJobInput, opts ...request.Option) (*forecastservice.DeleteForecastExportJobOutput, error)
	DeletePredictorWithContext(ctx context.Context, input *forecastservice.DeletePredictorInput, opts ...request.Option) (*forecastservice.DeletePredictorOutput, error)
	DescribeDatasetWithContext(ctx context.Context, input *forecastservice.DescribeDatasetInput, opts ...request.Option) (*forecastservice.DescribeDatasetOutput, error)
	DescribeDatasetGroupWithContext(ctx context.Context, input *forecastservice.DescribeDatasetGroupInput, opts ...request.Option) (*forecastservice.DescribeDatasetGroupOutput, error)
	DescribeDatasetImportJobWithContext(ctx context.Context, input *forecastservice.DescribeDatasetImportJobInput, opts ...request.Option) (*forecastservice.DescribeDatasetImportJobOutput, error)
	DescribeForecastWithContext(ctx context.Context, input *forecastservice.DescribeForecastInput, opts ...request.Option) (*forecastservice.DescribeForecastOutput, error)
	DescribeForecastExportJobWithContext(ctx context.Context, input *forecastservice.DescribeForecastExportJobInput, opts ...request.Option) (*forecastservice.DescribeForecastExportJobOutput, error)
	DescribePredictorWithContext(ctx context.Context, input *forecastservice.DescribePredictorInput, opts ...request.Option) (*forecastservice.DescribePredictorOutput, error)
	GetAccuracyMetricsWithContext(ctx context.Context, input *forecastservice.GetAccuracyMetricsInput, opts ...request.Option) (*forecastservice.GetAccuracyMetricsOutput, error)
	ListDatasetGroupsWithContext(ctx context.Context, input *forecastservice.ListDatasetGroupsInput, opts ...request.Option) (*forecastservice.ListDatasetGroupsOutput, error)
	ListDatasetImportJobsWithContext(ctx context.Context, input *forecastservice.ListDatasetImportJobsInput, opts ...request.Option) (*forecastservice.ListDatasetImportJobsOutput, error)
	ListDatasetsWithContext(ctx context.Context, input *forecastservice.ListDatasetsInput, opts ...request.Option) (*forecastservice.ListDatasetsOutput, error)
	ListForecastExportJobsWithContext(ctx context.Context, input *forecastservice.ListForecastExportJobsInput, opts ...request.Option) (*forecastservice.ListForecastExportJobsOutput, error)
	ListForecastsWithContext(ctx context.Context, input *forecastservice.ListForecastsInput, opts ...request.Option) (*forecastservice.ListForecastsOutput, error)
	ListPredictorsWithContext(ctx context.Context, input *forecastservice.ListPredictorsInput, opts ...request.Option) (*forecastservice.ListPredictorsOutput, error)
	UpdateDatasetGroupWithContext(ctx context.Context, input *forecastservice.UpdateDatasetGroupInput, opts ...request.Option) (*forecastservice.UpdateDatasetGroupOutput, error)
}

type Client struct {
	forecastserviceiface.ForecastServiceAPI
	Contexter awsctx.Contexter
}

func New(base forecastserviceiface.ForecastServiceAPI, ctxer awsctx.Contexter) ForecastService {
	return &Client{
		ForecastServiceAPI: base,
		Contexter: ctxer,
	}
}

var _ ForecastService = (*forecastservice.ForecastService)(nil)
var _ ForecastService = (*Client)(nil)

func (c *Client) CreateDatasetWithContext(ctx context.Context, input *forecastservice.CreateDatasetInput, opts ...request.Option) (*forecastservice.CreateDatasetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "CreateDataset",
		Input:   input,
		Output:  (*forecastservice.CreateDatasetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.CreateDatasetWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.CreateDatasetOutput), req.Error
}

func (c *Client) CreateDatasetGroupWithContext(ctx context.Context, input *forecastservice.CreateDatasetGroupInput, opts ...request.Option) (*forecastservice.CreateDatasetGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "CreateDatasetGroup",
		Input:   input,
		Output:  (*forecastservice.CreateDatasetGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.CreateDatasetGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.CreateDatasetGroupOutput), req.Error
}

func (c *Client) CreateDatasetImportJobWithContext(ctx context.Context, input *forecastservice.CreateDatasetImportJobInput, opts ...request.Option) (*forecastservice.CreateDatasetImportJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "CreateDatasetImportJob",
		Input:   input,
		Output:  (*forecastservice.CreateDatasetImportJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.CreateDatasetImportJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.CreateDatasetImportJobOutput), req.Error
}

func (c *Client) CreateForecastWithContext(ctx context.Context, input *forecastservice.CreateForecastInput, opts ...request.Option) (*forecastservice.CreateForecastOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "CreateForecast",
		Input:   input,
		Output:  (*forecastservice.CreateForecastOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.CreateForecastWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.CreateForecastOutput), req.Error
}

func (c *Client) CreateForecastExportJobWithContext(ctx context.Context, input *forecastservice.CreateForecastExportJobInput, opts ...request.Option) (*forecastservice.CreateForecastExportJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "CreateForecastExportJob",
		Input:   input,
		Output:  (*forecastservice.CreateForecastExportJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.CreateForecastExportJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.CreateForecastExportJobOutput), req.Error
}

func (c *Client) CreatePredictorWithContext(ctx context.Context, input *forecastservice.CreatePredictorInput, opts ...request.Option) (*forecastservice.CreatePredictorOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "CreatePredictor",
		Input:   input,
		Output:  (*forecastservice.CreatePredictorOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.CreatePredictorWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.CreatePredictorOutput), req.Error
}

func (c *Client) DeleteDatasetWithContext(ctx context.Context, input *forecastservice.DeleteDatasetInput, opts ...request.Option) (*forecastservice.DeleteDatasetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DeleteDataset",
		Input:   input,
		Output:  (*forecastservice.DeleteDatasetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DeleteDatasetWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DeleteDatasetOutput), req.Error
}

func (c *Client) DeleteDatasetGroupWithContext(ctx context.Context, input *forecastservice.DeleteDatasetGroupInput, opts ...request.Option) (*forecastservice.DeleteDatasetGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DeleteDatasetGroup",
		Input:   input,
		Output:  (*forecastservice.DeleteDatasetGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DeleteDatasetGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DeleteDatasetGroupOutput), req.Error
}

func (c *Client) DeleteDatasetImportJobWithContext(ctx context.Context, input *forecastservice.DeleteDatasetImportJobInput, opts ...request.Option) (*forecastservice.DeleteDatasetImportJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DeleteDatasetImportJob",
		Input:   input,
		Output:  (*forecastservice.DeleteDatasetImportJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DeleteDatasetImportJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DeleteDatasetImportJobOutput), req.Error
}

func (c *Client) DeleteForecastWithContext(ctx context.Context, input *forecastservice.DeleteForecastInput, opts ...request.Option) (*forecastservice.DeleteForecastOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DeleteForecast",
		Input:   input,
		Output:  (*forecastservice.DeleteForecastOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DeleteForecastWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DeleteForecastOutput), req.Error
}

func (c *Client) DeleteForecastExportJobWithContext(ctx context.Context, input *forecastservice.DeleteForecastExportJobInput, opts ...request.Option) (*forecastservice.DeleteForecastExportJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DeleteForecastExportJob",
		Input:   input,
		Output:  (*forecastservice.DeleteForecastExportJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DeleteForecastExportJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DeleteForecastExportJobOutput), req.Error
}

func (c *Client) DeletePredictorWithContext(ctx context.Context, input *forecastservice.DeletePredictorInput, opts ...request.Option) (*forecastservice.DeletePredictorOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DeletePredictor",
		Input:   input,
		Output:  (*forecastservice.DeletePredictorOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DeletePredictorWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DeletePredictorOutput), req.Error
}

func (c *Client) DescribeDatasetWithContext(ctx context.Context, input *forecastservice.DescribeDatasetInput, opts ...request.Option) (*forecastservice.DescribeDatasetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DescribeDataset",
		Input:   input,
		Output:  (*forecastservice.DescribeDatasetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DescribeDatasetWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DescribeDatasetOutput), req.Error
}

func (c *Client) DescribeDatasetGroupWithContext(ctx context.Context, input *forecastservice.DescribeDatasetGroupInput, opts ...request.Option) (*forecastservice.DescribeDatasetGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DescribeDatasetGroup",
		Input:   input,
		Output:  (*forecastservice.DescribeDatasetGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DescribeDatasetGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DescribeDatasetGroupOutput), req.Error
}

func (c *Client) DescribeDatasetImportJobWithContext(ctx context.Context, input *forecastservice.DescribeDatasetImportJobInput, opts ...request.Option) (*forecastservice.DescribeDatasetImportJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DescribeDatasetImportJob",
		Input:   input,
		Output:  (*forecastservice.DescribeDatasetImportJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DescribeDatasetImportJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DescribeDatasetImportJobOutput), req.Error
}

func (c *Client) DescribeForecastWithContext(ctx context.Context, input *forecastservice.DescribeForecastInput, opts ...request.Option) (*forecastservice.DescribeForecastOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DescribeForecast",
		Input:   input,
		Output:  (*forecastservice.DescribeForecastOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DescribeForecastWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DescribeForecastOutput), req.Error
}

func (c *Client) DescribeForecastExportJobWithContext(ctx context.Context, input *forecastservice.DescribeForecastExportJobInput, opts ...request.Option) (*forecastservice.DescribeForecastExportJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DescribeForecastExportJob",
		Input:   input,
		Output:  (*forecastservice.DescribeForecastExportJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DescribeForecastExportJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DescribeForecastExportJobOutput), req.Error
}

func (c *Client) DescribePredictorWithContext(ctx context.Context, input *forecastservice.DescribePredictorInput, opts ...request.Option) (*forecastservice.DescribePredictorOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "DescribePredictor",
		Input:   input,
		Output:  (*forecastservice.DescribePredictorOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.DescribePredictorWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.DescribePredictorOutput), req.Error
}

func (c *Client) GetAccuracyMetricsWithContext(ctx context.Context, input *forecastservice.GetAccuracyMetricsInput, opts ...request.Option) (*forecastservice.GetAccuracyMetricsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "GetAccuracyMetrics",
		Input:   input,
		Output:  (*forecastservice.GetAccuracyMetricsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.GetAccuracyMetricsWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.GetAccuracyMetricsOutput), req.Error
}

func (c *Client) ListDatasetGroupsWithContext(ctx context.Context, input *forecastservice.ListDatasetGroupsInput, opts ...request.Option) (*forecastservice.ListDatasetGroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "ListDatasetGroups",
		Input:   input,
		Output:  (*forecastservice.ListDatasetGroupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.ListDatasetGroupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.ListDatasetGroupsOutput), req.Error
}

func (c *Client) ListDatasetImportJobsWithContext(ctx context.Context, input *forecastservice.ListDatasetImportJobsInput, opts ...request.Option) (*forecastservice.ListDatasetImportJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "ListDatasetImportJobs",
		Input:   input,
		Output:  (*forecastservice.ListDatasetImportJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.ListDatasetImportJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.ListDatasetImportJobsOutput), req.Error
}

func (c *Client) ListDatasetsWithContext(ctx context.Context, input *forecastservice.ListDatasetsInput, opts ...request.Option) (*forecastservice.ListDatasetsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "ListDatasets",
		Input:   input,
		Output:  (*forecastservice.ListDatasetsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.ListDatasetsWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.ListDatasetsOutput), req.Error
}

func (c *Client) ListForecastExportJobsWithContext(ctx context.Context, input *forecastservice.ListForecastExportJobsInput, opts ...request.Option) (*forecastservice.ListForecastExportJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "ListForecastExportJobs",
		Input:   input,
		Output:  (*forecastservice.ListForecastExportJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.ListForecastExportJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.ListForecastExportJobsOutput), req.Error
}

func (c *Client) ListForecastsWithContext(ctx context.Context, input *forecastservice.ListForecastsInput, opts ...request.Option) (*forecastservice.ListForecastsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "ListForecasts",
		Input:   input,
		Output:  (*forecastservice.ListForecastsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.ListForecastsWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.ListForecastsOutput), req.Error
}

func (c *Client) ListPredictorsWithContext(ctx context.Context, input *forecastservice.ListPredictorsInput, opts ...request.Option) (*forecastservice.ListPredictorsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "ListPredictors",
		Input:   input,
		Output:  (*forecastservice.ListPredictorsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.ListPredictorsWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.ListPredictorsOutput), req.Error
}

func (c *Client) UpdateDatasetGroupWithContext(ctx context.Context, input *forecastservice.UpdateDatasetGroupInput, opts ...request.Option) (*forecastservice.UpdateDatasetGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "forecastservice",
		Action:  "UpdateDatasetGroup",
		Input:   input,
		Output:  (*forecastservice.UpdateDatasetGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ForecastServiceAPI.UpdateDatasetGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*forecastservice.UpdateDatasetGroupOutput), req.Error
}
