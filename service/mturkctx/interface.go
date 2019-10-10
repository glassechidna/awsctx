// Code generated by internal/generate/main.go. DO NOT EDIT.

package mturkctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aws/aws-sdk-go/service/mturk/mturkiface"
	"github.com/glassechidna/awsctx"
)

type MTurk interface {
	AcceptQualificationRequestWithContext(ctx context.Context, input *mturk.AcceptQualificationRequestInput, opts ...request.Option) (*mturk.AcceptQualificationRequestOutput, error)
	ApproveAssignmentWithContext(ctx context.Context, input *mturk.ApproveAssignmentInput, opts ...request.Option) (*mturk.ApproveAssignmentOutput, error)
	AssociateQualificationWithWorkerWithContext(ctx context.Context, input *mturk.AssociateQualificationWithWorkerInput, opts ...request.Option) (*mturk.AssociateQualificationWithWorkerOutput, error)
	CreateAdditionalAssignmentsForHITWithContext(ctx context.Context, input *mturk.CreateAdditionalAssignmentsForHITInput, opts ...request.Option) (*mturk.CreateAdditionalAssignmentsForHITOutput, error)
	CreateHITWithContext(ctx context.Context, input *mturk.CreateHITInput, opts ...request.Option) (*mturk.CreateHITOutput, error)
	CreateHITTypeWithContext(ctx context.Context, input *mturk.CreateHITTypeInput, opts ...request.Option) (*mturk.CreateHITTypeOutput, error)
	CreateHITWithHITTypeWithContext(ctx context.Context, input *mturk.CreateHITWithHITTypeInput, opts ...request.Option) (*mturk.CreateHITWithHITTypeOutput, error)
	CreateQualificationTypeWithContext(ctx context.Context, input *mturk.CreateQualificationTypeInput, opts ...request.Option) (*mturk.CreateQualificationTypeOutput, error)
	CreateWorkerBlockWithContext(ctx context.Context, input *mturk.CreateWorkerBlockInput, opts ...request.Option) (*mturk.CreateWorkerBlockOutput, error)
	DeleteHITWithContext(ctx context.Context, input *mturk.DeleteHITInput, opts ...request.Option) (*mturk.DeleteHITOutput, error)
	DeleteQualificationTypeWithContext(ctx context.Context, input *mturk.DeleteQualificationTypeInput, opts ...request.Option) (*mturk.DeleteQualificationTypeOutput, error)
	DeleteWorkerBlockWithContext(ctx context.Context, input *mturk.DeleteWorkerBlockInput, opts ...request.Option) (*mturk.DeleteWorkerBlockOutput, error)
	DisassociateQualificationFromWorkerWithContext(ctx context.Context, input *mturk.DisassociateQualificationFromWorkerInput, opts ...request.Option) (*mturk.DisassociateQualificationFromWorkerOutput, error)
	GetAccountBalanceWithContext(ctx context.Context, input *mturk.GetAccountBalanceInput, opts ...request.Option) (*mturk.GetAccountBalanceOutput, error)
	GetAssignmentWithContext(ctx context.Context, input *mturk.GetAssignmentInput, opts ...request.Option) (*mturk.GetAssignmentOutput, error)
	GetFileUploadURLWithContext(ctx context.Context, input *mturk.GetFileUploadURLInput, opts ...request.Option) (*mturk.GetFileUploadURLOutput, error)
	GetHITWithContext(ctx context.Context, input *mturk.GetHITInput, opts ...request.Option) (*mturk.GetHITOutput, error)
	GetQualificationScoreWithContext(ctx context.Context, input *mturk.GetQualificationScoreInput, opts ...request.Option) (*mturk.GetQualificationScoreOutput, error)
	GetQualificationTypeWithContext(ctx context.Context, input *mturk.GetQualificationTypeInput, opts ...request.Option) (*mturk.GetQualificationTypeOutput, error)
	ListAssignmentsForHITWithContext(ctx context.Context, input *mturk.ListAssignmentsForHITInput, opts ...request.Option) (*mturk.ListAssignmentsForHITOutput, error)
	ListBonusPaymentsWithContext(ctx context.Context, input *mturk.ListBonusPaymentsInput, opts ...request.Option) (*mturk.ListBonusPaymentsOutput, error)
	ListHITsWithContext(ctx context.Context, input *mturk.ListHITsInput, opts ...request.Option) (*mturk.ListHITsOutput, error)
	ListHITsForQualificationTypeWithContext(ctx context.Context, input *mturk.ListHITsForQualificationTypeInput, opts ...request.Option) (*mturk.ListHITsForQualificationTypeOutput, error)
	ListQualificationRequestsWithContext(ctx context.Context, input *mturk.ListQualificationRequestsInput, opts ...request.Option) (*mturk.ListQualificationRequestsOutput, error)
	ListQualificationTypesWithContext(ctx context.Context, input *mturk.ListQualificationTypesInput, opts ...request.Option) (*mturk.ListQualificationTypesOutput, error)
	ListReviewPolicyResultsForHITWithContext(ctx context.Context, input *mturk.ListReviewPolicyResultsForHITInput, opts ...request.Option) (*mturk.ListReviewPolicyResultsForHITOutput, error)
	ListReviewableHITsWithContext(ctx context.Context, input *mturk.ListReviewableHITsInput, opts ...request.Option) (*mturk.ListReviewableHITsOutput, error)
	ListWorkerBlocksWithContext(ctx context.Context, input *mturk.ListWorkerBlocksInput, opts ...request.Option) (*mturk.ListWorkerBlocksOutput, error)
	ListWorkersWithQualificationTypeWithContext(ctx context.Context, input *mturk.ListWorkersWithQualificationTypeInput, opts ...request.Option) (*mturk.ListWorkersWithQualificationTypeOutput, error)
	NotifyWorkersWithContext(ctx context.Context, input *mturk.NotifyWorkersInput, opts ...request.Option) (*mturk.NotifyWorkersOutput, error)
	RejectAssignmentWithContext(ctx context.Context, input *mturk.RejectAssignmentInput, opts ...request.Option) (*mturk.RejectAssignmentOutput, error)
	RejectQualificationRequestWithContext(ctx context.Context, input *mturk.RejectQualificationRequestInput, opts ...request.Option) (*mturk.RejectQualificationRequestOutput, error)
	SendBonusWithContext(ctx context.Context, input *mturk.SendBonusInput, opts ...request.Option) (*mturk.SendBonusOutput, error)
	SendTestEventNotificationWithContext(ctx context.Context, input *mturk.SendTestEventNotificationInput, opts ...request.Option) (*mturk.SendTestEventNotificationOutput, error)
	UpdateExpirationForHITWithContext(ctx context.Context, input *mturk.UpdateExpirationForHITInput, opts ...request.Option) (*mturk.UpdateExpirationForHITOutput, error)
	UpdateHITReviewStatusWithContext(ctx context.Context, input *mturk.UpdateHITReviewStatusInput, opts ...request.Option) (*mturk.UpdateHITReviewStatusOutput, error)
	UpdateHITTypeOfHITWithContext(ctx context.Context, input *mturk.UpdateHITTypeOfHITInput, opts ...request.Option) (*mturk.UpdateHITTypeOfHITOutput, error)
	UpdateNotificationSettingsWithContext(ctx context.Context, input *mturk.UpdateNotificationSettingsInput, opts ...request.Option) (*mturk.UpdateNotificationSettingsOutput, error)
	UpdateQualificationTypeWithContext(ctx context.Context, input *mturk.UpdateQualificationTypeInput, opts ...request.Option) (*mturk.UpdateQualificationTypeOutput, error)
}

