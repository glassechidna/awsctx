// Code generated by internal/generate/main.go. DO NOT EDIT.

package personalizeruntimectx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"github.com/aws/aws-sdk-go/service/personalizeruntime/personalizeruntimeiface"
	"github.com/glassechidna/awsctx"
)

type PersonalizeRuntime interface {
	GetPersonalizedRankingWithContext(ctx context.Context, input *personalizeruntime.GetPersonalizedRankingInput, opts ...request.Option) (*personalizeruntime.GetPersonalizedRankingOutput, error)
	GetRecommendationsWithContext(ctx context.Context, input *personalizeruntime.GetRecommendationsInput, opts ...request.Option) (*personalizeruntime.GetRecommendationsOutput, error)
}

type Client struct {
	personalizeruntimeiface.PersonalizeRuntimeAPI
	Contexter awsctx.Contexter
}

func New(base personalizeruntimeiface.PersonalizeRuntimeAPI, ctxer awsctx.Contexter) PersonalizeRuntime {
	return &Client{
		PersonalizeRuntimeAPI: base,
		Contexter: ctxer,
	}
}

var _ PersonalizeRuntime = (*personalizeruntime.PersonalizeRuntime)(nil)
var _ PersonalizeRuntime = (*Client)(nil)

func (c *Client) GetPersonalizedRankingWithContext(ctx context.Context, input *personalizeruntime.GetPersonalizedRankingInput, opts ...request.Option) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "personalizeruntime",
		Action:  "GetPersonalizedRankingWithContext",
		Input:   input,
		Output:  (*personalizeruntime.GetPersonalizedRankingOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PersonalizeRuntimeAPI.GetPersonalizedRankingWithContext(ctx, input, opts...)
	})

	return req.Output.(*personalizeruntime.GetPersonalizedRankingOutput), req.Error
}

func (c *Client) GetRecommendationsWithContext(ctx context.Context, input *personalizeruntime.GetRecommendationsInput, opts ...request.Option) (*personalizeruntime.GetRecommendationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "personalizeruntime",
		Action:  "GetRecommendationsWithContext",
		Input:   input,
		Output:  (*personalizeruntime.GetRecommendationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PersonalizeRuntimeAPI.GetRecommendationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*personalizeruntime.GetRecommendationsOutput), req.Error
}
