// Code generated by internal/generate/main.go. DO NOT EDIT.

package costexplorerctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
	"github.com/glassechidna/awsctx"
)

type CostExplorer interface {
	CreateAnomalyMonitorWithContext(ctx context.Context, input *costexplorer.CreateAnomalyMonitorInput, opts ...request.Option) (*costexplorer.CreateAnomalyMonitorOutput, error)
	CreateAnomalySubscriptionWithContext(ctx context.Context, input *costexplorer.CreateAnomalySubscriptionInput, opts ...request.Option) (*costexplorer.CreateAnomalySubscriptionOutput, error)
	CreateCostCategoryDefinitionWithContext(ctx context.Context, input *costexplorer.CreateCostCategoryDefinitionInput, opts ...request.Option) (*costexplorer.CreateCostCategoryDefinitionOutput, error)
	DeleteAnomalyMonitorWithContext(ctx context.Context, input *costexplorer.DeleteAnomalyMonitorInput, opts ...request.Option) (*costexplorer.DeleteAnomalyMonitorOutput, error)
	DeleteAnomalySubscriptionWithContext(ctx context.Context, input *costexplorer.DeleteAnomalySubscriptionInput, opts ...request.Option) (*costexplorer.DeleteAnomalySubscriptionOutput, error)
	DeleteCostCategoryDefinitionWithContext(ctx context.Context, input *costexplorer.DeleteCostCategoryDefinitionInput, opts ...request.Option) (*costexplorer.DeleteCostCategoryDefinitionOutput, error)
	DescribeCostCategoryDefinitionWithContext(ctx context.Context, input *costexplorer.DescribeCostCategoryDefinitionInput, opts ...request.Option) (*costexplorer.DescribeCostCategoryDefinitionOutput, error)
	GetAnomaliesWithContext(ctx context.Context, input *costexplorer.GetAnomaliesInput, opts ...request.Option) (*costexplorer.GetAnomaliesOutput, error)
	GetAnomalyMonitorsWithContext(ctx context.Context, input *costexplorer.GetAnomalyMonitorsInput, opts ...request.Option) (*costexplorer.GetAnomalyMonitorsOutput, error)
	GetAnomalySubscriptionsWithContext(ctx context.Context, input *costexplorer.GetAnomalySubscriptionsInput, opts ...request.Option) (*costexplorer.GetAnomalySubscriptionsOutput, error)
	GetCostAndUsageWithContext(ctx context.Context, input *costexplorer.GetCostAndUsageInput, opts ...request.Option) (*costexplorer.GetCostAndUsageOutput, error)
	GetCostAndUsageWithResourcesWithContext(ctx context.Context, input *costexplorer.GetCostAndUsageWithResourcesInput, opts ...request.Option) (*costexplorer.GetCostAndUsageWithResourcesOutput, error)
	GetCostCategoriesWithContext(ctx context.Context, input *costexplorer.GetCostCategoriesInput, opts ...request.Option) (*costexplorer.GetCostCategoriesOutput, error)
	GetCostForecastWithContext(ctx context.Context, input *costexplorer.GetCostForecastInput, opts ...request.Option) (*costexplorer.GetCostForecastOutput, error)
	GetDimensionValuesWithContext(ctx context.Context, input *costexplorer.GetDimensionValuesInput, opts ...request.Option) (*costexplorer.GetDimensionValuesOutput, error)
	GetReservationCoverageWithContext(ctx context.Context, input *costexplorer.GetReservationCoverageInput, opts ...request.Option) (*costexplorer.GetReservationCoverageOutput, error)
	GetReservationPurchaseRecommendationWithContext(ctx context.Context, input *costexplorer.GetReservationPurchaseRecommendationInput, opts ...request.Option) (*costexplorer.GetReservationPurchaseRecommendationOutput, error)
	GetReservationUtilizationWithContext(ctx context.Context, input *costexplorer.GetReservationUtilizationInput, opts ...request.Option) (*costexplorer.GetReservationUtilizationOutput, error)
	GetRightsizingRecommendationWithContext(ctx context.Context, input *costexplorer.GetRightsizingRecommendationInput, opts ...request.Option) (*costexplorer.GetRightsizingRecommendationOutput, error)
	GetSavingsPlansCoverageWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansCoverageInput, opts ...request.Option) (*costexplorer.GetSavingsPlansCoverageOutput, error)
	GetSavingsPlansCoveragePagesWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansCoverageInput, cb func(*costexplorer.GetSavingsPlansCoverageOutput, bool) bool, opts ...request.Option) error
	GetSavingsPlansPurchaseRecommendationWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput, opts ...request.Option) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error)
	GetSavingsPlansUtilizationWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationInput, opts ...request.Option) (*costexplorer.GetSavingsPlansUtilizationOutput, error)
	GetSavingsPlansUtilizationDetailsWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput, opts ...request.Option) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error)
	GetSavingsPlansUtilizationDetailsPagesWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput, cb func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) bool, opts ...request.Option) error
	GetTagsWithContext(ctx context.Context, input *costexplorer.GetTagsInput, opts ...request.Option) (*costexplorer.GetTagsOutput, error)
	GetUsageForecastWithContext(ctx context.Context, input *costexplorer.GetUsageForecastInput, opts ...request.Option) (*costexplorer.GetUsageForecastOutput, error)
	ListCostAllocationTagsWithContext(ctx context.Context, input *costexplorer.ListCostAllocationTagsInput, opts ...request.Option) (*costexplorer.ListCostAllocationTagsOutput, error)
	ListCostAllocationTagsPagesWithContext(ctx context.Context, input *costexplorer.ListCostAllocationTagsInput, cb func(*costexplorer.ListCostAllocationTagsOutput, bool) bool, opts ...request.Option) error
	ListCostCategoryDefinitionsWithContext(ctx context.Context, input *costexplorer.ListCostCategoryDefinitionsInput, opts ...request.Option) (*costexplorer.ListCostCategoryDefinitionsOutput, error)
	ListCostCategoryDefinitionsPagesWithContext(ctx context.Context, input *costexplorer.ListCostCategoryDefinitionsInput, cb func(*costexplorer.ListCostCategoryDefinitionsOutput, bool) bool, opts ...request.Option) error
	ListSavingsPlansPurchaseRecommendationGenerationWithContext(ctx context.Context, input *costexplorer.ListSavingsPlansPurchaseRecommendationGenerationInput, opts ...request.Option) (*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *costexplorer.ListTagsForResourceInput, opts ...request.Option) (*costexplorer.ListTagsForResourceOutput, error)
	ProvideAnomalyFeedbackWithContext(ctx context.Context, input *costexplorer.ProvideAnomalyFeedbackInput, opts ...request.Option) (*costexplorer.ProvideAnomalyFeedbackOutput, error)
	StartSavingsPlansPurchaseRecommendationGenerationWithContext(ctx context.Context, input *costexplorer.StartSavingsPlansPurchaseRecommendationGenerationInput, opts ...request.Option) (*costexplorer.StartSavingsPlansPurchaseRecommendationGenerationOutput, error)
	TagResourceWithContext(ctx context.Context, input *costexplorer.TagResourceInput, opts ...request.Option) (*costexplorer.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *costexplorer.UntagResourceInput, opts ...request.Option) (*costexplorer.UntagResourceOutput, error)
	UpdateAnomalyMonitorWithContext(ctx context.Context, input *costexplorer.UpdateAnomalyMonitorInput, opts ...request.Option) (*costexplorer.UpdateAnomalyMonitorOutput, error)
	UpdateAnomalySubscriptionWithContext(ctx context.Context, input *costexplorer.UpdateAnomalySubscriptionInput, opts ...request.Option) (*costexplorer.UpdateAnomalySubscriptionOutput, error)
	UpdateCostAllocationTagsStatusWithContext(ctx context.Context, input *costexplorer.UpdateCostAllocationTagsStatusInput, opts ...request.Option) (*costexplorer.UpdateCostAllocationTagsStatusOutput, error)
	UpdateCostCategoryDefinitionWithContext(ctx context.Context, input *costexplorer.UpdateCostCategoryDefinitionInput, opts ...request.Option) (*costexplorer.UpdateCostCategoryDefinitionOutput, error)
}

