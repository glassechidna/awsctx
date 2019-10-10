// Code generated by internal/generate/main.go. DO NOT EDIT.

package kmsctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
	"github.com/glassechidna/awsctx"
)

type KMS interface {
	CancelKeyDeletionWithContext(ctx context.Context, input *kms.CancelKeyDeletionInput, opts ...request.Option) (*kms.CancelKeyDeletionOutput, error)
	ConnectCustomKeyStoreWithContext(ctx context.Context, input *kms.ConnectCustomKeyStoreInput, opts ...request.Option) (*kms.ConnectCustomKeyStoreOutput, error)
	CreateAliasWithContext(ctx context.Context, input *kms.CreateAliasInput, opts ...request.Option) (*kms.CreateAliasOutput, error)
	CreateCustomKeyStoreWithContext(ctx context.Context, input *kms.CreateCustomKeyStoreInput, opts ...request.Option) (*kms.CreateCustomKeyStoreOutput, error)
	CreateGrantWithContext(ctx context.Context, input *kms.CreateGrantInput, opts ...request.Option) (*kms.CreateGrantOutput, error)
	CreateKeyWithContext(ctx context.Context, input *kms.CreateKeyInput, opts ...request.Option) (*kms.CreateKeyOutput, error)
	DecryptWithContext(ctx context.Context, input *kms.DecryptInput, opts ...request.Option) (*kms.DecryptOutput, error)
	DeleteAliasWithContext(ctx context.Context, input *kms.DeleteAliasInput, opts ...request.Option) (*kms.DeleteAliasOutput, error)
	DeleteCustomKeyStoreWithContext(ctx context.Context, input *kms.DeleteCustomKeyStoreInput, opts ...request.Option) (*kms.DeleteCustomKeyStoreOutput, error)
	DeleteImportedKeyMaterialWithContext(ctx context.Context, input *kms.DeleteImportedKeyMaterialInput, opts ...request.Option) (*kms.DeleteImportedKeyMaterialOutput, error)
	DescribeCustomKeyStoresWithContext(ctx context.Context, input *kms.DescribeCustomKeyStoresInput, opts ...request.Option) (*kms.DescribeCustomKeyStoresOutput, error)
	DescribeKeyWithContext(ctx context.Context, input *kms.DescribeKeyInput, opts ...request.Option) (*kms.DescribeKeyOutput, error)
	DisableKeyWithContext(ctx context.Context, input *kms.DisableKeyInput, opts ...request.Option) (*kms.DisableKeyOutput, error)
	DisableKeyRotationWithContext(ctx context.Context, input *kms.DisableKeyRotationInput, opts ...request.Option) (*kms.DisableKeyRotationOutput, error)
	DisconnectCustomKeyStoreWithContext(ctx context.Context, input *kms.DisconnectCustomKeyStoreInput, opts ...request.Option) (*kms.DisconnectCustomKeyStoreOutput, error)
	EnableKeyWithContext(ctx context.Context, input *kms.EnableKeyInput, opts ...request.Option) (*kms.EnableKeyOutput, error)
	EnableKeyRotationWithContext(ctx context.Context, input *kms.EnableKeyRotationInput, opts ...request.Option) (*kms.EnableKeyRotationOutput, error)
	EncryptWithContext(ctx context.Context, input *kms.EncryptInput, opts ...request.Option) (*kms.EncryptOutput, error)
	GenerateDataKeyWithContext(ctx context.Context, input *kms.GenerateDataKeyInput, opts ...request.Option) (*kms.GenerateDataKeyOutput, error)
	GenerateDataKeyWithoutPlaintextWithContext(ctx context.Context, input *kms.GenerateDataKeyWithoutPlaintextInput, opts ...request.Option) (*kms.GenerateDataKeyWithoutPlaintextOutput, error)
	GenerateRandomWithContext(ctx context.Context, input *kms.GenerateRandomInput, opts ...request.Option) (*kms.GenerateRandomOutput, error)
	GetKeyPolicyWithContext(ctx context.Context, input *kms.GetKeyPolicyInput, opts ...request.Option) (*kms.GetKeyPolicyOutput, error)
	GetKeyRotationStatusWithContext(ctx context.Context, input *kms.GetKeyRotationStatusInput, opts ...request.Option) (*kms.GetKeyRotationStatusOutput, error)
	GetParametersForImportWithContext(ctx context.Context, input *kms.GetParametersForImportInput, opts ...request.Option) (*kms.GetParametersForImportOutput, error)
	ImportKeyMaterialWithContext(ctx context.Context, input *kms.ImportKeyMaterialInput, opts ...request.Option) (*kms.ImportKeyMaterialOutput, error)
	ListAliasesWithContext(ctx context.Context, input *kms.ListAliasesInput, opts ...request.Option) (*kms.ListAliasesOutput, error)
	ListGrantsWithContext(ctx context.Context, input *kms.ListGrantsInput, opts ...request.Option) (*kms.ListGrantsResponse, error)
	ListKeyPoliciesWithContext(ctx context.Context, input *kms.ListKeyPoliciesInput, opts ...request.Option) (*kms.ListKeyPoliciesOutput, error)
	ListKeysWithContext(ctx context.Context, input *kms.ListKeysInput, opts ...request.Option) (*kms.ListKeysOutput, error)
	ListResourceTagsWithContext(ctx context.Context, input *kms.ListResourceTagsInput, opts ...request.Option) (*kms.ListResourceTagsOutput, error)
	ListRetirableGrantsWithContext(ctx context.Context, input *kms.ListRetirableGrantsInput, opts ...request.Option) (*kms.ListGrantsResponse, error)
	PutKeyPolicyWithContext(ctx context.Context, input *kms.PutKeyPolicyInput, opts ...request.Option) (*kms.PutKeyPolicyOutput, error)
	ReEncryptWithContext(ctx context.Context, input *kms.ReEncryptInput, opts ...request.Option) (*kms.ReEncryptOutput, error)
	RetireGrantWithContext(ctx context.Context, input *kms.RetireGrantInput, opts ...request.Option) (*kms.RetireGrantOutput, error)
	RevokeGrantWithContext(ctx context.Context, input *kms.RevokeGrantInput, opts ...request.Option) (*kms.RevokeGrantOutput, error)
	ScheduleKeyDeletionWithContext(ctx context.Context, input *kms.ScheduleKeyDeletionInput, opts ...request.Option) (*kms.ScheduleKeyDeletionOutput, error)
	TagResourceWithContext(ctx context.Context, input *kms.TagResourceInput, opts ...request.Option) (*kms.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *kms.UntagResourceInput, opts ...request.Option) (*kms.UntagResourceOutput, error)
	UpdateAliasWithContext(ctx context.Context, input *kms.UpdateAliasInput, opts ...request.Option) (*kms.UpdateAliasOutput, error)
	UpdateCustomKeyStoreWithContext(ctx context.Context, input *kms.UpdateCustomKeyStoreInput, opts ...request.Option) (*kms.UpdateCustomKeyStoreOutput, error)
	UpdateKeyDescriptionWithContext(ctx context.Context, input *kms.UpdateKeyDescriptionInput, opts ...request.Option) (*kms.UpdateKeyDescriptionOutput, error)
}

