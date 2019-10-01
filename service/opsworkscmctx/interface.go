// Code generated by internal/generate/main.go. DO NOT EDIT.

package opsworkscmctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"github.com/aws/aws-sdk-go/service/opsworkscm/opsworkscmiface"
	"github.com/glassechidna/awsctx"
)

type OpsWorksCM interface {
	AssociateNodeWithContext(ctx context.Context, input *opsworkscm.AssociateNodeInput, opts ...request.Option) (*opsworkscm.AssociateNodeOutput, error)
	CreateBackupWithContext(ctx context.Context, input *opsworkscm.CreateBackupInput, opts ...request.Option) (*opsworkscm.CreateBackupOutput, error)
	CreateServerWithContext(ctx context.Context, input *opsworkscm.CreateServerInput, opts ...request.Option) (*opsworkscm.CreateServerOutput, error)
	DeleteBackupWithContext(ctx context.Context, input *opsworkscm.DeleteBackupInput, opts ...request.Option) (*opsworkscm.DeleteBackupOutput, error)
	DeleteServerWithContext(ctx context.Context, input *opsworkscm.DeleteServerInput, opts ...request.Option) (*opsworkscm.DeleteServerOutput, error)
	DescribeAccountAttributesWithContext(ctx context.Context, input *opsworkscm.DescribeAccountAttributesInput, opts ...request.Option) (*opsworkscm.DescribeAccountAttributesOutput, error)
	DescribeBackupsWithContext(ctx context.Context, input *opsworkscm.DescribeBackupsInput, opts ...request.Option) (*opsworkscm.DescribeBackupsOutput, error)
	DescribeEventsWithContext(ctx context.Context, input *opsworkscm.DescribeEventsInput, opts ...request.Option) (*opsworkscm.DescribeEventsOutput, error)
	DescribeNodeAssociationStatusWithContext(ctx context.Context, input *opsworkscm.DescribeNodeAssociationStatusInput, opts ...request.Option) (*opsworkscm.DescribeNodeAssociationStatusOutput, error)
	DescribeServersWithContext(ctx context.Context, input *opsworkscm.DescribeServersInput, opts ...request.Option) (*opsworkscm.DescribeServersOutput, error)
	DisassociateNodeWithContext(ctx context.Context, input *opsworkscm.DisassociateNodeInput, opts ...request.Option) (*opsworkscm.DisassociateNodeOutput, error)
	ExportServerEngineAttributeWithContext(ctx context.Context, input *opsworkscm.ExportServerEngineAttributeInput, opts ...request.Option) (*opsworkscm.ExportServerEngineAttributeOutput, error)
	RestoreServerWithContext(ctx context.Context, input *opsworkscm.RestoreServerInput, opts ...request.Option) (*opsworkscm.RestoreServerOutput, error)
	StartMaintenanceWithContext(ctx context.Context, input *opsworkscm.StartMaintenanceInput, opts ...request.Option) (*opsworkscm.StartMaintenanceOutput, error)
	UpdateServerWithContext(ctx context.Context, input *opsworkscm.UpdateServerInput, opts ...request.Option) (*opsworkscm.UpdateServerOutput, error)
	UpdateServerEngineAttributesWithContext(ctx context.Context, input *opsworkscm.UpdateServerEngineAttributesInput, opts ...request.Option) (*opsworkscm.UpdateServerEngineAttributesOutput, error)
}

type Client struct {
	opsworkscmiface.OpsWorksCMAPI
	Contexter awsctx.Contexter
}

func New(base opsworkscmiface.OpsWorksCMAPI, ctxer awsctx.Contexter) OpsWorksCM {
	return &Client{
		OpsWorksCMAPI: base,
		Contexter: ctxer,
	}
}

var _ OpsWorksCM = (*opsworkscm.OpsWorksCM)(nil)
var _ OpsWorksCM = (*Client)(nil)

func (c *Client) AssociateNodeWithContext(ctx context.Context, input *opsworkscm.AssociateNodeInput, opts ...request.Option) (*opsworkscm.AssociateNodeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "AssociateNodeWithContext",
		Input:   input,
		Output:  (*opsworkscm.AssociateNodeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.AssociateNodeWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.AssociateNodeOutput), req.Error
}

func (c *Client) CreateBackupWithContext(ctx context.Context, input *opsworkscm.CreateBackupInput, opts ...request.Option) (*opsworkscm.CreateBackupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "CreateBackupWithContext",
		Input:   input,
		Output:  (*opsworkscm.CreateBackupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.CreateBackupWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.CreateBackupOutput), req.Error
}

func (c *Client) CreateServerWithContext(ctx context.Context, input *opsworkscm.CreateServerInput, opts ...request.Option) (*opsworkscm.CreateServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "CreateServerWithContext",
		Input:   input,
		Output:  (*opsworkscm.CreateServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.CreateServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.CreateServerOutput), req.Error
}

