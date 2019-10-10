// Code generated by internal/generate/main.go. DO NOT EDIT.

package ec2instanceconnectctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect/ec2instanceconnectiface"
	"github.com/glassechidna/awsctx"
)

type EC2InstanceConnect interface {
	SendSSHPublicKeyWithContext(ctx context.Context, input *ec2instanceconnect.SendSSHPublicKeyInput, opts ...request.Option) (*ec2instanceconnect.SendSSHPublicKeyOutput, error)
}

type Client struct {
	ec2instanceconnectiface.EC2InstanceConnectAPI
	Contexter awsctx.Contexter
}

func New(base ec2instanceconnectiface.EC2InstanceConnectAPI, ctxer awsctx.Contexter) EC2InstanceConnect {
	return &Client{
		EC2InstanceConnectAPI: base,
		Contexter: ctxer,
	}
}

var _ EC2InstanceConnect = (*ec2instanceconnect.EC2InstanceConnect)(nil)
var _ EC2InstanceConnect = (*Client)(nil)

func (c *Client) SendSSHPublicKeyWithContext(ctx context.Context, input *ec2instanceconnect.SendSSHPublicKeyInput, opts ...request.Option) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ec2instanceconnect",
		Action:  "SendSSHPublicKey",
		Input:   input,
		Output:  (*ec2instanceconnect.SendSSHPublicKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EC2InstanceConnectAPI.SendSSHPublicKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*ec2instanceconnect.SendSSHPublicKeyOutput), req.Error
}
