// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2RangeMapOptions] = (*Dashboardv2RangeMapOptionsBuilder)(nil)

type Dashboardv2RangeMapOptionsBuilder struct {
	internal *Dashboardv2RangeMapOptions
	errors   cog.BuildErrors
}

func NewDashboardv2RangeMapOptionsBuilder() *Dashboardv2RangeMapOptionsBuilder {
	resource := NewDashboardv2RangeMapOptions()
	builder := &Dashboardv2RangeMapOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2RangeMapOptionsBuilder) Build() (Dashboardv2RangeMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2RangeMapOptions{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2RangeMapOptions{}, cog.MakeBuildErrors("dashboardv2.dashboardv2RangeMapOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2RangeMapOptionsBuilder) RecordError(path string, err error) *Dashboardv2RangeMapOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Min value of the range. It can be null which means -Infinity
func (builder *Dashboardv2RangeMapOptionsBuilder) From(from float64) *Dashboardv2RangeMapOptionsBuilder {
	builder.internal.From = &from

	return builder
}

// Max value of the range. It can be null which means +Infinity
func (builder *Dashboardv2RangeMapOptionsBuilder) To(to float64) *Dashboardv2RangeMapOptionsBuilder {
	builder.internal.To = &to

	return builder
}

// Config to apply when the value is within the range
func (builder *Dashboardv2RangeMapOptionsBuilder) Result(result cog.Builder[ValueMappingResult]) *Dashboardv2RangeMapOptionsBuilder {
	resultResource, err := result.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Result = resultResource

	return builder
}