func (c *Client) DeleteBackupWithContext(ctx context.Context, input *opsworkscm.DeleteBackupInput, opts ...request.Option) (*opsworkscm.DeleteBackupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "DeleteBackupWithContext",
		Input:   input,
		Output:  (*opsworkscm.DeleteBackupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.DeleteBackupWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.DeleteBackupOutput), req.Error
}

func (c *Client) DeleteServerWithContext(ctx context.Context, input *opsworkscm.DeleteServerInput, opts ...request.Option) (*opsworkscm.DeleteServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "DeleteServerWithContext",
		Input:   input,
		Output:  (*opsworkscm.DeleteServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.DeleteServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.DeleteServerOutput), req.Error
}

func (c *Client) DescribeAccountAttributesWithContext(ctx context.Context, input *opsworkscm.DescribeAccountAttributesInput, opts ...request.Option) (*opsworkscm.DescribeAccountAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "DescribeAccountAttributesWithContext",
		Input:   input,
		Output:  (*opsworkscm.DescribeAccountAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.DescribeAccountAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.DescribeAccountAttributesOutput), req.Error
}

func (c *Client) DescribeBackupsWithContext(ctx context.Context, input *opsworkscm.DescribeBackupsInput, opts ...request.Option) (*opsworkscm.DescribeBackupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "DescribeBackupsWithContext",
		Input:   input,
		Output:  (*opsworkscm.DescribeBackupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.DescribeBackupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.DescribeBackupsOutput), req.Error
}

func (c *Client) DescribeEventsWithContext(ctx context.Context, input *opsworkscm.DescribeEventsInput, opts ...request.Option) (*opsworkscm.DescribeEventsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "DescribeEventsWithContext",
		Input:   input,
		Output:  (*opsworkscm.DescribeEventsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.DescribeEventsWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.DescribeEventsOutput), req.Error
}

func (c *Client) DescribeNodeAssociationStatusWithContext(ctx context.Context, input *opsworkscm.DescribeNodeAssociationStatusInput, opts ...request.Option) (*opsworkscm.DescribeNodeAssociationStatusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "DescribeNodeAssociationStatusWithContext",
		Input:   input,
		Output:  (*opsworkscm.DescribeNodeAssociationStatusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.DescribeNodeAssociationStatusWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.DescribeNodeAssociationStatusOutput), req.Error
}

func (c *Client) DescribeServersWithContext(ctx context.Context, input *opsworkscm.DescribeServersInput, opts ...request.Option) (*opsworkscm.DescribeServersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "DescribeServersWithContext",
		Input:   input,
		Output:  (*opsworkscm.DescribeServersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.DescribeServersWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.DescribeServersOutput), req.Error
}

func (c *Client) DisassociateNodeWithContext(ctx context.Context, input *opsworkscm.DisassociateNodeInput, opts ...request.Option) (*opsworkscm.DisassociateNodeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "DisassociateNodeWithContext",
		Input:   input,
		Output:  (*opsworkscm.DisassociateNodeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.DisassociateNodeWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.DisassociateNodeOutput), req.Error
}

func (c *Client) ExportServerEngineAttributeWithContext(ctx context.Context, input *opsworkscm.ExportServerEngineAttributeInput, opts ...request.Option) (*opsworkscm.ExportServerEngineAttributeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "ExportServerEngineAttributeWithContext",
		Input:   input,
		Output:  (*opsworkscm.ExportServerEngineAttributeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.ExportServerEngineAttributeWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.ExportServerEngineAttributeOutput), req.Error
}

func (c *Client) RestoreServerWithContext(ctx context.Context, input *opsworkscm.RestoreServerInput, opts ...request.Option) (*opsworkscm.RestoreServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "RestoreServerWithContext",
		Input:   input,
		Output:  (*opsworkscm.RestoreServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.RestoreServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.RestoreServerOutput), req.Error
}

func (c *Client) StartMaintenanceWithContext(ctx context.Context, input *opsworkscm.StartMaintenanceInput, opts ...request.Option) (*opsworkscm.StartMaintenanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "StartMaintenanceWithContext",
		Input:   input,
		Output:  (*opsworkscm.StartMaintenanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.StartMaintenanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.StartMaintenanceOutput), req.Error
}

func (c *Client) UpdateServerWithContext(ctx context.Context, input *opsworkscm.UpdateServerInput, opts ...request.Option) (*opsworkscm.UpdateServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "UpdateServerWithContext",
		Input:   input,
		Output:  (*opsworkscm.UpdateServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.UpdateServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.UpdateServerOutput), req.Error
}

func (c *Client) UpdateServerEngineAttributesWithContext(ctx context.Context, input *opsworkscm.UpdateServerEngineAttributesInput, opts ...request.Option) (*opsworkscm.UpdateServerEngineAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "opsworkscm",
		Action:  "UpdateServerEngineAttributesWithContext",
		Input:   input,
		Output:  (*opsworkscm.UpdateServerEngineAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.OpsWorksCMAPI.UpdateServerEngineAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*opsworkscm.UpdateServerEngineAttributesOutput), req.Error
}