type Client struct {
	kmsiface.KMSAPI
	Contexter awsctx.Contexter
}

func New(base kmsiface.KMSAPI, ctxer awsctx.Contexter) KMS {
	return &Client{
		KMSAPI: base,
		Contexter: ctxer,
	}
}

var _ KMS = (*kms.KMS)(nil)
var _ KMS = (*Client)(nil)

func (c *Client) CancelKeyDeletionWithContext(ctx context.Context, input *kms.CancelKeyDeletionInput, opts ...request.Option) (*kms.CancelKeyDeletionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "CancelKeyDeletion",
		Input:   input,
		Output:  (*kms.CancelKeyDeletionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.CancelKeyDeletionWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.CancelKeyDeletionOutput), req.Error
}

func (c *Client) ConnectCustomKeyStoreWithContext(ctx context.Context, input *kms.ConnectCustomKeyStoreInput, opts ...request.Option) (*kms.ConnectCustomKeyStoreOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ConnectCustomKeyStore",
		Input:   input,
		Output:  (*kms.ConnectCustomKeyStoreOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ConnectCustomKeyStoreWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ConnectCustomKeyStoreOutput), req.Error
}

func (c *Client) CreateAliasWithContext(ctx context.Context, input *kms.CreateAliasInput, opts ...request.Option) (*kms.CreateAliasOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "CreateAlias",
		Input:   input,
		Output:  (*kms.CreateAliasOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.CreateAliasWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.CreateAliasOutput), req.Error
}

