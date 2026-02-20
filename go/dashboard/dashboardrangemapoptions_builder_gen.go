// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardRangeMapOptions] = (*DashboardRangeMapOptionsBuilder)(nil)

type DashboardRangeMapOptionsBuilder struct {
	internal *DashboardRangeMapOptions
	errors   cog.BuildErrors
}

func NewDashboardRangeMapOptionsBuilder() *DashboardRangeMapOptionsBuilder {
	resource := NewDashboardRangeMapOptions()
	builder := &DashboardRangeMapOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DashboardRangeMapOptionsBuilder) Build() (DashboardRangeMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardRangeMapOptions{}, err
	}

	if len(builder.errors) > 0 {
		return DashboardRangeMapOptions{}, cog.MakeBuildErrors("dashboard.dashboardRangeMapOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DashboardRangeMapOptionsBuilder) RecordError(path string, err error) *DashboardRangeMapOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Min value of the range. It can be null which means -Infinity
func (builder *DashboardRangeMapOptionsBuilder) From(from float64) *DashboardRangeMapOptionsBuilder {
	builder.internal.From = &from

	return builder
}

// Max value of the range. It can be null which means +Infinity
func (builder *DashboardRangeMapOptionsBuilder) To(to float64) *DashboardRangeMapOptionsBuilder {
	builder.internal.To = &to

	return builder
}

// Config to apply when the value is within the range
func (builder *DashboardRangeMapOptionsBuilder) Result(result ValueMappingResult) *DashboardRangeMapOptionsBuilder {
	builder.internal.Result = result

	return builder
}
