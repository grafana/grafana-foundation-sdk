// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MonitorResource] = (*MonitorResourceBuilder)(nil)

type MonitorResourceBuilder struct {
	internal *MonitorResource
	errors   cog.BuildErrors
}

func NewMonitorResourceBuilder() *MonitorResourceBuilder {
	resource := NewMonitorResource()
	builder := &MonitorResourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MonitorResourceBuilder) Build() (MonitorResource, error) {
	if err := builder.internal.Validate(); err != nil {
		return MonitorResource{}, err
	}

	if len(builder.errors) > 0 {
		return MonitorResource{}, cog.MakeBuildErrors("azuremonitor.monitorResource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MonitorResourceBuilder) RecordError(path string, err error) *MonitorResourceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *MonitorResourceBuilder) Subscription(subscription string) *MonitorResourceBuilder {
	builder.internal.Subscription = &subscription

	return builder
}

func (builder *MonitorResourceBuilder) ResourceGroup(resourceGroup string) *MonitorResourceBuilder {
	builder.internal.ResourceGroup = &resourceGroup

	return builder
}

func (builder *MonitorResourceBuilder) ResourceName(resourceName string) *MonitorResourceBuilder {
	builder.internal.ResourceName = &resourceName

	return builder
}

func (builder *MonitorResourceBuilder) MetricNamespace(metricNamespace string) *MonitorResourceBuilder {
	builder.internal.MetricNamespace = &metricNamespace

	return builder
}

func (builder *MonitorResourceBuilder) Region(region string) *MonitorResourceBuilder {
	builder.internal.Region = &region

	return builder
}
