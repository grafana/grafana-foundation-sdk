// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[Query] = (*QueryBuilder)(nil)

type QueryBuilder struct {
	internal *Query
	errors   cog.BuildErrors
}

func NewQueryBuilder(refId string) *QueryBuilder {
	resource := NewQuery()
	builder := &QueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.RefId = &refId

	return builder
}

func (builder *QueryBuilder) Build() (Query, error) {
	if err := builder.internal.Validate(); err != nil {
		return Query{}, err
	}

	if len(builder.errors) > 0 {
		return Query{}, cog.MakeBuildErrors("alerting.query", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryBuilder) RecordError(path string, err error) *QueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
func (builder *QueryBuilder) DatasourceUid(datasourceUid string) *QueryBuilder {
	builder.internal.DatasourceUid = &datasourceUid

	return builder
}

// JSON is the raw JSON query and includes the above properties as well as custom properties.
func (builder *QueryBuilder) Model(model cog.Builder[variants.Dataquery]) *QueryBuilder {
	modelResource, err := model.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Model = modelResource

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder {
	builder.internal.RefId = &refId

	return builder
}

// RelativeTimeRange is the per query start and end time
// for requests.
func (builder *QueryBuilder) RelativeTimeRange(from Duration, to Duration) *QueryBuilder {
	if builder.internal.RelativeTimeRange == nil {
		builder.internal.RelativeTimeRange = NewRelativeTimeRange()
	}
	builder.internal.RelativeTimeRange.From = &from
	builder.internal.RelativeTimeRange.To = &to

	return builder
}