func (c *Client) CreateCustomKeyStoreWithContext(ctx context.Context, input *kms.CreateCustomKeyStoreInput, opts ...request.Option) (*kms.CreateCustomKeyStoreOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "CreateCustomKeyStore",
		Input:   input,
		Output:  (*kms.CreateCustomKeyStoreOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.CreateCustomKeyStoreWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.CreateCustomKeyStoreOutput), req.Error
}

func (c *Client) CreateGrantWithContext(ctx context.Context, input *kms.CreateGrantInput, opts ...request.Option) (*kms.CreateGrantOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "CreateGrant",
		Input:   input,
		Output:  (*kms.CreateGrantOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.CreateGrantWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.CreateGrantOutput), req.Error
}

func (c *Client) CreateKeyWithContext(ctx context.Context, input *kms.CreateKeyInput, opts ...request.Option) (*kms.CreateKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "CreateKey",
		Input:   input,
		Output:  (*kms.CreateKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.CreateKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.CreateKeyOutput), req.Error
}

func (c *Client) DecryptWithContext(ctx context.Context, input *kms.DecryptInput, opts ...request.Option) (*kms.DecryptOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "Decrypt",
		Input:   input,
		Output:  (*kms.DecryptOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.DecryptWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.DecryptOutput), req.Error
}

func (c *Client) DeleteAliasWithContext(ctx context.Context, input *kms.DeleteAliasInput, opts ...request.Option) (*kms.DeleteAliasOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "DeleteAlias",
		Input:   input,
		Output:  (*kms.DeleteAliasOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.DeleteAliasWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.DeleteAliasOutput), req.Error
}

func (c *Client) DeleteCustomKeyStoreWithContext(ctx context.Context, input *kms.DeleteCustomKeyStoreInput, opts ...request.Option) (*kms.DeleteCustomKeyStoreOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "DeleteCustomKeyStore",
		Input:   input,
		Output:  (*kms.DeleteCustomKeyStoreOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.DeleteCustomKeyStoreWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.DeleteCustomKeyStoreOutput), req.Error
}

func (c *Client) DeleteImportedKeyMaterialWithContext(ctx context.Context, input *kms.DeleteImportedKeyMaterialInput, opts ...request.Option) (*kms.DeleteImportedKeyMaterialOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "DeleteImportedKeyMaterial",
		Input:   input,
		Output:  (*kms.DeleteImportedKeyMaterialOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.DeleteImportedKeyMaterialWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.DeleteImportedKeyMaterialOutput), req.Error
}

func (c *Client) DescribeCustomKeyStoresWithContext(ctx context.Context, input *kms.DescribeCustomKeyStoresInput, opts ...request.Option) (*kms.DescribeCustomKeyStoresOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "DescribeCustomKeyStores",
		Input:   input,
		Output:  (*kms.DescribeCustomKeyStoresOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.DescribeCustomKeyStoresWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.DescribeCustomKeyStoresOutput), req.Error
}

func (c *Client) DescribeKeyWithContext(ctx context.Context, input *kms.DescribeKeyInput, opts ...request.Option) (*kms.DescribeKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "DescribeKey",
		Input:   input,
		Output:  (*kms.DescribeKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.DescribeKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.DescribeKeyOutput), req.Error
}

func (c *Client) DisableKeyWithContext(ctx context.Context, input *kms.DisableKeyInput, opts ...request.Option) (*kms.DisableKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "DisableKey",
		Input:   input,
		Output:  (*kms.DisableKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.DisableKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.DisableKeyOutput), req.Error
}

func (c *Client) DisableKeyRotationWithContext(ctx context.Context, input *kms.DisableKeyRotationInput, opts ...request.Option) (*kms.DisableKeyRotationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "DisableKeyRotation",
		Input:   input,
		Output:  (*kms.DisableKeyRotationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.DisableKeyRotationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.DisableKeyRotationOutput), req.Error
}

func (c *Client) DisconnectCustomKeyStoreWithContext(ctx context.Context, input *kms.DisconnectCustomKeyStoreInput, opts ...request.Option) (*kms.DisconnectCustomKeyStoreOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "DisconnectCustomKeyStore",
		Input:   input,
		Output:  (*kms.DisconnectCustomKeyStoreOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.DisconnectCustomKeyStoreWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.DisconnectCustomKeyStoreOutput), req.Error
}