type Client struct {
	costexploreriface.CostExplorerAPI
	Contexter awsctx.Contexter
}

func New(base costexploreriface.CostExplorerAPI, ctxer awsctx.Contexter) CostExplorer {
	return &Client{
		CostExplorerAPI: base,
		Contexter: ctxer,
	}
}

var _ CostExplorer = (*costexplorer.CostExplorer)(nil)
var _ CostExplorer = (*Client)(nil)

func (c *Client) CreateAnomalyMonitorWithContext(ctx context.Context, input *costexplorer.CreateAnomalyMonitorInput, opts ...request.Option) (*costexplorer.CreateAnomalyMonitorOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "CreateAnomalyMonitor",
		Input:   input,
		Output:  (*costexplorer.CreateAnomalyMonitorOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.CreateAnomalyMonitorWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.CreateAnomalyMonitorOutput), req.Error
}

func (c *Client) CreateAnomalySubscriptionWithContext(ctx context.Context, input *costexplorer.CreateAnomalySubscriptionInput, opts ...request.Option) (*costexplorer.CreateAnomalySubscriptionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "CreateAnomalySubscription",
		Input:   input,
		Output:  (*costexplorer.CreateAnomalySubscriptionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.CreateAnomalySubscriptionWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.CreateAnomalySubscriptionOutput), req.Error
}

