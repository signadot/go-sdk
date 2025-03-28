// Code generated by go-swagger; DO NOT EDIT.

package test_executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewQueryTestExecutionsParams creates a new QueryTestExecutionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryTestExecutionsParams() *QueryTestExecutionsParams {
	return &QueryTestExecutionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryTestExecutionsParamsWithTimeout creates a new QueryTestExecutionsParams object
// with the ability to set a timeout on a request.
func NewQueryTestExecutionsParamsWithTimeout(timeout time.Duration) *QueryTestExecutionsParams {
	return &QueryTestExecutionsParams{
		timeout: timeout,
	}
}

// NewQueryTestExecutionsParamsWithContext creates a new QueryTestExecutionsParams object
// with the ability to set a context for a request.
func NewQueryTestExecutionsParamsWithContext(ctx context.Context) *QueryTestExecutionsParams {
	return &QueryTestExecutionsParams{
		Context: ctx,
	}
}

// NewQueryTestExecutionsParamsWithHTTPClient creates a new QueryTestExecutionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryTestExecutionsParamsWithHTTPClient(client *http.Client) *QueryTestExecutionsParams {
	return &QueryTestExecutionsParams{
		HTTPClient: client,
	}
}

/*
QueryTestExecutionsParams contains all the parameters to send to the API endpoint

	for the query test executions operation.

	Typically these are written to a http.Request.
*/
type QueryTestExecutionsParams struct {

	/* Cluster.

	   Cluster name
	*/
	Cluster *string

	/* Cursor.

	   Start sending results from this cursor
	*/
	Cursor *string

	/* ExecutionPhase.

	   Test execution phase
	*/
	ExecutionPhase *string

	/* From.

	   Filter results starting from this time
	*/
	From *string

	/* IsAutoDiff.

	   Only include auto diff executions
	*/
	IsAutoDiff *string

	/* Label.

	   Label constraint key:value (may be specified multiple times)
	*/
	Label []string

	/* OrderDir.

	   Sorting direction
	*/
	OrderDir *string

	/* OrgName.

	   Signadot Org name
	*/
	OrgName string

	/* PageSize.

	   Number of rows to be included in the response
	*/
	PageSize *int64

	/* Published.

	   Whether the test is published to the UI by default
	*/
	Published *bool

	/* Repo.

	   Repository name
	*/
	Repo *string

	/* RepoBranch.

	   Repository path
	*/
	RepoBranch *string

	/* RepoCommitSHA.

	   Repository commit SHA
	*/
	RepoCommitSHA *string

	/* RepoPath.

	   Repository path
	*/
	RepoPath *string

	/* RunID.

	   CLI run identifier
	*/
	RunID *string

	/* SkipDeletedTests.

	   Do not include executions from deleted tests
	*/
	SkipDeletedTests *string

	/* TargetRevision.

	   Revision of the target routing context
	*/
	TargetRevision *string

	/* TargetRouteGroup.

	   Target route group name
	*/
	TargetRouteGroup *string

	/* TargetRoutingKey.

	   Target routing key
	*/
	TargetRoutingKey *string

	/* TargetSandbox.

	   Target sandbox name
	*/
	TargetSandbox *string

	/* TestName.

	   Test name
	*/
	TestName *string

	/* To.

	   Filter results until this time
	*/
	To *string

	/* TriggerID.

	   Trigger name
	*/
	TriggerID *string

	/* TriggerWorkload.

	   Trigger workload in JSON format
	*/
	TriggerWorkload *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query test executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryTestExecutionsParams) WithDefaults() *QueryTestExecutionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query test executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryTestExecutionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query test executions params