func (c *Client) EnableKeyWithContext(ctx context.Context, input *kms.EnableKeyInput, opts ...request.Option) (*kms.EnableKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "EnableKey",
		Input:   input,
		Output:  (*kms.EnableKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.EnableKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.EnableKeyOutput), req.Error
}

func (c *Client) EnableKeyRotationWithContext(ctx context.Context, input *kms.EnableKeyRotationInput, opts ...request.Option) (*kms.EnableKeyRotationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "EnableKeyRotation",
		Input:   input,
		Output:  (*kms.EnableKeyRotationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.EnableKeyRotationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.EnableKeyRotationOutput), req.Error
}

func (c *Client) EncryptWithContext(ctx context.Context, input *kms.EncryptInput, opts ...request.Option) (*kms.EncryptOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "Encrypt",
		Input:   input,
		Output:  (*kms.EncryptOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.EncryptWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.EncryptOutput), req.Error
}

func (c *Client) GenerateDataKeyWithContext(ctx context.Context, input *kms.GenerateDataKeyInput, opts ...request.Option) (*kms.GenerateDataKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "GenerateDataKey",
		Input:   input,
		Output:  (*kms.GenerateDataKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.GenerateDataKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.GenerateDataKeyOutput), req.Error
}

func (c *Client) GenerateDataKeyWithoutPlaintextWithContext(ctx context.Context, input *kms.GenerateDataKeyWithoutPlaintextInput, opts ...request.Option) (*kms.GenerateDataKeyWithoutPlaintextOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "GenerateDataKeyWithoutPlaintext",
		Input:   input,
		Output:  (*kms.GenerateDataKeyWithoutPlaintextOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.GenerateDataKeyWithoutPlaintextWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.GenerateDataKeyWithoutPlaintextOutput), req.Error
}

func (c *Client) GenerateRandomWithContext(ctx context.Context, input *kms.GenerateRandomInput, opts ...request.Option) (*kms.GenerateRandomOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "GenerateRandom",
		Input:   input,
		Output:  (*kms.GenerateRandomOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.GenerateRandomWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.GenerateRandomOutput), req.Error
}

func (c *Client) GetKeyPolicyWithContext(ctx context.Context, input *kms.GetKeyPolicyInput, opts ...request.Option) (*kms.GetKeyPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "GetKeyPolicy",
		Input:   input,
		Output:  (*kms.GetKeyPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.GetKeyPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.GetKeyPolicyOutput), req.Error
}

func (c *Client) GetKeyRotationStatusWithContext(ctx context.Context, input *kms.GetKeyRotationStatusInput, opts ...request.Option) (*kms.GetKeyRotationStatusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "GetKeyRotationStatus",
		Input:   input,
		Output:  (*kms.GetKeyRotationStatusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.GetKeyRotationStatusWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.GetKeyRotationStatusOutput), req.Error
}

func (c *Client) GetParametersForImportWithContext(ctx context.Context, input *kms.GetParametersForImportInput, opts ...request.Option) (*kms.GetParametersForImportOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "GetParametersForImport",
		Input:   input,
		Output:  (*kms.GetParametersForImportOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.GetParametersForImportWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.GetParametersForImportOutput), req.Error
}

func (c *Client) ImportKeyMaterialWithContext(ctx context.Context, input *kms.ImportKeyMaterialInput, opts ...request.Option) (*kms.ImportKeyMaterialOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ImportKeyMaterial",
		Input:   input,
		Output:  (*kms.ImportKeyMaterialOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ImportKeyMaterialWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ImportKeyMaterialOutput), req.Error
}

func (c *Client) ListAliasesWithContext(ctx context.Context, input *kms.ListAliasesInput, opts ...request.Option) (*kms.ListAliasesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ListAliases",
		Input:   input,
		Output:  (*kms.ListAliasesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ListAliasesWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ListAliasesOutput), req.Error
}

func (c *Client) ListGrantsWithContext(ctx context.Context, input *kms.ListGrantsInput, opts ...request.Option) (*kms.ListGrantsResponse, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ListGrants",
		Input:   input,
		Output:  (*kms.ListGrantsResponse)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ListGrantsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ListGrantsResponse), req.Error
}

