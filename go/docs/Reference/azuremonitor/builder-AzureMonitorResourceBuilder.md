---
title: <span class="badge builder"></span> AzureMonitorResourceBuilder
---
# <span class="badge builder"></span> AzureMonitorResourceBuilder

## Constructor

```go
func NewAzureMonitorResourceBuilder() *AzureMonitorResourceBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AzureMonitorResourceBuilder) Build() (AzureMonitorResource, error)
```

### <span class="badge object-method"></span> MetricNamespace

```go
func (builder *AzureMonitorResourceBuilder) MetricNamespace(metricNamespace string) *AzureMonitorResourceBuilder
```

### <span class="badge object-method"></span> Region

```go
func (builder *AzureMonitorResourceBuilder) Region(region string) *AzureMonitorResourceBuilder
```

### <span class="badge object-method"></span> ResourceGroup

```go
func (builder *AzureMonitorResourceBuilder) ResourceGroup(resourceGroup string) *AzureMonitorResourceBuilder
```

### <span class="badge object-method"></span> ResourceName

```go
func (builder *AzureMonitorResourceBuilder) ResourceName(resourceName string) *AzureMonitorResourceBuilder
```

### <span class="badge object-method"></span> Subscription

```go
func (builder *AzureMonitorResourceBuilder) Subscription(subscription string) *AzureMonitorResourceBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AzureMonitorResource](./object-AzureMonitorResource.md)
