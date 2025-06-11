// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeOption] = (*TimeOptionBuilder)(nil)

// Counterpart for TypeScript's TimeOption type.
type TimeOptionBuilder struct {
	internal *TimeOption
	errors   cog.BuildErrors
}

func NewTimeOptionBuilder() *TimeOptionBuilder {
	resource := NewTimeOption()
	builder := &TimeOptionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TimeOptionBuilder) Build() (TimeOption, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeOption{}, err
	}

	if len(builder.errors) > 0 {
		return TimeOption{}, cog.MakeBuildErrors("dashboard.timeOption", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TimeOptionBuilder) Display(display string) *TimeOptionBuilder {
	builder.internal.Display = display

	return builder
}

func (builder *TimeOptionBuilder) From(from string) *TimeOptionBuilder {
	builder.internal.From = from

	return builder
}

func (builder *TimeOptionBuilder) To(to string) *TimeOptionBuilder {
	builder.internal.To = to

	return builder
}
