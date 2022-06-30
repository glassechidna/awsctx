// Code generated by internal/generate/main.go. DO NOT EDIT.

package athenactx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
	"github.com/glassechidna/awsctx"
)

type Athena interface {
	BatchGetNamedQueryWithContext(ctx context.Context, input *athena.BatchGetNamedQueryInput, opts ...request.Option) (*athena.BatchGetNamedQueryOutput, error)
	BatchGetPreparedStatementWithContext(ctx context.Context, input *athena.BatchGetPreparedStatementInput, opts ...request.Option) (*athena.BatchGetPreparedStatementOutput, error)
	BatchGetQueryExecutionWithContext(ctx context.Context, input *athena.BatchGetQueryExecutionInput, opts ...request.Option) (*athena.BatchGetQueryExecutionOutput, error)
	CreateDataCatalogWithContext(ctx context.Context, input *athena.CreateDataCatalogInput, opts ...request.Option) (*athena.CreateDataCatalogOutput, error)
	CreateNamedQueryWithContext(ctx context.Context, input *athena.CreateNamedQueryInput, opts ...request.Option) (*athena.CreateNamedQueryOutput, error)
	CreatePreparedStatementWithContext(ctx context.Context, input *athena.CreatePreparedStatementInput, opts ...request.Option) (*athena.CreatePreparedStatementOutput, error)
	CreateWorkGroupWithContext(ctx context.Context, input *athena.CreateWorkGroupInput, opts ...request.Option) (*athena.CreateWorkGroupOutput, error)
	DeleteDataCatalogWithContext(ctx context.Context, input *athena.DeleteDataCatalogInput, opts ...request.Option) (*athena.DeleteDataCatalogOutput, error)
	DeleteNamedQueryWithContext(ctx context.Context, input *athena.DeleteNamedQueryInput, opts ...request.Option) (*athena.DeleteNamedQueryOutput, error)
	DeletePreparedStatementWithContext(ctx context.Context, input *athena.DeletePreparedStatementInput, opts ...request.Option) (*athena.DeletePreparedStatementOutput, error)
	DeleteWorkGroupWithContext(ctx context.Context, input *athena.DeleteWorkGroupInput, opts ...request.Option) (*athena.DeleteWorkGroupOutput, error)
	GetDataCatalogWithContext(ctx context.Context, input *athena.GetDataCatalogInput, opts ...request.Option) (*athena.GetDataCatalogOutput, error)
	GetDatabaseWithContext(ctx context.Context, input *athena.GetDatabaseInput, opts ...request.Option) (*athena.GetDatabaseOutput, error)
	GetNamedQueryWithContext(ctx context.Context, input *athena.GetNamedQueryInput, opts ...request.Option) (*athena.GetNamedQueryOutput, error)
	GetPreparedStatementWithContext(ctx context.Context, input *athena.GetPreparedStatementInput, opts ...request.Option) (*athena.GetPreparedStatementOutput, error)
	GetQueryExecutionWithContext(ctx context.Context, input *athena.GetQueryExecutionInput, opts ...request.Option) (*athena.GetQueryExecutionOutput, error)
	GetQueryResultsWithContext(ctx context.Context, input *athena.GetQueryResultsInput, opts ...request.Option) (*athena.GetQueryResultsOutput, error)
	GetQueryResultsPagesWithContext(ctx context.Context, input *athena.GetQueryResultsInput, cb func(*athena.GetQueryResultsOutput, bool) bool, opts ...request.Option) error
	GetTableMetadataWithContext(ctx context.Context, input *athena.GetTableMetadataInput, opts ...request.Option) (*athena.GetTableMetadataOutput, error)
	GetWorkGroupWithContext(ctx context.Context, input *athena.GetWorkGroupInput, opts ...request.Option) (*athena.GetWorkGroupOutput, error)
	ListDataCatalogsWithContext(ctx context.Context, input *athena.ListDataCatalogsInput, opts ...request.Option) (*athena.ListDataCatalogsOutput, error)
	ListDataCatalogsPagesWithContext(ctx context.Context, input *athena.ListDataCatalogsInput, cb func(*athena.ListDataCatalogsOutput, bool) bool, opts ...request.Option) error
	ListDatabasesWithContext(ctx context.Context, input *athena.ListDatabasesInput, opts ...request.Option) (*athena.ListDatabasesOutput, error)
	ListDatabasesPagesWithContext(ctx context.Context, input *athena.ListDatabasesInput, cb func(*athena.ListDatabasesOutput, bool) bool, opts ...request.Option) error
	ListEngineVersionsWithContext(ctx context.Context, input *athena.ListEngineVersionsInput, opts ...request.Option) (*athena.ListEngineVersionsOutput, error)
	ListNamedQueriesWithContext(ctx context.Context, input *athena.ListNamedQueriesInput, opts ...request.Option) (*athena.ListNamedQueriesOutput, error)
	ListNamedQueriesPagesWithContext(ctx context.Context, input *athena.ListNamedQueriesInput, cb func(*athena.ListNamedQueriesOutput, bool) bool, opts ...request.Option) error
	ListPreparedStatementsWithContext(ctx context.Context, input *athena.ListPreparedStatementsInput, opts ...request.Option) (*athena.ListPreparedStatementsOutput, error)
	ListPreparedStatementsPagesWithContext(ctx context.Context, input *athena.ListPreparedStatementsInput, cb func(*athena.ListPreparedStatementsOutput, bool) bool, opts ...request.Option) error
	ListQueryExecutionsWithContext(ctx context.Context, input *athena.ListQueryExecutionsInput, opts ...request.Option) (*athena.ListQueryExecutionsOutput, error)
	ListQueryExecutionsPagesWithContext(ctx context.Context, input *athena.ListQueryExecutionsInput, cb func(*athena.ListQueryExecutionsOutput, bool) bool, opts ...request.Option) error
	ListTableMetadataWithContext(ctx context.Context, input *athena.ListTableMetadataInput, opts ...request.Option) (*athena.ListTableMetadataOutput, error)
	ListTableMetadataPagesWithContext(ctx context.Context, input *athena.ListTableMetadataInput, cb func(*athena.ListTableMetadataOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *athena.ListTagsForResourceInput, opts ...request.Option) (*athena.ListTagsForResourceOutput, error)
	ListTagsForResourcePagesWithContext(ctx context.Context, input *athena.ListTagsForResourceInput, cb func(*athena.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error
	ListWorkGroupsWithContext(ctx context.Context, input *athena.ListWorkGroupsInput, opts ...request.Option) (*athena.ListWorkGroupsOutput, error)
	ListWorkGroupsPagesWithContext(ctx context.Context, input *athena.ListWorkGroupsInput, cb func(*athena.ListWorkGroupsOutput, bool) bool, opts ...request.Option) error
	StartQueryExecutionWithContext(ctx context.Context, input *athena.StartQueryExecutionInput, opts ...request.Option) (*athena.StartQueryExecutionOutput, error)
	StopQueryExecutionWithContext(ctx context.Context, input *athena.StopQueryExecutionInput, opts ...request.Option) (*athena.StopQueryExecutionOutput, error)
	TagResourceWithContext(ctx context.Context, input *athena.TagResourceInput, opts ...request.Option) (*athena.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *athena.UntagResourceInput, opts ...request.Option) (*athena.UntagResourceOutput, error)
	UpdateDataCatalogWithContext(ctx context.Context, input *athena.UpdateDataCatalogInput, opts ...request.Option) (*athena.UpdateDataCatalogOutput, error)
	UpdateNamedQueryWithContext(ctx context.Context, input *athena.UpdateNamedQueryInput, opts ...request.Option) (*athena.UpdateNamedQueryOutput, error)
	UpdatePreparedStatementWithContext(ctx context.Context, input *athena.UpdatePreparedStatementInput, opts ...request.Option) (*athena.UpdatePreparedStatementOutput, error)
	UpdateWorkGroupWithContext(ctx context.Context, input *athena.UpdateWorkGroupInput, opts ...request.Option) (*athena.UpdateWorkGroupOutput, error)
}

type Client struct {
	athenaiface.AthenaAPI
	Contexter awsctx.Contexter
}

func New(base athenaiface.AthenaAPI, ctxer awsctx.Contexter) Athena {
	return &Client{
		AthenaAPI: base,
		Contexter: ctxer,
	}
}

var _ Athena = (*athena.Athena)(nil)
var _ Athena = (*Client)(nil)

func (c *Client) BatchGetNamedQueryWithContext(ctx context.Context, input *athena.BatchGetNamedQueryInput, opts ...request.Option) (*athena.BatchGetNamedQueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "BatchGetNamedQuery",
		Input:   input,
		Output:  (*athena.BatchGetNamedQueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.BatchGetNamedQueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.BatchGetNamedQueryOutput), req.Error
}

func (c *Client) BatchGetPreparedStatementWithContext(ctx context.Context, input *athena.BatchGetPreparedStatementInput, opts ...request.Option) (*athena.BatchGetPreparedStatementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "BatchGetPreparedStatement",
		Input:   input,
		Output:  (*athena.BatchGetPreparedStatementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.BatchGetPreparedStatementWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.BatchGetPreparedStatementOutput), req.Error
}

func (c *Client) BatchGetQueryExecutionWithContext(ctx context.Context, input *athena.BatchGetQueryExecutionInput, opts ...request.Option) (*athena.BatchGetQueryExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "BatchGetQueryExecution",
		Input:   input,
		Output:  (*athena.BatchGetQueryExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.BatchGetQueryExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.BatchGetQueryExecutionOutput), req.Error
}

func (c *Client) CreateDataCatalogWithContext(ctx context.Context, input *athena.CreateDataCatalogInput, opts ...request.Option) (*athena.CreateDataCatalogOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "CreateDataCatalog",
		Input:   input,
		Output:  (*athena.CreateDataCatalogOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.CreateDataCatalogWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.CreateDataCatalogOutput), req.Error
}

func (c *Client) CreateNamedQueryWithContext(ctx context.Context, input *athena.CreateNamedQueryInput, opts ...request.Option) (*athena.CreateNamedQueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "CreateNamedQuery",
		Input:   input,
		Output:  (*athena.CreateNamedQueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.CreateNamedQueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.CreateNamedQueryOutput), req.Error
}

func (c *Client) CreatePreparedStatementWithContext(ctx context.Context, input *athena.CreatePreparedStatementInput, opts ...request.Option) (*athena.CreatePreparedStatementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "CreatePreparedStatement",
		Input:   input,
		Output:  (*athena.CreatePreparedStatementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.CreatePreparedStatementWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.CreatePreparedStatementOutput), req.Error
}

func (c *Client) CreateWorkGroupWithContext(ctx context.Context, input *athena.CreateWorkGroupInput, opts ...request.Option) (*athena.CreateWorkGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "CreateWorkGroup",
		Input:   input,
		Output:  (*athena.CreateWorkGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.CreateWorkGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.CreateWorkGroupOutput), req.Error
}

func (c *Client) DeleteDataCatalogWithContext(ctx context.Context, input *athena.DeleteDataCatalogInput, opts ...request.Option) (*athena.DeleteDataCatalogOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "DeleteDataCatalog",
		Input:   input,
		Output:  (*athena.DeleteDataCatalogOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.DeleteDataCatalogWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.DeleteDataCatalogOutput), req.Error
}

func (c *Client) DeleteNamedQueryWithContext(ctx context.Context, input *athena.DeleteNamedQueryInput, opts ...request.Option) (*athena.DeleteNamedQueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "DeleteNamedQuery",
		Input:   input,
		Output:  (*athena.DeleteNamedQueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.DeleteNamedQueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.DeleteNamedQueryOutput), req.Error
}

func (c *Client) DeletePreparedStatementWithContext(ctx context.Context, input *athena.DeletePreparedStatementInput, opts ...request.Option) (*athena.DeletePreparedStatementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "DeletePreparedStatement",
		Input:   input,
		Output:  (*athena.DeletePreparedStatementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.DeletePreparedStatementWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.DeletePreparedStatementOutput), req.Error
}

func (c *Client) DeleteWorkGroupWithContext(ctx context.Context, input *athena.DeleteWorkGroupInput, opts ...request.Option) (*athena.DeleteWorkGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "DeleteWorkGroup",
		Input:   input,
		Output:  (*athena.DeleteWorkGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.DeleteWorkGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.DeleteWorkGroupOutput), req.Error
}

func (c *Client) GetDataCatalogWithContext(ctx context.Context, input *athena.GetDataCatalogInput, opts ...request.Option) (*athena.GetDataCatalogOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "GetDataCatalog",
		Input:   input,
		Output:  (*athena.GetDataCatalogOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.GetDataCatalogWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.GetDataCatalogOutput), req.Error
}

func (c *Client) GetDatabaseWithContext(ctx context.Context, input *athena.GetDatabaseInput, opts ...request.Option) (*athena.GetDatabaseOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "GetDatabase",
		Input:   input,
		Output:  (*athena.GetDatabaseOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.GetDatabaseWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.GetDatabaseOutput), req.Error
}

func (c *Client) GetNamedQueryWithContext(ctx context.Context, input *athena.GetNamedQueryInput, opts ...request.Option) (*athena.GetNamedQueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "GetNamedQuery",
		Input:   input,
		Output:  (*athena.GetNamedQueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.GetNamedQueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.GetNamedQueryOutput), req.Error
}

func (c *Client) GetPreparedStatementWithContext(ctx context.Context, input *athena.GetPreparedStatementInput, opts ...request.Option) (*athena.GetPreparedStatementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "GetPreparedStatement",
		Input:   input,
		Output:  (*athena.GetPreparedStatementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.GetPreparedStatementWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.GetPreparedStatementOutput), req.Error
}

func (c *Client) GetQueryExecutionWithContext(ctx context.Context, input *athena.GetQueryExecutionInput, opts ...request.Option) (*athena.GetQueryExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "GetQueryExecution",
		Input:   input,
		Output:  (*athena.GetQueryExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.GetQueryExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.GetQueryExecutionOutput), req.Error
}

func (c *Client) GetQueryResultsWithContext(ctx context.Context, input *athena.GetQueryResultsInput, opts ...request.Option) (*athena.GetQueryResultsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "GetQueryResults",
		Input:   input,
		Output:  (*athena.GetQueryResultsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.GetQueryResultsWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.GetQueryResultsOutput), req.Error
}

func (c *Client) GetQueryResultsPagesWithContext(ctx context.Context, input *athena.GetQueryResultsInput, cb func(*athena.GetQueryResultsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "GetQueryResults",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AthenaAPI.GetQueryResultsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetTableMetadataWithContext(ctx context.Context, input *athena.GetTableMetadataInput, opts ...request.Option) (*athena.GetTableMetadataOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "GetTableMetadata",
		Input:   input,
		Output:  (*athena.GetTableMetadataOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.GetTableMetadataWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.GetTableMetadataOutput), req.Error
}

func (c *Client) GetWorkGroupWithContext(ctx context.Context, input *athena.GetWorkGroupInput, opts ...request.Option) (*athena.GetWorkGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "GetWorkGroup",
		Input:   input,
		Output:  (*athena.GetWorkGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.GetWorkGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.GetWorkGroupOutput), req.Error
}

func (c *Client) ListDataCatalogsWithContext(ctx context.Context, input *athena.ListDataCatalogsInput, opts ...request.Option) (*athena.ListDataCatalogsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListDataCatalogs",
		Input:   input,
		Output:  (*athena.ListDataCatalogsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.ListDataCatalogsWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.ListDataCatalogsOutput), req.Error
}

func (c *Client) ListDataCatalogsPagesWithContext(ctx context.Context, input *athena.ListDataCatalogsInput, cb func(*athena.ListDataCatalogsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListDataCatalogs",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AthenaAPI.ListDataCatalogsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListDatabasesWithContext(ctx context.Context, input *athena.ListDatabasesInput, opts ...request.Option) (*athena.ListDatabasesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListDatabases",
		Input:   input,
		Output:  (*athena.ListDatabasesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.ListDatabasesWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.ListDatabasesOutput), req.Error
}

func (c *Client) ListDatabasesPagesWithContext(ctx context.Context, input *athena.ListDatabasesInput, cb func(*athena.ListDatabasesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListDatabases",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AthenaAPI.ListDatabasesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListEngineVersionsWithContext(ctx context.Context, input *athena.ListEngineVersionsInput, opts ...request.Option) (*athena.ListEngineVersionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListEngineVersions",
		Input:   input,
		Output:  (*athena.ListEngineVersionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.ListEngineVersionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.ListEngineVersionsOutput), req.Error
}

func (c *Client) ListNamedQueriesWithContext(ctx context.Context, input *athena.ListNamedQueriesInput, opts ...request.Option) (*athena.ListNamedQueriesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListNamedQueries",
		Input:   input,
		Output:  (*athena.ListNamedQueriesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.ListNamedQueriesWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.ListNamedQueriesOutput), req.Error
}

func (c *Client) ListNamedQueriesPagesWithContext(ctx context.Context, input *athena.ListNamedQueriesInput, cb func(*athena.ListNamedQueriesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListNamedQueries",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AthenaAPI.ListNamedQueriesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListPreparedStatementsWithContext(ctx context.Context, input *athena.ListPreparedStatementsInput, opts ...request.Option) (*athena.ListPreparedStatementsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListPreparedStatements",
		Input:   input,
		Output:  (*athena.ListPreparedStatementsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.ListPreparedStatementsWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.ListPreparedStatementsOutput), req.Error
}

func (c *Client) ListPreparedStatementsPagesWithContext(ctx context.Context, input *athena.ListPreparedStatementsInput, cb func(*athena.ListPreparedStatementsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListPreparedStatements",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AthenaAPI.ListPreparedStatementsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListQueryExecutionsWithContext(ctx context.Context, input *athena.ListQueryExecutionsInput, opts ...request.Option) (*athena.ListQueryExecutionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListQueryExecutions",
		Input:   input,
		Output:  (*athena.ListQueryExecutionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.ListQueryExecutionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.ListQueryExecutionsOutput), req.Error
}

func (c *Client) ListQueryExecutionsPagesWithContext(ctx context.Context, input *athena.ListQueryExecutionsInput, cb func(*athena.ListQueryExecutionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListQueryExecutions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AthenaAPI.ListQueryExecutionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTableMetadataWithContext(ctx context.Context, input *athena.ListTableMetadataInput, opts ...request.Option) (*athena.ListTableMetadataOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListTableMetadata",
		Input:   input,
		Output:  (*athena.ListTableMetadataOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.ListTableMetadataWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.ListTableMetadataOutput), req.Error
}

func (c *Client) ListTableMetadataPagesWithContext(ctx context.Context, input *athena.ListTableMetadataInput, cb func(*athena.ListTableMetadataOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListTableMetadata",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AthenaAPI.ListTableMetadataPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *athena.ListTagsForResourceInput, opts ...request.Option) (*athena.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*athena.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTagsForResourcePagesWithContext(ctx context.Context, input *athena.ListTagsForResourceInput, cb func(*athena.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListTagsForResource",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AthenaAPI.ListTagsForResourcePagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListWorkGroupsWithContext(ctx context.Context, input *athena.ListWorkGroupsInput, opts ...request.Option) (*athena.ListWorkGroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListWorkGroups",
		Input:   input,
		Output:  (*athena.ListWorkGroupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.ListWorkGroupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.ListWorkGroupsOutput), req.Error
}

func (c *Client) ListWorkGroupsPagesWithContext(ctx context.Context, input *athena.ListWorkGroupsInput, cb func(*athena.ListWorkGroupsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "ListWorkGroups",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.AthenaAPI.ListWorkGroupsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) StartQueryExecutionWithContext(ctx context.Context, input *athena.StartQueryExecutionInput, opts ...request.Option) (*athena.StartQueryExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "StartQueryExecution",
		Input:   input,
		Output:  (*athena.StartQueryExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.StartQueryExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.StartQueryExecutionOutput), req.Error
}

func (c *Client) StopQueryExecutionWithContext(ctx context.Context, input *athena.StopQueryExecutionInput, opts ...request.Option) (*athena.StopQueryExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "StopQueryExecution",
		Input:   input,
		Output:  (*athena.StopQueryExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.StopQueryExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.StopQueryExecutionOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *athena.TagResourceInput, opts ...request.Option) (*athena.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "TagResource",
		Input:   input,
		Output:  (*athena.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *athena.UntagResourceInput, opts ...request.Option) (*athena.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*athena.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.UntagResourceOutput), req.Error
}

func (c *Client) UpdateDataCatalogWithContext(ctx context.Context, input *athena.UpdateDataCatalogInput, opts ...request.Option) (*athena.UpdateDataCatalogOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "UpdateDataCatalog",
		Input:   input,
		Output:  (*athena.UpdateDataCatalogOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.UpdateDataCatalogWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.UpdateDataCatalogOutput), req.Error
}

func (c *Client) UpdateNamedQueryWithContext(ctx context.Context, input *athena.UpdateNamedQueryInput, opts ...request.Option) (*athena.UpdateNamedQueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "UpdateNamedQuery",
		Input:   input,
		Output:  (*athena.UpdateNamedQueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.UpdateNamedQueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.UpdateNamedQueryOutput), req.Error
}

func (c *Client) UpdatePreparedStatementWithContext(ctx context.Context, input *athena.UpdatePreparedStatementInput, opts ...request.Option) (*athena.UpdatePreparedStatementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "UpdatePreparedStatement",
		Input:   input,
		Output:  (*athena.UpdatePreparedStatementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.UpdatePreparedStatementWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.UpdatePreparedStatementOutput), req.Error
}

func (c *Client) UpdateWorkGroupWithContext(ctx context.Context, input *athena.UpdateWorkGroupInput, opts ...request.Option) (*athena.UpdateWorkGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "athena",
		Action:  "UpdateWorkGroup",
		Input:   input,
		Output:  (*athena.UpdateWorkGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AthenaAPI.UpdateWorkGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*athena.UpdateWorkGroupOutput), req.Error
}
