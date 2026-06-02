---
title: <span class="badge builder"></span> MonitorResourceBuilder
---
# <span class="badge builder"></span> MonitorResourceBuilder

## Constructor

```go
func NewMonitorResourceBuilder() *MonitorResourceBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MonitorResourceBuilder) Build() (MonitorResource, error)
```

### <span class="badge object-method"></span> MetricNamespace

```go
func (builder *MonitorResourceBuilder) MetricNamespace(metricNamespace string) *MonitorResourceBuilder
```

### <span class="badge object-method"></span> Region

```go
func (builder *MonitorResourceBuilder) Region(region string) *MonitorResourceBuilder
```

### <span class="badge object-method"></span> ResourceGroup

```go
func (builder *MonitorResourceBuilder) ResourceGroup(resourceGroup string) *MonitorResourceBuilder
```

### <span class="badge object-method"></span> ResourceName

```go
func (builder *MonitorResourceBuilder) ResourceName(resourceName string) *MonitorResourceBuilder
```

### <span class="badge object-method"></span> Subscription

```go
func (builder *MonitorResourceBuilder) Subscription(subscription string) *MonitorResourceBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MonitorResource](./object-MonitorResource.md)
