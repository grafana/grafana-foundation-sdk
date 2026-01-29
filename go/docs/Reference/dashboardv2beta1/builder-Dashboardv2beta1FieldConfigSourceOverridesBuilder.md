---
title: <span class="badge builder"></span> Dashboardv2beta1FieldConfigSourceOverridesBuilder
---
# <span class="badge builder"></span> Dashboardv2beta1FieldConfigSourceOverridesBuilder

## Constructor

```go
func NewDashboardv2beta1FieldConfigSourceOverridesBuilder() *Dashboardv2beta1FieldConfigSourceOverridesBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *Dashboardv2beta1FieldConfigSourceOverridesBuilder) Build() (Dashboardv2beta1FieldConfigSourceOverrides, error)
```

### <span class="badge object-method"></span> SystemRef

Describes config override rules created when interacting with Grafana.

```go
func (builder *Dashboardv2beta1FieldConfigSourceOverridesBuilder) SystemRef(systemRef string) *Dashboardv2beta1FieldConfigSourceOverridesBuilder
```

### <span class="badge object-method"></span> Matcher

```go
func (builder *Dashboardv2beta1FieldConfigSourceOverridesBuilder) Matcher(matcher dashboardv2beta1.MatcherConfig) *Dashboardv2beta1FieldConfigSourceOverridesBuilder
```

### <span class="badge object-method"></span> Properties

```go
func (builder *Dashboardv2beta1FieldConfigSourceOverridesBuilder) Properties(properties []dashboardv2beta1.DynamicConfigValue) *Dashboardv2beta1FieldConfigSourceOverridesBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dashboardv2beta1FieldConfigSourceOverrides](./object-Dashboardv2beta1FieldConfigSourceOverrides.md)
