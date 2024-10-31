// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureMonitorResource] = (*AzureMonitorResourceBuilder)(nil)

type AzureMonitorResourceBuilder struct {
	internal *AzureMonitorResource
	errors   map[string]cog.BuildErrors
}

func NewAzureMonitorResourceBuilder() *AzureMonitorResourceBuilder {
	resource := &AzureMonitorResource{}
	builder := &AzureMonitorResourceBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AzureMonitorResourceBuilder) Build() (AzureMonitorResource, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureMonitorResource{}, err
	}

	return *builder.internal, nil
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

func (builder *AzureMonitorResourceBuilder) applyDefaults() {
}
