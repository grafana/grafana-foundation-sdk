// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryOptionsSpec] = (*QueryOptionsSpecBuilder)(nil)

type QueryOptionsSpecBuilder struct {
	internal *QueryOptionsSpec
	errors   cog.BuildErrors
}

func NewQueryOptionsSpecBuilder() *QueryOptionsSpecBuilder {
	resource := NewQueryOptionsSpec()
	builder := &QueryOptionsSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *QueryOptionsSpecBuilder) Build() (QueryOptionsSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryOptionsSpec{}, err
	}

	if len(builder.errors) > 0 {
		return QueryOptionsSpec{}, cog.MakeBuildErrors("dashboardv2beta1.queryOptionsSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryOptionsSpecBuilder) TimeFrom(timeFrom string) *QueryOptionsSpecBuilder {
	builder.internal.TimeFrom = &timeFrom

	return builder
}

func (builder *QueryOptionsSpecBuilder) MaxDataPoints(maxDataPoints int64) *QueryOptionsSpecBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

func (builder *QueryOptionsSpecBuilder) TimeShift(timeShift string) *QueryOptionsSpecBuilder {
	builder.internal.TimeShift = &timeShift

	return builder
}

func (builder *QueryOptionsSpecBuilder) QueryCachingTTL(queryCachingTTL int64) *QueryOptionsSpecBuilder {
	builder.internal.QueryCachingTTL = &queryCachingTTL

	return builder
}

func (builder *QueryOptionsSpecBuilder) Interval(interval string) *QueryOptionsSpecBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *QueryOptionsSpecBuilder) CacheTimeout(cacheTimeout string) *QueryOptionsSpecBuilder {
	builder.internal.CacheTimeout = &cacheTimeout

	return builder
}

func (builder *QueryOptionsSpecBuilder) HideTimeOverride(hideTimeOverride bool) *QueryOptionsSpecBuilder {
	builder.internal.HideTimeOverride = &hideTimeOverride

	return builder
}

func (builder *QueryOptionsSpecBuilder) TimeCompare(timeCompare string) *QueryOptionsSpecBuilder {
	builder.internal.TimeCompare = &timeCompare

	return builder
}