type Client struct {
	mturkiface.MTurkAPI
	Contexter awsctx.Contexter
}

func New(base mturkiface.MTurkAPI, ctxer awsctx.Contexter) MTurk {
	return &Client{
		MTurkAPI: base,
		Contexter: ctxer,
	}
}

var _ MTurk = (*mturk.MTurk)(nil)
var _ MTurk = (*Client)(nil)

func (c *Client) AcceptQualificationRequestWithContext(ctx context.Context, input *mturk.AcceptQualificationRequestInput, opts ...request.Option) (*mturk.AcceptQualificationRequestOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "AcceptQualificationRequest",
		Input:   input,
		Output:  (*mturk.AcceptQualificationRequestOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.AcceptQualificationRequestWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.AcceptQualificationRequestOutput), req.Error
}

func (c *Client) ApproveAssignmentWithContext(ctx context.Context, input *mturk.ApproveAssignmentInput, opts ...request.Option) (*mturk.ApproveAssignmentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ApproveAssignment",
		Input:   input,
		Output:  (*mturk.ApproveAssignmentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ApproveAssignmentWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ApproveAssignmentOutput), req.Error
}

func (c *Client) AssociateQualificationWithWorkerWithContext(ctx context.Context, input *mturk.AssociateQualificationWithWorkerInput, opts ...request.Option) (*mturk.AssociateQualificationWithWorkerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "AssociateQualificationWithWorker",
		Input:   input,
		Output:  (*mturk.AssociateQualificationWithWorkerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.AssociateQualificationWithWorkerWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.AssociateQualificationWithWorkerOutput), req.Error
}

