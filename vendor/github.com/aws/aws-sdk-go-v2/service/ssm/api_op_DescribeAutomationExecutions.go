// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides details about all active and terminated Automation executions.
func (c *Client) DescribeAutomationExecutions(ctx context.Context, params *DescribeAutomationExecutionsInput, optFns ...func(*Options)) (*DescribeAutomationExecutionsOutput, error) {
	if params == nil {
		params = &DescribeAutomationExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutomationExecutions", params, optFns, c.addOperationDescribeAutomationExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutomationExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutomationExecutionsInput struct {

	// Filters used to limit the scope of executions that are requested.
	Filters []types.AutomationExecutionFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAutomationExecutionsOutput struct {

	// The list of details about each automation execution which has occurred which
	// matches the filter specification, if any.
	AutomationExecutionMetadataList []types.AutomationExecutionMetadata

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAutomationExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAutomationExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAutomationExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAutomationExecutions"); err != nil {
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
	if err = addOpDescribeAutomationExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutomationExecutions(options.Region), middleware.Before); err != nil {
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

// DescribeAutomationExecutionsPaginatorOptions is the paginator options for
// DescribeAutomationExecutions
type DescribeAutomationExecutionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAutomationExecutionsPaginator is a paginator for
// DescribeAutomationExecutions
type DescribeAutomationExecutionsPaginator struct {
	options   DescribeAutomationExecutionsPaginatorOptions
	client    DescribeAutomationExecutionsAPIClient
	params    *DescribeAutomationExecutionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAutomationExecutionsPaginator returns a new
// DescribeAutomationExecutionsPaginator
func NewDescribeAutomationExecutionsPaginator(client DescribeAutomationExecutionsAPIClient, params *DescribeAutomationExecutionsInput, optFns ...func(*DescribeAutomationExecutionsPaginatorOptions)) *DescribeAutomationExecutionsPaginator {
	if params == nil {
		params = &DescribeAutomationExecutionsInput{}
	}

	options := DescribeAutomationExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAutomationExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAutomationExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAutomationExecutions page.
func (p *DescribeAutomationExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAutomationExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeAutomationExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeAutomationExecutionsAPIClient is a client that implements the
// DescribeAutomationExecutions operation.
type DescribeAutomationExecutionsAPIClient interface {
	DescribeAutomationExecutions(context.Context, *DescribeAutomationExecutionsInput, ...func(*Options)) (*DescribeAutomationExecutionsOutput, error)
}

var _ DescribeAutomationExecutionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeAutomationExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAutomationExecutions",
	}
}