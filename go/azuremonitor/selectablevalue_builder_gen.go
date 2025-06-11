// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SelectableValue] = (*SelectableValueBuilder)(nil)

type SelectableValueBuilder struct {
	internal *SelectableValue
	errors   cog.BuildErrors
}

func NewSelectableValueBuilder() *SelectableValueBuilder {
	resource := NewSelectableValue()
	builder := &SelectableValueBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SelectableValueBuilder) Build() (SelectableValue, error) {
	if err := builder.internal.Validate(); err != nil {
		return SelectableValue{}, err
	}

	if len(builder.errors) > 0 {
		return SelectableValue{}, cog.MakeBuildErrors("azuremonitor.selectableValue", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SelectableValueBuilder) Label(label string) *SelectableValueBuilder {
	builder.internal.Label = label

	return builder
}

func (builder *SelectableValueBuilder) Value(value string) *SelectableValueBuilder {
	builder.internal.Value = value

	return builder
}