func (c *Client) ListKeyPoliciesWithContext(ctx context.Context, input *kms.ListKeyPoliciesInput, opts ...request.Option) (*kms.ListKeyPoliciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ListKeyPolicies",
		Input:   input,
		Output:  (*kms.ListKeyPoliciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ListKeyPoliciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ListKeyPoliciesOutput), req.Error
}

func (c *Client) ListKeysWithContext(ctx context.Context, input *kms.ListKeysInput, opts ...request.Option) (*kms.ListKeysOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ListKeys",
		Input:   input,
		Output:  (*kms.ListKeysOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ListKeysWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ListKeysOutput), req.Error
}

func (c *Client) ListResourceTagsWithContext(ctx context.Context, input *kms.ListResourceTagsInput, opts ...request.Option) (*kms.ListResourceTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ListResourceTags",
		Input:   input,
		Output:  (*kms.ListResourceTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ListResourceTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ListResourceTagsOutput), req.Error
}

func (c *Client) ListRetirableGrantsWithContext(ctx context.Context, input *kms.ListRetirableGrantsInput, opts ...request.Option) (*kms.ListGrantsResponse, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ListRetirableGrants",
		Input:   input,
		Output:  (*kms.ListGrantsResponse)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ListRetirableGrantsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ListGrantsResponse), req.Error
}

func (c *Client) PutKeyPolicyWithContext(ctx context.Context, input *kms.PutKeyPolicyInput, opts ...request.Option) (*kms.PutKeyPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "PutKeyPolicy",
		Input:   input,
		Output:  (*kms.PutKeyPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.PutKeyPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.PutKeyPolicyOutput), req.Error
}

func (c *Client) ReEncryptWithContext(ctx context.Context, input *kms.ReEncryptInput, opts ...request.Option) (*kms.ReEncryptOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ReEncrypt",
		Input:   input,
		Output:  (*kms.ReEncryptOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ReEncryptWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ReEncryptOutput), req.Error
}

func (c *Client) RetireGrantWithContext(ctx context.Context, input *kms.RetireGrantInput, opts ...request.Option) (*kms.RetireGrantOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "RetireGrant",
		Input:   input,
		Output:  (*kms.RetireGrantOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.RetireGrantWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.RetireGrantOutput), req.Error
}

func (c *Client) RevokeGrantWithContext(ctx context.Context, input *kms.RevokeGrantInput, opts ...request.Option) (*kms.RevokeGrantOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "RevokeGrant",
		Input:   input,
		Output:  (*kms.RevokeGrantOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.RevokeGrantWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.RevokeGrantOutput), req.Error
}

func (c *Client) ScheduleKeyDeletionWithContext(ctx context.Context, input *kms.ScheduleKeyDeletionInput, opts ...request.Option) (*kms.ScheduleKeyDeletionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "ScheduleKeyDeletion",
		Input:   input,
		Output:  (*kms.ScheduleKeyDeletionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.ScheduleKeyDeletionWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.ScheduleKeyDeletionOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *kms.TagResourceInput, opts ...request.Option) (*kms.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "TagResource",
		Input:   input,
		Output:  (*kms.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *kms.UntagResourceInput, opts ...request.Option) (*kms.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*kms.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.UntagResourceOutput), req.Error
}

func (c *Client) UpdateAliasWithContext(ctx context.Context, input *kms.UpdateAliasInput, opts ...request.Option) (*kms.UpdateAliasOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "UpdateAlias",
		Input:   input,
		Output:  (*kms.UpdateAliasOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.UpdateAliasWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.UpdateAliasOutput), req.Error
}

func (c *Client) UpdateCustomKeyStoreWithContext(ctx context.Context, input *kms.UpdateCustomKeyStoreInput, opts ...request.Option) (*kms.UpdateCustomKeyStoreOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "UpdateCustomKeyStore",
		Input:   input,
		Output:  (*kms.UpdateCustomKeyStoreOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.UpdateCustomKeyStoreWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.UpdateCustomKeyStoreOutput), req.Error
}

func (c *Client) UpdateKeyDescriptionWithContext(ctx context.Context, input *kms.UpdateKeyDescriptionInput, opts ...request.Option) (*kms.UpdateKeyDescriptionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kms",
		Action:  "UpdateKeyDescription",
		Input:   input,
		Output:  (*kms.UpdateKeyDescriptionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KMSAPI.UpdateKeyDescriptionWithContext(ctx, input, opts...)
	})

	return req.Output.(*kms.UpdateKeyDescriptionOutput), req.Error
}