func (c *Client) CreateAdditionalAssignmentsForHITWithContext(ctx context.Context, input *mturk.CreateAdditionalAssignmentsForHITInput, opts ...request.Option) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "CreateAdditionalAssignmentsForHIT",
		Input:   input,
		Output:  (*mturk.CreateAdditionalAssignmentsForHITOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.CreateAdditionalAssignmentsForHITWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.CreateAdditionalAssignmentsForHITOutput), req.Error
}

func (c *Client) CreateHITWithContext(ctx context.Context, input *mturk.CreateHITInput, opts ...request.Option) (*mturk.CreateHITOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "CreateHIT",
		Input:   input,
		Output:  (*mturk.CreateHITOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.CreateHITWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.CreateHITOutput), req.Error
}

func (c *Client) CreateHITTypeWithContext(ctx context.Context, input *mturk.CreateHITTypeInput, opts ...request.Option) (*mturk.CreateHITTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "CreateHITType",
		Input:   input,
		Output:  (*mturk.CreateHITTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.CreateHITTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.CreateHITTypeOutput), req.Error
}

func (c *Client) CreateHITWithHITTypeWithContext(ctx context.Context, input *mturk.CreateHITWithHITTypeInput, opts ...request.Option) (*mturk.CreateHITWithHITTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "CreateHITWithHITType",
		Input:   input,
		Output:  (*mturk.CreateHITWithHITTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.CreateHITWithHITTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.CreateHITWithHITTypeOutput), req.Error
}

func (c *Client) CreateQualificationTypeWithContext(ctx context.Context, input *mturk.CreateQualificationTypeInput, opts ...request.Option) (*mturk.CreateQualificationTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "CreateQualificationType",
		Input:   input,
		Output:  (*mturk.CreateQualificationTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.CreateQualificationTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.CreateQualificationTypeOutput), req.Error
}

func (c *Client) CreateWorkerBlockWithContext(ctx context.Context, input *mturk.CreateWorkerBlockInput, opts ...request.Option) (*mturk.CreateWorkerBlockOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "CreateWorkerBlock",
		Input:   input,
		Output:  (*mturk.CreateWorkerBlockOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.CreateWorkerBlockWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.CreateWorkerBlockOutput), req.Error
}

func (c *Client) DeleteHITWithContext(ctx context.Context, input *mturk.DeleteHITInput, opts ...request.Option) (*mturk.DeleteHITOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "DeleteHIT",
		Input:   input,
		Output:  (*mturk.DeleteHITOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.DeleteHITWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.DeleteHITOutput), req.Error
}

func (c *Client) DeleteQualificationTypeWithContext(ctx context.Context, input *mturk.DeleteQualificationTypeInput, opts ...request.Option) (*mturk.DeleteQualificationTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "DeleteQualificationType",
		Input:   input,
		Output:  (*mturk.DeleteQualificationTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.DeleteQualificationTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.DeleteQualificationTypeOutput), req.Error
}

func (c *Client) DeleteWorkerBlockWithContext(ctx context.Context, input *mturk.DeleteWorkerBlockInput, opts ...request.Option) (*mturk.DeleteWorkerBlockOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "DeleteWorkerBlock",
		Input:   input,
		Output:  (*mturk.DeleteWorkerBlockOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.DeleteWorkerBlockWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.DeleteWorkerBlockOutput), req.Error
}

func (c *Client) DisassociateQualificationFromWorkerWithContext(ctx context.Context, input *mturk.DisassociateQualificationFromWorkerInput, opts ...request.Option) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "DisassociateQualificationFromWorker",
		Input:   input,
		Output:  (*mturk.DisassociateQualificationFromWorkerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.DisassociateQualificationFromWorkerWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.DisassociateQualificationFromWorkerOutput), req.Error
}

func (c *Client) GetAccountBalanceWithContext(ctx context.Context, input *mturk.GetAccountBalanceInput, opts ...request.Option) (*mturk.GetAccountBalanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "GetAccountBalance",
		Input:   input,
		Output:  (*mturk.GetAccountBalanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.GetAccountBalanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.GetAccountBalanceOutput), req.Error
}

func (c *Client) GetAssignmentWithContext(ctx context.Context, input *mturk.GetAssignmentInput, opts ...request.Option) (*mturk.GetAssignmentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "GetAssignment",
		Input:   input,
		Output:  (*mturk.GetAssignmentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.GetAssignmentWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.GetAssignmentOutput), req.Error
}

