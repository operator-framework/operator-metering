// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilBucketExists uses the Amazon S3 API operation
// HeadBucket to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *S3) WaitUntilBucketExists(input *HeadBucketInput) error {
	return c.WaitUntilBucketExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilBucketExistsWithContext is an extended version of WaitUntilBucketExists.
// With the support for passing in a context and options to con***REMOVED***gure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *S3) WaitUntilBucketExistsWithContext(ctx aws.Context, input *HeadBucketInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilBucketExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 301,
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 403,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 404,
			},
		},
		Logger: c.Con***REMOVED***g.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *HeadBucketInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.HeadBucketRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilBucketNotExists uses the Amazon S3 API operation
// HeadBucket to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *S3) WaitUntilBucketNotExists(input *HeadBucketInput) error {
	return c.WaitUntilBucketNotExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilBucketNotExistsWithContext is an extended version of WaitUntilBucketNotExists.
// With the support for passing in a context and options to con***REMOVED***gure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *S3) WaitUntilBucketNotExistsWithContext(ctx aws.Context, input *HeadBucketInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilBucketNotExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 404,
			},
		},
		Logger: c.Con***REMOVED***g.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *HeadBucketInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.HeadBucketRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilObjectExists uses the Amazon S3 API operation
// HeadObject to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *S3) WaitUntilObjectExists(input *HeadObjectInput) error {
	return c.WaitUntilObjectExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilObjectExistsWithContext is an extended version of WaitUntilObjectExists.
// With the support for passing in a context and options to con***REMOVED***gure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *S3) WaitUntilObjectExistsWithContext(ctx aws.Context, input *HeadObjectInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilObjectExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 404,
			},
		},
		Logger: c.Con***REMOVED***g.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *HeadObjectInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.HeadObjectRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilObjectNotExists uses the Amazon S3 API operation
// HeadObject to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *S3) WaitUntilObjectNotExists(input *HeadObjectInput) error {
	return c.WaitUntilObjectNotExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilObjectNotExistsWithContext is an extended version of WaitUntilObjectNotExists.
// With the support for passing in a context and options to con***REMOVED***gure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *S3) WaitUntilObjectNotExistsWithContext(ctx aws.Context, input *HeadObjectInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilObjectNotExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 404,
			},
		},
		Logger: c.Con***REMOVED***g.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *HeadObjectInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.HeadObjectRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
