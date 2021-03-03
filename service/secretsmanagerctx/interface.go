// Code generated by internal/generate/main.go. DO NOT EDIT.

package secretsmanagerctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/glassechidna/awsctx"
)

type SecretsManager interface {
	CancelRotateSecretWithContext(ctx context.Context, input *secretsmanager.CancelRotateSecretInput, opts ...request.Option) (*secretsmanager.CancelRotateSecretOutput, error)
	CreateSecretWithContext(ctx context.Context, input *secretsmanager.CreateSecretInput, opts ...request.Option) (*secretsmanager.CreateSecretOutput, error)
	DeleteResourcePolicyWithContext(ctx context.Context, input *secretsmanager.DeleteResourcePolicyInput, opts ...request.Option) (*secretsmanager.DeleteResourcePolicyOutput, error)
	DeleteSecretWithContext(ctx context.Context, input *secretsmanager.DeleteSecretInput, opts ...request.Option) (*secretsmanager.DeleteSecretOutput, error)
	DescribeSecretWithContext(ctx context.Context, input *secretsmanager.DescribeSecretInput, opts ...request.Option) (*secretsmanager.DescribeSecretOutput, error)
	GetRandomPasswordWithContext(ctx context.Context, input *secretsmanager.GetRandomPasswordInput, opts ...request.Option) (*secretsmanager.GetRandomPasswordOutput, error)
	GetResourcePolicyWithContext(ctx context.Context, input *secretsmanager.GetResourcePolicyInput, opts ...request.Option) (*secretsmanager.GetResourcePolicyOutput, error)
	GetSecretValueWithContext(ctx context.Context, input *secretsmanager.GetSecretValueInput, opts ...request.Option) (*secretsmanager.GetSecretValueOutput, error)
	ListSecretVersionIdsWithContext(ctx context.Context, input *secretsmanager.ListSecretVersionIdsInput, opts ...request.Option) (*secretsmanager.ListSecretVersionIdsOutput, error)
	ListSecretVersionIdsPagesWithContext(ctx context.Context, input *secretsmanager.ListSecretVersionIdsInput, cb func(*secretsmanager.ListSecretVersionIdsOutput, bool) bool, opts ...request.Option) error
	ListSecretsWithContext(ctx context.Context, input *secretsmanager.ListSecretsInput, opts ...request.Option) (*secretsmanager.ListSecretsOutput, error)
	ListSecretsPagesWithContext(ctx context.Context, input *secretsmanager.ListSecretsInput, cb func(*secretsmanager.ListSecretsOutput, bool) bool, opts ...request.Option) error
	PutResourcePolicyWithContext(ctx context.Context, input *secretsmanager.PutResourcePolicyInput, opts ...request.Option) (*secretsmanager.PutResourcePolicyOutput, error)
	PutSecretValueWithContext(ctx context.Context, input *secretsmanager.PutSecretValueInput, opts ...request.Option) (*secretsmanager.PutSecretValueOutput, error)
	RemoveRegionsFromReplicationWithContext(ctx context.Context, input *secretsmanager.RemoveRegionsFromReplicationInput, opts ...request.Option) (*secretsmanager.RemoveRegionsFromReplicationOutput, error)
	ReplicateSecretToRegionsWithContext(ctx context.Context, input *secretsmanager.ReplicateSecretToRegionsInput, opts ...request.Option) (*secretsmanager.ReplicateSecretToRegionsOutput, error)
	RestoreSecretWithContext(ctx context.Context, input *secretsmanager.RestoreSecretInput, opts ...request.Option) (*secretsmanager.RestoreSecretOutput, error)
	RotateSecretWithContext(ctx context.Context, input *secretsmanager.RotateSecretInput, opts ...request.Option) (*secretsmanager.RotateSecretOutput, error)
	StopReplicationToReplicaWithContext(ctx context.Context, input *secretsmanager.StopReplicationToReplicaInput, opts ...request.Option) (*secretsmanager.StopReplicationToReplicaOutput, error)
	TagResourceWithContext(ctx context.Context, input *secretsmanager.TagResourceInput, opts ...request.Option) (*secretsmanager.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *secretsmanager.UntagResourceInput, opts ...request.Option) (*secretsmanager.UntagResourceOutput, error)
	UpdateSecretWithContext(ctx context.Context, input *secretsmanager.UpdateSecretInput, opts ...request.Option) (*secretsmanager.UpdateSecretOutput, error)
	UpdateSecretVersionStageWithContext(ctx context.Context, input *secretsmanager.UpdateSecretVersionStageInput, opts ...request.Option) (*secretsmanager.UpdateSecretVersionStageOutput, error)
	ValidateResourcePolicyWithContext(ctx context.Context, input *secretsmanager.ValidateResourcePolicyInput, opts ...request.Option) (*secretsmanager.ValidateResourcePolicyOutput, error)
}

