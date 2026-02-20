// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureMonitorResource] = (*AzureMonitorResourceBuilder)(nil)

type AzureMonitorResourceBuilder struct {
	internal *AzureMonitorResource
	errors   cog.BuildErrors
}

func NewAzureMonitorResourceBuilder() *AzureMonitorResourceBuilder {
	resource := NewAzureMonitorResource()
	builder := &AzureMonitorResourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AzureMonitorResourceBuilder) Build() (AzureMonitorResource, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureMonitorResource{}, err
	}

	if len(builder.errors) > 0 {
		return AzureMonitorResource{}, cog.MakeBuildErrors("azuremonitor.azureMonitorResource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AzureMonitorResourceBuilder) RecordError(path string, err error) *AzureMonitorResourceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AzureMonitorResourceBuilder) Subscription(subscription string) *AzureMonitorResourceBuilder {
	builder.internal.Subscription = &subscription

	return builder
}

func (builder *AzureMonitorResourceBuilder) ResourceGroup(resourceGroup string) *AzureMonitorResourceBuilder {
	builder.internal.ResourceGroup = &resourceGroup

	return builder
}

func (builder *AzureMonitorResourceBuilder) ResourceName(resourceName string) *AzureMonitorResourceBuilder {
	builder.internal.ResourceName = &resourceName

	return builder
}

func (builder *AzureMonitorResourceBuilder) MetricNamespace(metricNamespace string) *AzureMonitorResourceBuilder {
	builder.internal.MetricNamespace = &metricNamespace

	return builder
}

func (builder *AzureMonitorResourceBuilder) Region(region string) *AzureMonitorResourceBuilder {
	builder.internal.Region = &region

	return builder
}
