// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorOperator] = (*QueryEditorOperatorBuilder)(nil)

// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
type QueryEditorOperatorBuilder struct {
	internal *QueryEditorOperator
	errors   map[string]cog.BuildErrors
}

func NewQueryEditorOperatorBuilder() *QueryEditorOperatorBuilder {
	resource := &QueryEditorOperator{}
	builder := &QueryEditorOperatorBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *QueryEditorOperatorBuilder) Build() (QueryEditorOperator, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorOperator", err)...)
	}

	if len(errs) != 0 {
		return QueryEditorOperator{}, errs
	}

	return *builder.internal, nil
}

func (builder *QueryEditorOperatorBuilder) Name(name string) *QueryEditorOperatorBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *QueryEditorOperatorBuilder) Value(value StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) *QueryEditorOperatorBuilder {
	builder.internal.Value = &value

	return builder
}

func (builder *QueryEditorOperatorBuilder) applyDefaults() {
}