func (c *Client) GetFileUploadURLWithContext(ctx context.Context, input *mturk.GetFileUploadURLInput, opts ...request.Option) (*mturk.GetFileUploadURLOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "GetFileUploadURL",
		Input:   input,
		Output:  (*mturk.GetFileUploadURLOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.GetFileUploadURLWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.GetFileUploadURLOutput), req.Error
}

func (c *Client) GetHITWithContext(ctx context.Context, input *mturk.GetHITInput, opts ...request.Option) (*mturk.GetHITOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "GetHIT",
		Input:   input,
		Output:  (*mturk.GetHITOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.GetHITWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.GetHITOutput), req.Error
}

func (c *Client) GetQualificationScoreWithContext(ctx context.Context, input *mturk.GetQualificationScoreInput, opts ...request.Option) (*mturk.GetQualificationScoreOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "GetQualificationScore",
		Input:   input,
		Output:  (*mturk.GetQualificationScoreOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.GetQualificationScoreWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.GetQualificationScoreOutput), req.Error
}

func (c *Client) GetQualificationTypeWithContext(ctx context.Context, input *mturk.GetQualificationTypeInput, opts ...request.Option) (*mturk.GetQualificationTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "GetQualificationType",
		Input:   input,
		Output:  (*mturk.GetQualificationTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.GetQualificationTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.GetQualificationTypeOutput), req.Error
}

func (c *Client) ListAssignmentsForHITWithContext(ctx context.Context, input *mturk.ListAssignmentsForHITInput, opts ...request.Option) (*mturk.ListAssignmentsForHITOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListAssignmentsForHIT",
		Input:   input,
		Output:  (*mturk.ListAssignmentsForHITOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListAssignmentsForHITWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListAssignmentsForHITOutput), req.Error
}

func (c *Client) ListBonusPaymentsWithContext(ctx context.Context, input *mturk.ListBonusPaymentsInput, opts ...request.Option) (*mturk.ListBonusPaymentsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListBonusPayments",
		Input:   input,
		Output:  (*mturk.ListBonusPaymentsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListBonusPaymentsWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListBonusPaymentsOutput), req.Error
}

func (c *Client) ListHITsWithContext(ctx context.Context, input *mturk.ListHITsInput, opts ...request.Option) (*mturk.ListHITsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListHITs",
		Input:   input,
		Output:  (*mturk.ListHITsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListHITsWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListHITsOutput), req.Error
}

func (c *Client) ListHITsForQualificationTypeWithContext(ctx context.Context, input *mturk.ListHITsForQualificationTypeInput, opts ...request.Option) (*mturk.ListHITsForQualificationTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListHITsForQualificationType",
		Input:   input,
		Output:  (*mturk.ListHITsForQualificationTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListHITsForQualificationTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListHITsForQualificationTypeOutput), req.Error
}

func (c *Client) ListQualificationRequestsWithContext(ctx context.Context, input *mturk.ListQualificationRequestsInput, opts ...request.Option) (*mturk.ListQualificationRequestsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListQualificationRequests",
		Input:   input,
		Output:  (*mturk.ListQualificationRequestsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListQualificationRequestsWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListQualificationRequestsOutput), req.Error
}

func (c *Client) ListQualificationTypesWithContext(ctx context.Context, input *mturk.ListQualificationTypesInput, opts ...request.Option) (*mturk.ListQualificationTypesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListQualificationTypes",
		Input:   input,
		Output:  (*mturk.ListQualificationTypesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListQualificationTypesWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListQualificationTypesOutput), req.Error
}

func (c *Client) ListReviewPolicyResultsForHITWithContext(ctx context.Context, input *mturk.ListReviewPolicyResultsForHITInput, opts ...request.Option) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListReviewPolicyResultsForHIT",
		Input:   input,
		Output:  (*mturk.ListReviewPolicyResultsForHITOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListReviewPolicyResultsForHITWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListReviewPolicyResultsForHITOutput), req.Error
}

func (c *Client) ListReviewableHITsWithContext(ctx context.Context, input *mturk.ListReviewableHITsInput, opts ...request.Option) (*mturk.ListReviewableHITsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListReviewableHITs",
		Input:   input,
		Output:  (*mturk.ListReviewableHITsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListReviewableHITsWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListReviewableHITsOutput), req.Error
}

