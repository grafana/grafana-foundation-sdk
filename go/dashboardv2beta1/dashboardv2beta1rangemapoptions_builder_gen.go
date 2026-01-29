// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2beta1RangeMapOptions] = (*Dashboardv2beta1RangeMapOptionsBuilder)(nil)

type Dashboardv2beta1RangeMapOptionsBuilder struct {
	internal *Dashboardv2beta1RangeMapOptions
	errors   cog.BuildErrors
}

func NewDashboardv2beta1RangeMapOptionsBuilder() *Dashboardv2beta1RangeMapOptionsBuilder {
	resource := NewDashboardv2beta1RangeMapOptions()
	builder := &Dashboardv2beta1RangeMapOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2beta1RangeMapOptionsBuilder) Build() (Dashboardv2beta1RangeMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2beta1RangeMapOptions{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2beta1RangeMapOptions{}, cog.MakeBuildErrors("dashboardv2beta1.dashboardv2beta1RangeMapOptions", builder.errors)
	}

	return *builder.internal, nil
}

// Min value of the range. It can be null which means -Infinity
func (builder *Dashboardv2beta1RangeMapOptionsBuilder) From(from float64) *Dashboardv2beta1RangeMapOptionsBuilder {
	builder.internal.From = &from

	return builder
}

// Max value of the range. It can be null which means +Infinity
func (builder *Dashboardv2beta1RangeMapOptionsBuilder) To(to float64) *Dashboardv2beta1RangeMapOptionsBuilder {
	builder.internal.To = &to

	return builder
}

// Config to apply when the value is within the range
func (builder *Dashboardv2beta1RangeMapOptionsBuilder) Result(result cog.Builder[ValueMappingResult]) *Dashboardv2beta1RangeMapOptionsBuilder {
	resultResource, err := result.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Result = resultResource

	return builder
}