type Client struct {
	secretsmanageriface.SecretsManagerAPI
	Contexter awsctx.Contexter
}

func New(base secretsmanageriface.SecretsManagerAPI, ctxer awsctx.Contexter) SecretsManager {
	return &Client{
		SecretsManagerAPI: base,
		Contexter: ctxer,
	}
}

var _ SecretsManager = (*secretsmanager.SecretsManager)(nil)
var _ SecretsManager = (*Client)(nil)

func (c *Client) CancelRotateSecretWithContext(ctx context.Context, input *secretsmanager.CancelRotateSecretInput, opts ...request.Option) (*secretsmanager.CancelRotateSecretOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "CancelRotateSecret",
		Input:   input,
		Output:  (*secretsmanager.CancelRotateSecretOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.CancelRotateSecretWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.CancelRotateSecretOutput), req.Error
}

func (c *Client) CreateSecretWithContext(ctx context.Context, input *secretsmanager.CreateSecretInput, opts ...request.Option) (*secretsmanager.CreateSecretOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "CreateSecret",
		Input:   input,
		Output:  (*secretsmanager.CreateSecretOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.CreateSecretWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.CreateSecretOutput), req.Error
}

func (c *Client) DeleteResourcePolicyWithContext(ctx context.Context, input *secretsmanager.DeleteResourcePolicyInput, opts ...request.Option) (*secretsmanager.DeleteResourcePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "DeleteResourcePolicy",
		Input:   input,
		Output:  (*secretsmanager.DeleteResourcePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.DeleteResourcePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.DeleteResourcePolicyOutput), req.Error
}

func (c *Client) DeleteSecretWithContext(ctx context.Context, input *secretsmanager.DeleteSecretInput, opts ...request.Option) (*secretsmanager.DeleteSecretOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "DeleteSecret",
		Input:   input,
		Output:  (*secretsmanager.DeleteSecretOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.DeleteSecretWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.DeleteSecretOutput), req.Error
}

func (c *Client) DescribeSecretWithContext(ctx context.Context, input *secretsmanager.DescribeSecretInput, opts ...request.Option) (*secretsmanager.DescribeSecretOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "DescribeSecret",
		Input:   input,
		Output:  (*secretsmanager.DescribeSecretOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.DescribeSecretWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.DescribeSecretOutput), req.Error
}

func (c *Client) GetRandomPasswordWithContext(ctx context.Context, input *secretsmanager.GetRandomPasswordInput, opts ...request.Option) (*secretsmanager.GetRandomPasswordOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "GetRandomPassword",
		Input:   input,
		Output:  (*secretsmanager.GetRandomPasswordOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.GetRandomPasswordWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.GetRandomPasswordOutput), req.Error
}

func (c *Client) GetResourcePolicyWithContext(ctx context.Context, input *secretsmanager.GetResourcePolicyInput, opts ...request.Option) (*secretsmanager.GetResourcePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "GetResourcePolicy",
		Input:   input,
		Output:  (*secretsmanager.GetResourcePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.GetResourcePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.GetResourcePolicyOutput), req.Error
}

func (c *Client) GetSecretValueWithContext(ctx context.Context, input *secretsmanager.GetSecretValueInput, opts ...request.Option) (*secretsmanager.GetSecretValueOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "GetSecretValue",
		Input:   input,
		Output:  (*secretsmanager.GetSecretValueOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.GetSecretValueWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.GetSecretValueOutput), req.Error
}

func (c *Client) ListSecretVersionIdsWithContext(ctx context.Context, input *secretsmanager.ListSecretVersionIdsInput, opts ...request.Option) (*secretsmanager.ListSecretVersionIdsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "ListSecretVersionIds",
		Input:   input,
		Output:  (*secretsmanager.ListSecretVersionIdsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.ListSecretVersionIdsWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.ListSecretVersionIdsOutput), req.Error
}

