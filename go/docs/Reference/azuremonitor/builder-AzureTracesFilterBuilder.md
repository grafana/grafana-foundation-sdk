---
title: <span class="badge builder"></span> AzureTracesFilterBuilder
---
# <span class="badge builder"></span> AzureTracesFilterBuilder

## Constructor

```go
func NewAzureTracesFilterBuilder() *AzureTracesFilterBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AzureTracesFilterBuilder) Build() (AzureTracesFilter, error)
```

### <span class="badge object-method"></span> Filters

Values to filter by.

```go
func (builder *AzureTracesFilterBuilder) Filters(filters []string) *AzureTracesFilterBuilder
```

### <span class="badge object-method"></span> Operation

Comparison operator to use. Either equals or not equals.

```go
func (builder *AzureTracesFilterBuilder) Operation(operation string) *AzureTracesFilterBuilder
```

### <span class="badge object-method"></span> Property

Property name, auto-populated based on available traces.

```go
func (builder *AzureTracesFilterBuilder) Property(property string) *AzureTracesFilterBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AzureTracesFilter](./object-AzureTracesFilter.md)
