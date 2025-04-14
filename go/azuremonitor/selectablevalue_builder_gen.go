// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SelectableValue] = (*SelectableValueBuilder)(nil)

type SelectableValueBuilder struct {
	internal *SelectableValue
	errors   map[string]cog.BuildErrors
}

func NewSelectableValueBuilder() *SelectableValueBuilder {
	resource := NewSelectableValue()
	builder := &SelectableValueBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *SelectableValueBuilder) Build() (SelectableValue, error) {
	if err := builder.internal.Validate(); err != nil {
		return SelectableValue{}, err
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
