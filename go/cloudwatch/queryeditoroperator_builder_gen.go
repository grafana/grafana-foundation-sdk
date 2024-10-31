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
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorOperator{}, err
	}

	return *builder.internal, nil
}

func (builder *QueryEditorOperatorBuilder) Name(name string) *QueryEditorOperatorBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *QueryEditorOperatorBuilder) Value(operatorTypes []QueryEditorOperatorType) *QueryEditorOperatorBuilder {
	if builder.internal.Value == nil {
		builder.internal.Value = &StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType{}
	}
	builder.internal.Value.ArrayOfQueryEditorOperatorType = operatorTypes

	return builder
}

func (builder *QueryEditorOperatorBuilder) applyDefaults() {
}
