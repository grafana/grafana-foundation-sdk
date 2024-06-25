// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[Query] = (*QueryBuilder)(nil)

type QueryBuilder struct {
	internal *Query
	errors   map[string]cog.BuildErrors
}

func NewQueryBuilder(refId string) *QueryBuilder {
	resource := &Query{}
	builder := &QueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.RefId = &refId

	return builder
}

func (builder *QueryBuilder) Build() (Query, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Query", err)...)
	}

	if len(errs) != 0 {
		return Query{}, errs
	}

	return *builder.internal, nil
}

func (builder *QueryBuilder) DatasourceUid(datasourceUid string) *QueryBuilder {
	builder.internal.DatasourceUid = &datasourceUid

	return builder
}

func (builder *QueryBuilder) Model(model cog.Builder[cogvariants.Dataquery]) *QueryBuilder {
	modelResource, err := model.Build()
	if err != nil {
		builder.errors["model"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Model = modelResource

	return builder
}

func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

func (builder *QueryBuilder) RefId(refId string) *QueryBuilder {
	builder.internal.RefId = &refId

	return builder
}

func (builder *QueryBuilder) RelativeTimeRange(from Duration, to Duration) *QueryBuilder {
	if builder.internal.RelativeTimeRange == nil {
		builder.internal.RelativeTimeRange = &RelativeTimeRange{}
	}
	builder.internal.RelativeTimeRange.From = &from
	builder.internal.RelativeTimeRange.To = &to

	return builder
}

func (builder *QueryBuilder) applyDefaults() {
}
