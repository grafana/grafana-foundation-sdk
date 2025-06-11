// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableModel] = (*IntervalVariableBuilder)(nil)

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
type IntervalVariableBuilder struct {
	internal *VariableModel
	errors   cog.BuildErrors
}

func NewIntervalVariableBuilder(name string) *IntervalVariableBuilder {
	resource := NewVariableModel()
	builder := &IntervalVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Name = name
	builder.internal.Type = "interval"

	return builder
}

func (builder *IntervalVariableBuilder) Build() (VariableModel, error) {
	if err := builder.internal.Validate(); err != nil {
		return VariableModel{}, err
	}

	if len(builder.errors) > 0 {
		return VariableModel{}, cog.MakeBuildErrors("dashboard.intervalVariable", builder.errors)
	}

	return *builder.internal, nil
}

// Name of variable
func (builder *IntervalVariableBuilder) Name(name string) *IntervalVariableBuilder {
	builder.internal.Name = name

	return builder
}

// Optional display name
func (builder *IntervalVariableBuilder) Label(label string) *IntervalVariableBuilder {
	builder.internal.Label = &label

	return builder
}

// Visibility configuration for the variable
func (builder *IntervalVariableBuilder) Hide(hide VariableHide) *IntervalVariableBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Description of variable. It can be defined but `null`.
func (builder *IntervalVariableBuilder) Description(description string) *IntervalVariableBuilder {
	builder.internal.Description = &description

	return builder
}

// Query used to fetch values for a variable
func (builder *IntervalVariableBuilder) Values(query StringOrMap) *IntervalVariableBuilder {
	builder.internal.Query = &query

	return builder
}

// Shows current selected variable text/value on the dashboard
func (builder *IntervalVariableBuilder) Current(current VariableOption) *IntervalVariableBuilder {
	builder.internal.Current = &current

	return builder
}

// Options that can be selected for a variable.
func (builder *IntervalVariableBuilder) Options(options []VariableOption) *IntervalVariableBuilder {
	builder.internal.Options = options

	return builder
}

// Dynamically calculates interval by dividing time range by the count specified.
func (builder *IntervalVariableBuilder) Auto(auto bool) *IntervalVariableBuilder {
	builder.internal.Auto = &auto

	return builder
}

// The minimum threshold below which the step count intervals will not divide the time.
func (builder *IntervalVariableBuilder) MinInterval(autoMin string) *IntervalVariableBuilder {
	builder.internal.AutoMin = &autoMin

	return builder
}

// How many times the current time range should be divided to calculate the value, similar to the Max data points query option.
// For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.
func (builder *IntervalVariableBuilder) StepCount(autoCount int32) *IntervalVariableBuilder {
	builder.internal.AutoCount = &autoCount

	return builder
}