func (c *Client) ListSecretVersionIdsPagesWithContext(ctx context.Context, input *secretsmanager.ListSecretVersionIdsInput, cb func(*secretsmanager.ListSecretVersionIdsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "ListSecretVersionIds",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SecretsManagerAPI.ListSecretVersionIdsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListSecretsWithContext(ctx context.Context, input *secretsmanager.ListSecretsInput, opts ...request.Option) (*secretsmanager.ListSecretsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "ListSecrets",
		Input:   input,
		Output:  (*secretsmanager.ListSecretsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.ListSecretsWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.ListSecretsOutput), req.Error
}

func (c *Client) ListSecretsPagesWithContext(ctx context.Context, input *secretsmanager.ListSecretsInput, cb func(*secretsmanager.ListSecretsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "ListSecrets",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SecretsManagerAPI.ListSecretsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) PutResourcePolicyWithContext(ctx context.Context, input *secretsmanager.PutResourcePolicyInput, opts ...request.Option) (*secretsmanager.PutResourcePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "PutResourcePolicy",
		Input:   input,
		Output:  (*secretsmanager.PutResourcePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.PutResourcePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.PutResourcePolicyOutput), req.Error
}

func (c *Client) PutSecretValueWithContext(ctx context.Context, input *secretsmanager.PutSecretValueInput, opts ...request.Option) (*secretsmanager.PutSecretValueOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "PutSecretValue",
		Input:   input,
		Output:  (*secretsmanager.PutSecretValueOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.PutSecretValueWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.PutSecretValueOutput), req.Error
}

func (c *Client) RemoveRegionsFromReplicationWithContext(ctx context.Context, input *secretsmanager.RemoveRegionsFromReplicationInput, opts ...request.Option) (*secretsmanager.RemoveRegionsFromReplicationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "RemoveRegionsFromReplication",
		Input:   input,
		Output:  (*secretsmanager.RemoveRegionsFromReplicationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.RemoveRegionsFromReplicationWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.RemoveRegionsFromReplicationOutput), req.Error
}

func (c *Client) ReplicateSecretToRegionsWithContext(ctx context.Context, input *secretsmanager.ReplicateSecretToRegionsInput, opts ...request.Option) (*secretsmanager.ReplicateSecretToRegionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "ReplicateSecretToRegions",
		Input:   input,
		Output:  (*secretsmanager.ReplicateSecretToRegionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.ReplicateSecretToRegionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.ReplicateSecretToRegionsOutput), req.Error
}

func (c *Client) RestoreSecretWithContext(ctx context.Context, input *secretsmanager.RestoreSecretInput, opts ...request.Option) (*secretsmanager.RestoreSecretOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "RestoreSecret",
		Input:   input,
		Output:  (*secretsmanager.RestoreSecretOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.RestoreSecretWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.RestoreSecretOutput), req.Error
}

func (c *Client) RotateSecretWithContext(ctx context.Context, input *secretsmanager.RotateSecretInput, opts ...request.Option) (*secretsmanager.RotateSecretOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "RotateSecret",
		Input:   input,
		Output:  (*secretsmanager.RotateSecretOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.RotateSecretWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.RotateSecretOutput), req.Error
}

func (c *Client) StopReplicationToReplicaWithContext(ctx context.Context, input *secretsmanager.StopReplicationToReplicaInput, opts ...request.Option) (*secretsmanager.StopReplicationToReplicaOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "StopReplicationToReplica",
		Input:   input,
		Output:  (*secretsmanager.StopReplicationToReplicaOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.StopReplicationToReplicaWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.StopReplicationToReplicaOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *secretsmanager.TagResourceInput, opts ...request.Option) (*secretsmanager.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "TagResource",
		Input:   input,
		Output:  (*secretsmanager.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *secretsmanager.UntagResourceInput, opts ...request.Option) (*secretsmanager.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*secretsmanager.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.UntagResourceOutput), req.Error
}

func (c *Client) UpdateSecretWithContext(ctx context.Context, input *secretsmanager.UpdateSecretInput, opts ...request.Option) (*secretsmanager.UpdateSecretOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "UpdateSecret",
		Input:   input,
		Output:  (*secretsmanager.UpdateSecretOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.UpdateSecretWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.UpdateSecretOutput), req.Error
}

func (c *Client) UpdateSecretVersionStageWithContext(ctx context.Context, input *secretsmanager.UpdateSecretVersionStageInput, opts ...request.Option) (*secretsmanager.UpdateSecretVersionStageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "UpdateSecretVersionStage",
		Input:   input,
		Output:  (*secretsmanager.UpdateSecretVersionStageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.UpdateSecretVersionStageWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.UpdateSecretVersionStageOutput), req.Error
}

func (c *Client) ValidateResourcePolicyWithContext(ctx context.Context, input *secretsmanager.ValidateResourcePolicyInput, opts ...request.Option) (*secretsmanager.ValidateResourcePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "secretsmanager",
		Action:  "ValidateResourcePolicy",
		Input:   input,
		Output:  (*secretsmanager.ValidateResourcePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SecretsManagerAPI.ValidateResourcePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*secretsmanager.ValidateResourcePolicyOutput), req.Error
}
