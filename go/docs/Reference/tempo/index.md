# tempo

## Objects

 * <span class="badge object-type-enum"></span> [SearchStreamingState](./object-SearchStreamingState.md)
 * <span class="badge object-type-struct"></span> [StringOrArrayOfString](./object-StringOrArrayOfString.md)
 * <span class="badge object-type-struct"></span> [TempoQuery](./object-TempoQuery.md)
 * <span class="badge object-type-enum"></span> [TempoQueryType](./object-TempoQueryType.md)
 * <span class="badge object-type-struct"></span> [TraceqlFilter](./object-TraceqlFilter.md)
 * <span class="badge object-type-enum"></span> [TraceqlSearchScope](./object-TraceqlSearchScope.md)
## Builders

 * <span class="badge builder"></span> [TempoQueryBuilder](./builder-TempoQueryBuilder.md)
 * <span class="badge builder"></span> [TraceqlFilterBuilder](./builder-TraceqlFilterBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to tempo dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> TempoQueryConverter

TempoQueryConverter accepts a `TempoQuery` object and generates the Go code to build this object using builders.

```go
func TempoQueryConverter(input TempoQuery) string
```

### <span class="badge function"></span> TraceqlFilterConverter

TraceqlFilterConverter accepts a `TraceqlFilter` object and generates the Go code to build this object using builders.

```go
func TraceqlFilterConverter(input TraceqlFilter) string
```