func (c *Client) CreateCostCategoryDefinitionWithContext(ctx context.Context, input *costexplorer.CreateCostCategoryDefinitionInput, opts ...request.Option) (*costexplorer.CreateCostCategoryDefinitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "CreateCostCategoryDefinition",
		Input:   input,
		Output:  (*costexplorer.CreateCostCategoryDefinitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.CreateCostCategoryDefinitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.CreateCostCategoryDefinitionOutput), req.Error
}

func (c *Client) DeleteAnomalyMonitorWithContext(ctx context.Context, input *costexplorer.DeleteAnomalyMonitorInput, opts ...request.Option) (*costexplorer.DeleteAnomalyMonitorOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "DeleteAnomalyMonitor",
		Input:   input,
		Output:  (*costexplorer.DeleteAnomalyMonitorOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.DeleteAnomalyMonitorWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.DeleteAnomalyMonitorOutput), req.Error
}

func (c *Client) DeleteAnomalySubscriptionWithContext(ctx context.Context, input *costexplorer.DeleteAnomalySubscriptionInput, opts ...request.Option) (*costexplorer.DeleteAnomalySubscriptionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "DeleteAnomalySubscription",
		Input:   input,
		Output:  (*costexplorer.DeleteAnomalySubscriptionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.DeleteAnomalySubscriptionWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.DeleteAnomalySubscriptionOutput), req.Error
}

func (c *Client) DeleteCostCategoryDefinitionWithContext(ctx context.Context, input *costexplorer.DeleteCostCategoryDefinitionInput, opts ...request.Option) (*costexplorer.DeleteCostCategoryDefinitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "DeleteCostCategoryDefinition",
		Input:   input,
		Output:  (*costexplorer.DeleteCostCategoryDefinitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.DeleteCostCategoryDefinitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.DeleteCostCategoryDefinitionOutput), req.Error
}

func (c *Client) DescribeCostCategoryDefinitionWithContext(ctx context.Context, input *costexplorer.DescribeCostCategoryDefinitionInput, opts ...request.Option) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "DescribeCostCategoryDefinition",
		Input:   input,
		Output:  (*costexplorer.DescribeCostCategoryDefinitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.DescribeCostCategoryDefinitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.DescribeCostCategoryDefinitionOutput), req.Error
}

