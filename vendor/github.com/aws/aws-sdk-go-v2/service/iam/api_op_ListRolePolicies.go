// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the names of the inline policies that are embedded in the specified IAM
// role.
//
// An IAM role can also have managed policies attached to it. To list the managed
// policies that are attached to a role, use [ListAttachedRolePolicies]. For more information about
// policies, see [Managed policies and inline policies]in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. If there
// are no inline policies embedded with the specified role, the operation returns
// an empty list.
//
// [ListAttachedRolePolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListAttachedRolePolicies.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func (c *Client) ListRolePolicies(ctx context.Context, params *ListRolePoliciesInput, optFns ...func(*Options)) (*ListRolePoliciesOutput, error) {
	if params == nil {
		params = &ListRolePoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRolePolicies", params, optFns, c.addOperationListRolePoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRolePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRolePoliciesInput struct {

	// The name of the role to list policies for.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	RoleName *string

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	MaxItems *int32

	noSmithyDocumentSerde
}

// Contains the response to a successful [ListRolePolicies] request.
//
// [ListRolePolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListRolePolicies.html
type ListRolePoliciesOutput struct {

	// A list of policy names.
	//
	// This member is required.
	PolicyNames []string

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRolePoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListRolePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListRolePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRolePolicies"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpListRolePoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRolePolicies(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// ListRolePoliciesPaginatorOptions is the paginator options for ListRolePolicies
type ListRolePoliciesPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRolePoliciesPaginator is a paginator for ListRolePolicies
type ListRolePoliciesPaginator struct {
	options   ListRolePoliciesPaginatorOptions
	client    ListRolePoliciesAPIClient
	params    *ListRolePoliciesInput
	nextToken *string
	firstPage bool
}

// NewListRolePoliciesPaginator returns a new ListRolePoliciesPaginator
func NewListRolePoliciesPaginator(client ListRolePoliciesAPIClient, params *ListRolePoliciesInput, optFns ...func(*ListRolePoliciesPaginatorOptions)) *ListRolePoliciesPaginator {
	if params == nil {
		params = &ListRolePoliciesInput{}
	}

	options := ListRolePoliciesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRolePoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRolePoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRolePolicies page.
func (p *ListRolePoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRolePoliciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListRolePolicies(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListRolePoliciesAPIClient is a client that implements the ListRolePolicies
// operation.
type ListRolePoliciesAPIClient interface {
	ListRolePolicies(context.Context, *ListRolePoliciesInput, ...func(*Options)) (*ListRolePoliciesOutput, error)
}

var _ ListRolePoliciesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRolePolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRolePolicies",
	}
}
