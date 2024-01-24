// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TraceqlFilter] = (*TraceqlFilterBuilder)(nil)

type TraceqlFilterBuilder struct {
	internal *TraceqlFilter
	errors   map[string]cog.BuildErrors
}

func NewTraceqlFilterBuilder() *TraceqlFilterBuilder {
	resource := &TraceqlFilter{}
	builder := &TraceqlFilterBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TraceqlFilterBuilder) Build() (TraceqlFilter, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("TraceqlFilter", err)...)
	}

	if len(errs) != 0 {
		return TraceqlFilter{}, errs
	}

	return *builder.internal, nil
}

// Uniquely identify the filter, will not be used in the query generation
func (builder *TraceqlFilterBuilder) Id(id string) *TraceqlFilterBuilder {
	builder.internal.Id = id

	return builder
}

// The tag for the search filter, for example: .http.status_code, .service.name, status
func (builder *TraceqlFilterBuilder) Tag(tag string) *TraceqlFilterBuilder {
	builder.internal.Tag = &tag

	return builder
}

// The operator that connects the tag to the value, for example: =, >, !=, =~
func (builder *TraceqlFilterBuilder) Operator(operator string) *TraceqlFilterBuilder {
	builder.internal.Operator = &operator

	return builder
}

// The value for the search filter
func (builder *TraceqlFilterBuilder) Value(arrayOfString []string) *TraceqlFilterBuilder {
	if builder.internal.Value == nil {
		builder.internal.Value = &StringOrArrayOfString{}
	}
	builder.internal.Value.ArrayOfString = arrayOfString

	return builder
}

// The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
func (builder *TraceqlFilterBuilder) ValueType(valueType string) *TraceqlFilterBuilder {
	builder.internal.ValueType = &valueType

	return builder
}

// The scope of the filter, can either be unscoped/all scopes, resource or span
func (builder *TraceqlFilterBuilder) Scope(scope TraceqlSearchScope) *TraceqlFilterBuilder {
	builder.internal.Scope = &scope

	return builder
}

func (builder *TraceqlFilterBuilder) applyDefaults() {
}
