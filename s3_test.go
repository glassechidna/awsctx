package awsctx_test

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/glassechidna/awsctx"
	"github.com/glassechidna/awsctx/service/s3ctx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockS3 struct {
	mock.Mock
	s3iface.S3API
}

func (m *mockS3) ListObjects(input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	f := m.Called(input)
	out, _ := f.Get(0).(*s3.ListObjectsOutput)
	return out, f.Error(1)
}

func (m *mockS3) ListObjectsWithContext(ctx aws.Context, input *s3.ListObjectsInput, opts ...request.Option) (*s3.ListObjectsOutput, error) {
	f := m.Called(ctx, input, opts)
	out, _ := f.Get(0).(*s3.ListObjectsOutput)
	return out, f.Error(1)
}

func (m *mockS3) ListObjectsPagesWithContext(ctx aws.Context, input *s3.ListObjectsInput, cb func(*s3.ListObjectsOutput, bool) bool, opts ...request.Option) error {
	f := m.Called(ctx, input, cb, opts)
	return f.Error(0)
}

func TestPassthroughWithoutContext(t *testing.T) {
	api := &mockS3{}
	api.On("ListObjects", &s3.ListObjectsInput{Bucket: aws.String("bucket")}).Return(&s3.ListObjectsOutput{}, nil)

	wrapped := &s3ctx.Client{
		S3API: api,
		Contexter: awsctx.ContexterFunc(func(ctx context.Context, request *awsctx.AwsRequest, inner func(ctx context.Context)) {
			assert.Fail(t, "shouldn't call contexter for non-context methods")
		}),
	}

	out, err := wrapped.ListObjects(&s3.ListObjectsInput{Bucket: aws.String("bucket")})
	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestPaginatedWithContexter(t *testing.T) {
	api := &mockS3{}
	api.
		On("ListObjectsPagesWithContext", mock.Anything, &s3.ListObjectsInput{Bucket: aws.String("bucket")}, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			ctx := args.Get(0).(context.Context)
			val := ctx.Value("mykey").(string)
			assert.Equal(t, "hello world", val)
		}).
		Return(nil)

	wrapped := &s3ctx.Client{
		S3API: api,
		Contexter: awsctx.ContexterFunc(func(ctx context.Context, request *awsctx.AwsRequest, inner func(ctx context.Context)) {
			inner(context.WithValue(ctx, "mykey", "hello world"))
		}),
	}

	err := wrapped.ListObjectsPagesWithContext(context.Background(), &s3.ListObjectsInput{Bucket: aws.String("bucket")}, func(page *s3.ListObjectsOutput, lastPage bool) bool {
		return false
	})
	assert.NoError(t, err)
}

func TestWithContexter(t *testing.T) {
	api := &mockS3{}
	api.
		On("ListObjectsWithContext", mock.Anything, &s3.ListObjectsInput{Bucket: aws.String("bucket")}, mock.Anything).
		Run(func(args mock.Arguments) {
			ctx := args.Get(0).(context.Context)
			val := ctx.Value("mykey").(string)
			assert.Equal(t, "hello world", val)
		}).
		Return(&s3.ListObjectsOutput{}, nil)

	wrapped := &s3ctx.Client{
		S3API: api,
		Contexter: awsctx.ContexterFunc(func(ctx context.Context, request *awsctx.AwsRequest, inner func(ctx context.Context)) {
			assert.Equal(t, "s3", request.Service)
			assert.Equal(t, "ListObjects", request.Action)
			assert.NotNil(t, request.Input)
			assert.Nil(t, request.Output)
			inner(context.WithValue(ctx, "mykey", "hello world"))
			assert.NotNil(t, request.Output)
		}),
	}

	out, err := wrapped.ListObjectsWithContext(context.Background(), &s3.ListObjectsInput{Bucket: aws.String("bucket")})
	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestInnerApiReturnsError(t *testing.T) {
	api := &mockS3{}
	api.On("ListObjectsWithContext", mock.Anything, &s3.ListObjectsInput{Bucket: aws.String("bucket")}, mock.Anything).Return(nil, errors.New("some error"))

	wrapped := &s3ctx.Client{S3API: api}

	out, err := wrapped.ListObjectsWithContext(context.Background(), &s3.ListObjectsInput{Bucket: aws.String("bucket")})
	assert.Error(t, err)
	assert.Nil(t, out)
}

func TestContexterNotInvokingInnerDoesntCrash(t *testing.T) {
	api := &mockS3{}

	wrapped := &s3ctx.Client{
		S3API:     api,
		Contexter: awsctx.ContexterFunc(func(ctx context.Context, request *awsctx.AwsRequest, inner func(ctx context.Context)) {}),
	}

	assert.NotPanics(t, func() {
		out, err := wrapped.ListObjectsWithContext(context.Background(), &s3.ListObjectsInput{Bucket: aws.String("bucket")})
		assert.NoError(t, err)
		assert.Nil(t, out)
	})
}