func (c *Client) GetAnomaliesWithContext(ctx context.Context, input *costexplorer.GetAnomaliesInput, opts ...request.Option) (*costexplorer.GetAnomaliesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetAnomalies",
		Input:   input,
		Output:  (*costexplorer.GetAnomaliesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetAnomaliesWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetAnomaliesOutput), req.Error
}

func (c *Client) GetAnomalyMonitorsWithContext(ctx context.Context, input *costexplorer.GetAnomalyMonitorsInput, opts ...request.Option) (*costexplorer.GetAnomalyMonitorsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetAnomalyMonitors",
		Input:   input,
		Output:  (*costexplorer.GetAnomalyMonitorsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetAnomalyMonitorsWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetAnomalyMonitorsOutput), req.Error
}

func (c *Client) GetAnomalySubscriptionsWithContext(ctx context.Context, input *costexplorer.GetAnomalySubscriptionsInput, opts ...request.Option) (*costexplorer.GetAnomalySubscriptionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetAnomalySubscriptions",
		Input:   input,
		Output:  (*costexplorer.GetAnomalySubscriptionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetAnomalySubscriptionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetAnomalySubscriptionsOutput), req.Error
}

func (c *Client) GetCostAndUsageWithContext(ctx context.Context, input *costexplorer.GetCostAndUsageInput, opts ...request.Option) (*costexplorer.GetCostAndUsageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetCostAndUsage",
		Input:   input,
		Output:  (*costexplorer.GetCostAndUsageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetCostAndUsageWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetCostAndUsageOutput), req.Error
}

func (c *Client) GetCostAndUsageWithResourcesWithContext(ctx context.Context, input *costexplorer.GetCostAndUsageWithResourcesInput, opts ...request.Option) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetCostAndUsageWithResources",
		Input:   input,
		Output:  (*costexplorer.GetCostAndUsageWithResourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetCostAndUsageWithResourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetCostAndUsageWithResourcesOutput), req.Error
}

func (c *Client) GetCostCategoriesWithContext(ctx context.Context, input *costexplorer.GetCostCategoriesInput, opts ...request.Option) (*costexplorer.GetCostCategoriesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetCostCategories",
		Input:   input,
		Output:  (*costexplorer.GetCostCategoriesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetCostCategoriesWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetCostCategoriesOutput), req.Error
}

func (c *Client) GetCostForecastWithContext(ctx context.Context, input *costexplorer.GetCostForecastInput, opts ...request.Option) (*costexplorer.GetCostForecastOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetCostForecast",
		Input:   input,
		Output:  (*costexplorer.GetCostForecastOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetCostForecastWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetCostForecastOutput), req.Error
}

func (c *Client) GetDimensionValuesWithContext(ctx context.Context, input *costexplorer.GetDimensionValuesInput, opts ...request.Option) (*costexplorer.GetDimensionValuesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetDimensionValues",
		Input:   input,
		Output:  (*costexplorer.GetDimensionValuesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetDimensionValuesWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetDimensionValuesOutput), req.Error
}

func (c *Client) GetReservationCoverageWithContext(ctx context.Context, input *costexplorer.GetReservationCoverageInput, opts ...request.Option) (*costexplorer.GetReservationCoverageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetReservationCoverage",
		Input:   input,
		Output:  (*costexplorer.GetReservationCoverageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetReservationCoverageWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetReservationCoverageOutput), req.Error
}

func (c *Client) GetReservationPurchaseRecommendationWithContext(ctx context.Context, input *costexplorer.GetReservationPurchaseRecommendationInput, opts ...request.Option) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetReservationPurchaseRecommendation",
		Input:   input,
		Output:  (*costexplorer.GetReservationPurchaseRecommendationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetReservationPurchaseRecommendationWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetReservationPurchaseRecommendationOutput), req.Error
}

func (c *Client) GetReservationUtilizationWithContext(ctx context.Context, input *costexplorer.GetReservationUtilizationInput, opts ...request.Option) (*costexplorer.GetReservationUtilizationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetReservationUtilization",
		Input:   input,
		Output:  (*costexplorer.GetReservationUtilizationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetReservationUtilizationWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetReservationUtilizationOutput), req.Error
}

