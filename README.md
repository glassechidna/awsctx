# awsctx

## What

The AWS SDK for Go ([aws-sdk-go][sdk]) has [supported][ctx-support] the context pattern since 2017.
This facilitates support for cancellation of long-running requests, etc. This library aims to rectify
what I personally see as some shortcomings of the official library.

**Contexts shouldn't be optional**. As [blogged][ctx-mandatory] by AWS, they are making context-passing
mandatory in more situations in the SDK v2 beta. This is because dropping contexts is a common source
of tricky-to-diagnose bugs. AWS can't break backwards compatibility by *removing* the context-less
SDK methods in v1, but there's no way to *enforce* usage of the contextual methods. This library fixes
that by creating interfaces that *only* expose the contextual methods, i.e. a limited subset of the
AWS service client interfaces.

**Inability to easily hook all methods in a high level way**. The AWS service clients have support for
adding _request handlers_ in a generic fashion, but these are invoked later in the request lifecycle
than I would like. They only have access to serialised request objects, whereas I want to be able to
modify the context object before a request is invoked and have access to metadata about requests. This
helps facilitate application-wide observability.

## How

We'll use S3 for these examples, but they apply equally well to any other services supported by the SDK.
Typical usage of the SDK involves passing around either the client struct directly, or objects implementing
the interface equivalent. In code:

```go
package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/glassechidna/awsctx/service/s3ctx"
)

func main() {
	sess := session.Must(session.NewSession())
	api := s3.New(sess)
	SomeFuncThatUsesClientStruct(api) // good
	SomeFuncThatUsesInterface(api)    // better  
	SomeFuncThatUsesCtxInterface(api) // best (in my opinion!)
}

// this is how basic programs typically get the job done. it's quick and it works.
func SomeFuncThatUsesClientStruct(api *s3.S3) {
	_, _ = api.ListObjectsWithContext(context.Background(), &s3.ListObjectsInput{})
	// ...
}

// using the iface interface instead has a few advantages. the biggest is that it
// facilitates much easier testability, because a "mock" s3 client can be passed
// in for tests, so that network calls to the real s3 are not needed for unit tests.
func SomeFuncThatUsesInterface(api s3iface.S3API) {
	_, _ = api.ListObjectsWithContext(context.Background(), &s3.ListObjectsInput{})
	// ...
}

// s3ctx.S3 is a much smaller interface that s3iface.S3API as it has 88 methods 
// instead of 282. it's a strict subset of the aws-provided interface, so it 
// should be a drop-in replacement - assuming your code currently only uses the 
// *WithContext methods.
func SomeFuncThatUsesCtxInterface(api s3ctx.S3) {
	_, _ = api.ListObjectsWithContext(context.Background(), &s3.ListObjectsInput{})
	// ...
}
```

## Advanced

So the above is fine for ensuring that only contextual methods are fine. But maybe you want to hook
all methods as well. To achieve that, a `Contexter` interface is used. Like `http.Handler`, it has a
`ContexterFunc` wrapper for simple cases. Here's how you use it:

```go
package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/glassechidna/awsctx"
	"github.com/glassechidna/awsctx/service/stsctx"
	"time"
)

func main() {
	sess := session.Must(session.NewSession())
	base := sts.New(sess)
		
	ctxer := awsctx.ContexterFunc(func(ctx context.Context, request *awsctx.AwsRequest, inner func(ctx context.Context)) {
	    defer recordApiCall(ctx, request, time.Now())
	    inner(ctx)
	})
		
	ctxual := stsctx.New(base, ctxer)
	_, _ = ctxual.GetCallerIdentityWithContext(context.Background(), &sts.GetCallerIdentityInput{})
}

func recordApiCall(ctx context.Context, request *awsctx.AwsRequest, start time.Time) {
	// you could even do a switch on the type of request.Input and enrich your logs with bucket names, etc.

	seconds := time.Now().Sub(start).Seconds()
	fmt.Printf("api=%s action=%s took %f seconds\n", request.Service, request.Action, seconds)
	
	// maybe you will also want to send this to honeycomb.io, new relic, other apm, etc.
	// these kinds of services usually want the context object as it allows for correlation
	// with other events - you can do that here!
}
```

## Outstanding issues

**Automation**. To be maximally useful, I need to set up CI to automatically publish a tag in
this repo every time the upstream AWS SDK has a new version published. Ideally the version numbers
should be in sync with AWS to avoid hair being torn out.

**Pagination and waiters**. I haven't yet implemented support for the `*PagesWithContext` methods. Those 
are useful and should be added. Likewise for the `WaitUntil***WithContext` methods.

[sdk]: https://github.com/aws/aws-sdk-go
[ctx-support]: https://aws.amazon.com/blogs/developer/context-pattern-added-to-the-aws-sdk-for-go/
[ctx-mandatory]: https://aws.amazon.com/blogs/developer/v2-aws-sdk-for-go-adds-context-to-api-operations/