func (c *Client) ListWorkerBlocksWithContext(ctx context.Context, input *mturk.ListWorkerBlocksInput, opts ...request.Option) (*mturk.ListWorkerBlocksOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListWorkerBlocks",
		Input:   input,
		Output:  (*mturk.ListWorkerBlocksOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListWorkerBlocksWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListWorkerBlocksOutput), req.Error
}

func (c *Client) ListWorkersWithQualificationTypeWithContext(ctx context.Context, input *mturk.ListWorkersWithQualificationTypeInput, opts ...request.Option) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "ListWorkersWithQualificationType",
		Input:   input,
		Output:  (*mturk.ListWorkersWithQualificationTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.ListWorkersWithQualificationTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.ListWorkersWithQualificationTypeOutput), req.Error
}

func (c *Client) NotifyWorkersWithContext(ctx context.Context, input *mturk.NotifyWorkersInput, opts ...request.Option) (*mturk.NotifyWorkersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "NotifyWorkers",
		Input:   input,
		Output:  (*mturk.NotifyWorkersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.NotifyWorkersWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.NotifyWorkersOutput), req.Error
}

func (c *Client) RejectAssignmentWithContext(ctx context.Context, input *mturk.RejectAssignmentInput, opts ...request.Option) (*mturk.RejectAssignmentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "RejectAssignment",
		Input:   input,
		Output:  (*mturk.RejectAssignmentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.RejectAssignmentWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.RejectAssignmentOutput), req.Error
}

func (c *Client) RejectQualificationRequestWithContext(ctx context.Context, input *mturk.RejectQualificationRequestInput, opts ...request.Option) (*mturk.RejectQualificationRequestOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "RejectQualificationRequest",
		Input:   input,
		Output:  (*mturk.RejectQualificationRequestOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.RejectQualificationRequestWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.RejectQualificationRequestOutput), req.Error
}

func (c *Client) SendBonusWithContext(ctx context.Context, input *mturk.SendBonusInput, opts ...request.Option) (*mturk.SendBonusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "SendBonus",
		Input:   input,
		Output:  (*mturk.SendBonusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.SendBonusWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.SendBonusOutput), req.Error
}

func (c *Client) SendTestEventNotificationWithContext(ctx context.Context, input *mturk.SendTestEventNotificationInput, opts ...request.Option) (*mturk.SendTestEventNotificationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "SendTestEventNotification",
		Input:   input,
		Output:  (*mturk.SendTestEventNotificationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.SendTestEventNotificationWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.SendTestEventNotificationOutput), req.Error
}

func (c *Client) UpdateExpirationForHITWithContext(ctx context.Context, input *mturk.UpdateExpirationForHITInput, opts ...request.Option) (*mturk.UpdateExpirationForHITOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "UpdateExpirationForHIT",
		Input:   input,
		Output:  (*mturk.UpdateExpirationForHITOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.UpdateExpirationForHITWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.UpdateExpirationForHITOutput), req.Error
}

func (c *Client) UpdateHITReviewStatusWithContext(ctx context.Context, input *mturk.UpdateHITReviewStatusInput, opts ...request.Option) (*mturk.UpdateHITReviewStatusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "UpdateHITReviewStatus",
		Input:   input,
		Output:  (*mturk.UpdateHITReviewStatusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.UpdateHITReviewStatusWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.UpdateHITReviewStatusOutput), req.Error
}

func (c *Client) UpdateHITTypeOfHITWithContext(ctx context.Context, input *mturk.UpdateHITTypeOfHITInput, opts ...request.Option) (*mturk.UpdateHITTypeOfHITOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "UpdateHITTypeOfHIT",
		Input:   input,
		Output:  (*mturk.UpdateHITTypeOfHITOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.UpdateHITTypeOfHITWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.UpdateHITTypeOfHITOutput), req.Error
}

func (c *Client) UpdateNotificationSettingsWithContext(ctx context.Context, input *mturk.UpdateNotificationSettingsInput, opts ...request.Option) (*mturk.UpdateNotificationSettingsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "UpdateNotificationSettings",
		Input:   input,
		Output:  (*mturk.UpdateNotificationSettingsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.UpdateNotificationSettingsWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.UpdateNotificationSettingsOutput), req.Error
}

func (c *Client) UpdateQualificationTypeWithContext(ctx context.Context, input *mturk.UpdateQualificationTypeInput, opts ...request.Option) (*mturk.UpdateQualificationTypeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mturk",
		Action:  "UpdateQualificationType",
		Input:   input,
		Output:  (*mturk.UpdateQualificationTypeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MTurkAPI.UpdateQualificationTypeWithContext(ctx, input, opts...)
	})

	return req.Output.(*mturk.UpdateQualificationTypeOutput), req.Error
}