func (c *Client) GetRightsizingRecommendationWithContext(ctx context.Context, input *costexplorer.GetRightsizingRecommendationInput, opts ...request.Option) (*costexplorer.GetRightsizingRecommendationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetRightsizingRecommendation",
		Input:   input,
		Output:  (*costexplorer.GetRightsizingRecommendationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetRightsizingRecommendationWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetRightsizingRecommendationOutput), req.Error
}

func (c *Client) GetSavingsPlansCoverageWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansCoverageInput, opts ...request.Option) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetSavingsPlansCoverage",
		Input:   input,
		Output:  (*costexplorer.GetSavingsPlansCoverageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetSavingsPlansCoverageWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetSavingsPlansCoverageOutput), req.Error
}

func (c *Client) GetSavingsPlansCoveragePagesWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansCoverageInput, cb func(*costexplorer.GetSavingsPlansCoverageOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetSavingsPlansCoverage",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CostExplorerAPI.GetSavingsPlansCoveragePagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetSavingsPlansPurchaseRecommendationWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput, opts ...request.Option) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetSavingsPlansPurchaseRecommendation",
		Input:   input,
		Output:  (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetSavingsPlansPurchaseRecommendationWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetSavingsPlansPurchaseRecommendationOutput), req.Error
}

func (c *Client) GetSavingsPlansUtilizationWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationInput, opts ...request.Option) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetSavingsPlansUtilization",
		Input:   input,
		Output:  (*costexplorer.GetSavingsPlansUtilizationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetSavingsPlansUtilizationWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetSavingsPlansUtilizationOutput), req.Error
}

func (c *Client) GetSavingsPlansUtilizationDetailsWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput, opts ...request.Option) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetSavingsPlansUtilizationDetails",
		Input:   input,
		Output:  (*costexplorer.GetSavingsPlansUtilizationDetailsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetSavingsPlansUtilizationDetailsWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetSavingsPlansUtilizationDetailsOutput), req.Error
}

func (c *Client) GetSavingsPlansUtilizationDetailsPagesWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput, cb func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetSavingsPlansUtilizationDetails",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CostExplorerAPI.GetSavingsPlansUtilizationDetailsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetTagsWithContext(ctx context.Context, input *costexplorer.GetTagsInput, opts ...request.Option) (*costexplorer.GetTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetTags",
		Input:   input,
		Output:  (*costexplorer.GetTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetTagsOutput), req.Error
}

func (c *Client) GetUsageForecastWithContext(ctx context.Context, input *costexplorer.GetUsageForecastInput, opts ...request.Option) (*costexplorer.GetUsageForecastOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "GetUsageForecast",
		Input:   input,
		Output:  (*costexplorer.GetUsageForecastOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.GetUsageForecastWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.GetUsageForecastOutput), req.Error
}

func (c *Client) ListCostAllocationTagsWithContext(ctx context.Context, input *costexplorer.ListCostAllocationTagsInput, opts ...request.Option) (*costexplorer.ListCostAllocationTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "ListCostAllocationTags",
		Input:   input,
		Output:  (*costexplorer.ListCostAllocationTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.ListCostAllocationTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.ListCostAllocationTagsOutput), req.Error
}

