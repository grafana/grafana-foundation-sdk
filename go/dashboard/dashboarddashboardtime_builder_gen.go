// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardDashboardTime] = (*DashboardDashboardTimeBuilder)(nil)

type DashboardDashboardTimeBuilder struct {
	internal *DashboardDashboardTime
	errors   map[string]cog.BuildErrors
}

func NewDashboardDashboardTimeBuilder() *DashboardDashboardTimeBuilder {
	resource := NewDashboardDashboardTime()
	builder := &DashboardDashboardTimeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *DashboardDashboardTimeBuilder) Build() (DashboardDashboardTime, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardDashboardTime{}, err
	}

	return *builder.internal, nil
}

func (builder *DashboardDashboardTimeBuilder) From(from string) *DashboardDashboardTimeBuilder {
	builder.internal.From = from

	return builder
}

func (builder *DashboardDashboardTimeBuilder) To(to string) *DashboardDashboardTimeBuilder {
	builder.internal.To = to

	return builder
}