func (o *QueryTestExecutionsParams) WithTimeout(timeout time.Duration) *QueryTestExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query test executions params
func (o *QueryTestExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query test executions params
func (o *QueryTestExecutionsParams) WithContext(ctx context.Context) *QueryTestExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query test executions params
func (o *QueryTestExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query test executions params
func (o *QueryTestExecutionsParams) WithHTTPClient(client *http.Client) *QueryTestExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query test executions params
func (o *QueryTestExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the query test executions params
func (o *QueryTestExecutionsParams) WithCluster(cluster *string) *QueryTestExecutionsParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the query test executions params
func (o *QueryTestExecutionsParams) SetCluster(cluster *string) {
	o.Cluster = cluster
}

// WithCursor adds the cursor to the query test executions params
func (o *QueryTestExecutionsParams) WithCursor(cursor *string) *QueryTestExecutionsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the query test executions params
func (o *QueryTestExecutionsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithExecutionPhase adds the executionPhase to the query test executions params
func (o *QueryTestExecutionsParams) WithExecutionPhase(executionPhase *string) *QueryTestExecutionsParams {
	o.SetExecutionPhase(executionPhase)
	return o
}

// SetExecutionPhase adds the executionPhase to the query test executions params
func (o *QueryTestExecutionsParams) SetExecutionPhase(executionPhase *string) {
	o.ExecutionPhase = executionPhase
}

// WithFrom adds the from to the query test executions params
func (o *QueryTestExecutionsParams) WithFrom(from *string) *QueryTestExecutionsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the query test executions params
func (o *QueryTestExecutionsParams) SetFrom(from *string) {
	o.From = from
}

// WithIsAutoDiff adds the isAutoDiff to the query test executions params
func (o *QueryTestExecutionsParams) WithIsAutoDiff(isAutoDiff *string) *QueryTestExecutionsParams {
	o.SetIsAutoDiff(isAutoDiff)
	return o
}

// SetIsAutoDiff adds the isAutoDiff to the query test executions params
func (o *QueryTestExecutionsParams) SetIsAutoDiff(isAutoDiff *string) {
	o.IsAutoDiff = isAutoDiff
}

// WithLabel adds the label to the query test executions params
func (o *QueryTestExecutionsParams) WithLabel(label []string) *QueryTestExecutionsParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the query test executions params
func (o *QueryTestExecutionsParams) SetLabel(label []string) {
	o.Label = label
}

// WithOrderDir adds the orderDir to the query test executions params
func (o *QueryTestExecutionsParams) WithOrderDir(orderDir *string) *QueryTestExecutionsParams {
	o.SetOrderDir(orderDir)
	return o
}

// SetOrderDir adds the orderDir to the query test executions params
func (o *QueryTestExecutionsParams) SetOrderDir(orderDir *string) {
	o.OrderDir = orderDir
}

// WithOrgName adds the orgName to the query test executions params
func (o *QueryTestExecutionsParams) WithOrgName(orgName string) *QueryTestExecutionsParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the query test executions params
func (o *QueryTestExecutionsParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithPageSize adds the pageSize to the query test executions params
func (o *QueryTestExecutionsParams) WithPageSize(pageSize *int64) *QueryTestExecutionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the query test executions params
func (o *QueryTestExecutionsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPublished adds the published to the query test executions params
func (o *QueryTestExecutionsParams) WithPublished(published *bool) *QueryTestExecutionsParams {
	o.SetPublished(published)
	return o
}

// SetPublished adds the published to the query test executions params
func (o *QueryTestExecutionsParams) SetPublished(published *bool) {
	o.Published = published
}

// WithRepo adds the repo to the query test executions params
func (o *QueryTestExecutionsParams) WithRepo(repo *string) *QueryTestExecutionsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the query test executions params
func (o *QueryTestExecutionsParams) SetRepo(repo *string) {
	o.Repo = repo
}

// WithRepoBranch adds the repoBranch to the query test executions params
func (o *QueryTestExecutionsParams) WithRepoBranch(repoBranch *string) *QueryTestExecutionsParams {
	o.SetRepoBranch(repoBranch)
	return o
}

// SetRepoBranch adds the repoBranch to the query test executions params
func (o *QueryTestExecutionsParams) SetRepoBranch(repoBranch *string) {
	o.RepoBranch = repoBranch
}

// WithRepoCommitSHA adds the repoCommitSHA to the query test executions params
func (o *QueryTestExecutionsParams) WithRepoCommitSHA(repoCommitSHA *string) *QueryTestExecutionsParams {
	o.SetRepoCommitSHA(repoCommitSHA)
	return o
}

// SetRepoCommitSHA adds the repoCommitSHA to the query test executions params
func (o *QueryTestExecutionsParams) SetRepoCommitSHA(repoCommitSHA *string) {
	o.RepoCommitSHA = repoCommitSHA
}

// WithRepoPath adds the repoPath to the query test executions params
func (o *QueryTestExecutionsParams) WithRepoPath(repoPath *string) *QueryTestExecutionsParams {
	o.SetRepoPath(repoPath)
	return o
}

// SetRepoPath adds the repoPath to the query test executions params
func (o *QueryTestExecutionsParams) SetRepoPath(repoPath *string) {
	o.RepoPath = repoPath
}

// WithRunID adds the runID to the query test executions params
func (o *QueryTestExecutionsParams) WithRunID(runID *string) *QueryTestExecutionsParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the query test executions params
func (o *QueryTestExecutionsParams) SetRunID(runID *string) {
	o.RunID = runID
}

// WithSkipDeletedTests adds the skipDeletedTests to the query test executions params
func (o *QueryTestExecutionsParams) WithSkipDeletedTests(skipDeletedTests *string) *QueryTestExecutionsParams {
	o.SetSkipDeletedTests(skipDeletedTests)
	return o
}

// SetSkipDeletedTests adds the skipDeletedTests to the query test executions params
func (o *QueryTestExecutionsParams) SetSkipDeletedTests(skipDeletedTests *string) {
	o.SkipDeletedTests = skipDeletedTests
}

// WithTargetRevision adds the targetRevision to the query test executions params
func (o *QueryTestExecutionsParams) WithTargetRevision(targetRevision *string) *QueryTestExecutionsParams {
	o.SetTargetRevision(targetRevision)
	return o
}

// SetTargetRevision adds the targetRevision to the query test executions params
func (o *QueryTestExecutionsParams) SetTargetRevision(targetRevision *string) {
	o.TargetRevision = targetRevision
}

// WithTargetRouteGroup adds the targetRouteGroup to the query test executions params
func (o *QueryTestExecutionsParams) WithTargetRouteGroup(targetRouteGroup *string) *QueryTestExecutionsParams {
	o.SetTargetRouteGroup(targetRouteGroup)
	return o
}

// SetTargetRouteGroup adds the targetRouteGroup to the query test executions params
func (o *QueryTestExecutionsParams) SetTargetRouteGroup(targetRouteGroup *string) {
	o.TargetRouteGroup = targetRouteGroup
}

// WithTargetRoutingKey adds the targetRoutingKey to the query test executions params
func (o *QueryTestExecutionsParams) WithTargetRoutingKey(targetRoutingKey *string) *QueryTestExecutionsParams {
	o.SetTargetRoutingKey(targetRoutingKey)
	return o
}

// SetTargetRoutingKey adds the targetRoutingKey to the query test executions params
func (o *QueryTestExecutionsParams) SetTargetRoutingKey(targetRoutingKey *string) {
	o.TargetRoutingKey = targetRoutingKey
}

// WithTargetSandbox adds the targetSandbox to the query test executions params
func (o *QueryTestExecutionsParams) WithTargetSandbox(targetSandbox *string) *QueryTestExecutionsParams {
	o.SetTargetSandbox(targetSandbox)
	return o
}

// SetTargetSandbox adds the targetSandbox to the query test executions params
func (o *QueryTestExecutionsParams) SetTargetSandbox(targetSandbox *string) {
	o.TargetSandbox = targetSandbox
}

// WithTestName adds the testName to the query test executions params
func (o *QueryTestExecutionsParams) WithTestName(testName *string) *QueryTestExecutionsParams {
	o.SetTestName(testName)
	return o
}

// SetTestName adds the testName to the query test executions params
func (o *QueryTestExecutionsParams) SetTestName(testName *string) {
	o.TestName = testName
}

// WithTo adds the to to the query test executions params
func (o *QueryTestExecutionsParams) WithTo(to *string) *QueryTestExecutionsParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the query test executions params
func (o *QueryTestExecutionsParams) SetTo(to *string) {
	o.To = to
}

// WithTriggerID adds the triggerID to the query test executions params
func (o *QueryTestExecutionsParams) WithTriggerID(triggerID *string) *QueryTestExecutionsParams {
	o.SetTriggerID(triggerID)
	return o
}

// SetTriggerID adds the triggerId to the query test executions params
func (o *QueryTestExecutionsParams) SetTriggerID(triggerID *string) {
	o.TriggerID = triggerID
}

// WithTriggerWorkload adds the triggerWorkload to the query test executions params
func (o *QueryTestExecutionsParams) WithTriggerWorkload(triggerWorkload *string) *QueryTestExecutionsParams {
	o.SetTriggerWorkload(triggerWorkload)
	return o
}

// SetTriggerWorkload adds the triggerWorkload to the query test executions params
func (o *QueryTestExecutionsParams) SetTriggerWorkload(triggerWorkload *string) {
	o.TriggerWorkload = triggerWorkload
}

// WriteToRequest writes these params to a swagger request
func (o *QueryTestExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {

		// query param cluster
		var qrCluster string

		if o.Cluster != nil {
			qrCluster = *o.Cluster
		}
		qCluster := qrCluster
		if qCluster != "" {

			if err := r.SetQueryParam("cluster", qCluster); err != nil {
				return err
			}
		}
	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string

		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {

			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}
	}

	if o.ExecutionPhase != nil {

		// query param executionPhase
		var qrExecutionPhase string

		if o.ExecutionPhase != nil {
			qrExecutionPhase = *o.ExecutionPhase
		}
		qExecutionPhase := qrExecutionPhase
		if qExecutionPhase != "" {

			if err := r.SetQueryParam("executionPhase", qExecutionPhase); err != nil {
				return err
			}
		}
	}

	if o.From != nil {

		// query param from
		var qrFrom string

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	if o.IsAutoDiff != nil {

		// query param isAutoDiff
		var qrIsAutoDiff string

		if o.IsAutoDiff != nil {
			qrIsAutoDiff = *o.IsAutoDiff
		}
		qIsAutoDiff := qrIsAutoDiff
		if qIsAutoDiff != "" {

			if err := r.SetQueryParam("isAutoDiff", qIsAutoDiff); err != nil {
				return err
			}
		}
	}

	if o.Label != nil {

		// binding items for label
		joinedLabel := o.bindParamLabel(reg)

		// query array param label
		if err := r.SetQueryParam("label", joinedLabel...); err != nil {
			return err
		}
	}

	if o.OrderDir != nil {

		// query param orderDir
		var qrOrderDir string

		if o.OrderDir != nil {
			qrOrderDir = *o.OrderDir
		}
		qOrderDir := qrOrderDir
		if qOrderDir != "" {

			if err := r.SetQueryParam("orderDir", qOrderDir); err != nil {
				return err
			}
		}
	}

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Published != nil {

		// query param published
		var qrPublished bool

		if o.Published != nil {
			qrPublished = *o.Published
		}
		qPublished := swag.FormatBool(qrPublished)
		if qPublished != "" {

			if err := r.SetQueryParam("published", qPublished); err != nil {
				return err
			}
		}
	}

	if o.Repo != nil {

		// query param repo
		var qrRepo string

		if o.Repo != nil {
			qrRepo = *o.Repo
		}
		qRepo := qrRepo
		if qRepo != "" {

			if err := r.SetQueryParam("repo", qRepo); err != nil {
				return err
			}
		}
	}

	if o.RepoBranch != nil {

		// query param repoBranch
		var qrRepoBranch string

		if o.RepoBranch != nil {
			qrRepoBranch = *o.RepoBranch
		}
		qRepoBranch := qrRepoBranch
		if qRepoBranch != "" {

			if err := r.SetQueryParam("repoBranch", qRepoBranch); err != nil {
				return err
			}
		}
	}

	if o.RepoCommitSHA != nil {

		// query param repoCommitSHA
		var qrRepoCommitSHA string

		if o.RepoCommitSHA != nil {
			qrRepoCommitSHA = *o.RepoCommitSHA
		}
		qRepoCommitSHA := qrRepoCommitSHA
		if qRepoCommitSHA != "" {

			if err := r.SetQueryParam("repoCommitSHA", qRepoCommitSHA); err != nil {
				return err
			}
		}
	}

	if o.RepoPath != nil {

		// query param repoPath
		var qrRepoPath string

		if o.RepoPath != nil {
			qrRepoPath = *o.RepoPath
		}
		qRepoPath := qrRepoPath
		if qRepoPath != "" {

			if err := r.SetQueryParam("repoPath", qRepoPath); err != nil {
				return err
			}
		}
	}

	if o.RunID != nil {

		// query param runID
		var qrRunID string

		if o.RunID != nil {
			qrRunID = *o.RunID
		}
		qRunID := qrRunID
		if qRunID != "" {

			if err := r.SetQueryParam("runID", qRunID); err != nil {
				return err
			}
		}
	}

	if o.SkipDeletedTests != nil {

		// query param skipDeletedTests
		var qrSkipDeletedTests string

		if o.SkipDeletedTests != nil {
			qrSkipDeletedTests = *o.SkipDeletedTests
		}
		qSkipDeletedTests := qrSkipDeletedTests
		if qSkipDeletedTests != "" {

			if err := r.SetQueryParam("skipDeletedTests", qSkipDeletedTests); err != nil {
				return err
			}
		}
	}

	if o.TargetRevision != nil {

		// query param targetRevision
		var qrTargetRevision string

		if o.TargetRevision != nil {
			qrTargetRevision = *o.TargetRevision
		}
		qTargetRevision := qrTargetRevision
		if qTargetRevision != "" {

			if err := r.SetQueryParam("targetRevision", qTargetRevision); err != nil {
				return err
			}
		}
	}

	if o.TargetRouteGroup != nil {

		// query param targetRouteGroup
		var qrTargetRouteGroup string

		if o.TargetRouteGroup != nil {
			qrTargetRouteGroup = *o.TargetRouteGroup
		}
		qTargetRouteGroup := qrTargetRouteGroup
		if qTargetRouteGroup != "" {

			if err := r.SetQueryParam("targetRouteGroup", qTargetRouteGroup); err != nil {
				return err
			}
		}
	}

	if o.TargetRoutingKey != nil {

		// query param targetRoutingKey
		var qrTargetRoutingKey string

		if o.TargetRoutingKey != nil {
			qrTargetRoutingKey = *o.TargetRoutingKey
		}
		qTargetRoutingKey := qrTargetRoutingKey
		if qTargetRoutingKey != "" {

			if err := r.SetQueryParam("targetRoutingKey", qTargetRoutingKey); err != nil {
				return err
			}
		}
	}

	if o.TargetSandbox != nil {

		// query param targetSandbox
		var qrTargetSandbox string

		if o.TargetSandbox != nil {
			qrTargetSandbox = *o.TargetSandbox
		}
		qTargetSandbox := qrTargetSandbox
		if qTargetSandbox != "" {

			if err := r.SetQueryParam("targetSandbox", qTargetSandbox); err != nil {
				return err
			}
		}
	}

	if o.TestName != nil {

		// query param testName
		var qrTestName string

		if o.TestName != nil {
			qrTestName = *o.TestName
		}
		qTestName := qrTestName
		if qTestName != "" {

			if err := r.SetQueryParam("testName", qTestName); err != nil {
				return err
			}
		}
	}

	if o.To != nil {

		// query param to
		var qrTo string

		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo
		if qTo != "" {

			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}
	}

	if o.TriggerID != nil {

		// query param triggerID
		var qrTriggerID string

		if o.TriggerID != nil {
			qrTriggerID = *o.TriggerID
		}
		qTriggerID := qrTriggerID
		if qTriggerID != "" {

			if err := r.SetQueryParam("triggerID", qTriggerID); err != nil {
				return err
			}
		}
	}

	if o.TriggerWorkload != nil {

		// query param triggerWorkload
		var qrTriggerWorkload string

		if o.TriggerWorkload != nil {
			qrTriggerWorkload = *o.TriggerWorkload
		}
		qTriggerWorkload := qrTriggerWorkload
		if qTriggerWorkload != "" {

			if err := r.SetQueryParam("triggerWorkload", qTriggerWorkload); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamQueryTestExecutions binds the parameter label
func (o *QueryTestExecutionsParams) bindParamLabel(formats strfmt.Registry) []string {
	labelIR := o.Label

	var labelIC []string
	for _, labelIIR := range labelIR { // explode []string

		labelIIV := labelIIR // string as string
		labelIC = append(labelIC, labelIIV)
	}

	// items.CollectionFormat: ""
	labelIS := swag.JoinByFormat(labelIC, "")

	return labelIS
}