func (c *Client) ListCostAllocationTagsPagesWithContext(ctx context.Context, input *costexplorer.ListCostAllocationTagsInput, cb func(*costexplorer.ListCostAllocationTagsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "ListCostAllocationTags",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CostExplorerAPI.ListCostAllocationTagsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListCostCategoryDefinitionsWithContext(ctx context.Context, input *costexplorer.ListCostCategoryDefinitionsInput, opts ...request.Option) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "ListCostCategoryDefinitions",
		Input:   input,
		Output:  (*costexplorer.ListCostCategoryDefinitionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.ListCostCategoryDefinitionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.ListCostCategoryDefinitionsOutput), req.Error
}

func (c *Client) ListCostCategoryDefinitionsPagesWithContext(ctx context.Context, input *costexplorer.ListCostCategoryDefinitionsInput, cb func(*costexplorer.ListCostCategoryDefinitionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "ListCostCategoryDefinitions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.CostExplorerAPI.ListCostCategoryDefinitionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListSavingsPlansPurchaseRecommendationGenerationWithContext(ctx context.Context, input *costexplorer.ListSavingsPlansPurchaseRecommendationGenerationInput, opts ...request.Option) (*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "ListSavingsPlansPurchaseRecommendationGeneration",
		Input:   input,
		Output:  (*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.ListSavingsPlansPurchaseRecommendationGenerationWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *costexplorer.ListTagsForResourceInput, opts ...request.Option) (*costexplorer.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*costexplorer.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.ListTagsForResourceOutput), req.Error
}

func (c *Client) ProvideAnomalyFeedbackWithContext(ctx context.Context, input *costexplorer.ProvideAnomalyFeedbackInput, opts ...request.Option) (*costexplorer.ProvideAnomalyFeedbackOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "ProvideAnomalyFeedback",
		Input:   input,
		Output:  (*costexplorer.ProvideAnomalyFeedbackOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.ProvideAnomalyFeedbackWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.ProvideAnomalyFeedbackOutput), req.Error
}

func (c *Client) StartSavingsPlansPurchaseRecommendationGenerationWithContext(ctx context.Context, input *costexplorer.StartSavingsPlansPurchaseRecommendationGenerationInput, opts ...request.Option) (*costexplorer.StartSavingsPlansPurchaseRecommendationGenerationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "StartSavingsPlansPurchaseRecommendationGeneration",
		Input:   input,
		Output:  (*costexplorer.StartSavingsPlansPurchaseRecommendationGenerationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.StartSavingsPlansPurchaseRecommendationGenerationWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.StartSavingsPlansPurchaseRecommendationGenerationOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *costexplorer.TagResourceInput, opts ...request.Option) (*costexplorer.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "TagResource",
		Input:   input,
		Output:  (*costexplorer.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *costexplorer.UntagResourceInput, opts ...request.Option) (*costexplorer.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*costexplorer.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.UntagResourceOutput), req.Error
}

func (c *Client) UpdateAnomalyMonitorWithContext(ctx context.Context, input *costexplorer.UpdateAnomalyMonitorInput, opts ...request.Option) (*costexplorer.UpdateAnomalyMonitorOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "UpdateAnomalyMonitor",
		Input:   input,
		Output:  (*costexplorer.UpdateAnomalyMonitorOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.UpdateAnomalyMonitorWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.UpdateAnomalyMonitorOutput), req.Error
}

func (c *Client) UpdateAnomalySubscriptionWithContext(ctx context.Context, input *costexplorer.UpdateAnomalySubscriptionInput, opts ...request.Option) (*costexplorer.UpdateAnomalySubscriptionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "UpdateAnomalySubscription",
		Input:   input,
		Output:  (*costexplorer.UpdateAnomalySubscriptionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.UpdateAnomalySubscriptionWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.UpdateAnomalySubscriptionOutput), req.Error
}

func (c *Client) UpdateCostAllocationTagsStatusWithContext(ctx context.Context, input *costexplorer.UpdateCostAllocationTagsStatusInput, opts ...request.Option) (*costexplorer.UpdateCostAllocationTagsStatusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "UpdateCostAllocationTagsStatus",
		Input:   input,
		Output:  (*costexplorer.UpdateCostAllocationTagsStatusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.UpdateCostAllocationTagsStatusWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.UpdateCostAllocationTagsStatusOutput), req.Error
}

func (c *Client) UpdateCostCategoryDefinitionWithContext(ctx context.Context, input *costexplorer.UpdateCostCategoryDefinitionInput, opts ...request.Option) (*costexplorer.UpdateCostCategoryDefinitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "costexplorer",
		Action:  "UpdateCostCategoryDefinition",
		Input:   input,
		Output:  (*costexplorer.UpdateCostCategoryDefinitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CostExplorerAPI.UpdateCostCategoryDefinitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*costexplorer.UpdateCostCategoryDefinitionOutput), req.Error
}
