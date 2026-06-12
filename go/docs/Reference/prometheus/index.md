# <span class="badge package-variant-dataquery"></span> prometheus

## Objects

 * <span class="badge object-type-struct"></span> [AdhocFilters](./object-AdhocFilters.md)
 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
 * <span class="badge object-type-enum"></span> [PromQueryFormat](./object-PromQueryFormat.md)
 * <span class="badge object-type-enum"></span> [QueryEditorMode](./object-QueryEditorMode.md)
 * <span class="badge object-type-struct"></span> [ResultAssertions](./object-ResultAssertions.md)
 * <span class="badge object-type-enum"></span> [ResultAssertionsType](./object-ResultAssertionsType.md)
 * <span class="badge object-type-struct"></span> [Scopes](./object-Scopes.md)
 * <span class="badge object-type-struct"></span> [ScopesFilters](./object-ScopesFilters.md)
 * <span class="badge object-type-struct"></span> [TimeRange](./object-TimeRange.md)
## Builders

 * <span class="badge builder"></span> [AdhocFiltersBuilder](./builder-AdhocFiltersBuilder.md)
 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
 * <span class="badge builder"></span> [QueryBuilder](./builder-QueryBuilder.md)
 * <span class="badge builder"></span> [QueryV2Builder](./builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [ResultAssertionsBuilder](./builder-ResultAssertionsBuilder.md)
 * <span class="badge builder"></span> [ScopesBuilder](./builder-ScopesBuilder.md)
 * <span class="badge builder"></span> [ScopesFiltersBuilder](./builder-ScopesFiltersBuilder.md)
 * <span class="badge builder"></span> [TimeRangeBuilder](./builder-TimeRangeBuilder.md)
## Functions

### <span class="badge function"></span> NewAdhocFilters

NewAdhocFilters creates a new AdhocFilters object.

```go
func NewAdhocFilters() *AdhocFilters
```

### <span class="badge function"></span> NewResultAssertions

NewResultAssertions creates a new ResultAssertions object.

```go
func NewResultAssertions() *ResultAssertions
```

### <span class="badge function"></span> NewScopesFilters

NewScopesFilters creates a new ScopesFilters object.

```go
func NewScopesFilters() *ScopesFilters
```

### <span class="badge function"></span> NewScopes

NewScopes creates a new Scopes object.

```go
func NewScopes() *Scopes
```

### <span class="badge function"></span> NewTimeRange

NewTimeRange creates a new TimeRange object.

```go
func NewTimeRange() *TimeRange
```

### <span class="badge function"></span> NewDataquery

NewDataquery creates a new Dataquery object.

```go
func NewDataquery() *Dataquery
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to prometheus dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> QueryV2Converter

QueryV2Converter accepts a `QueryV2` object and generates the Go code to build this object using builders.

```go
func QueryV2Converter(input dashboardv2.DataQueryKind) string
```

### <span class="badge function"></span> AdhocFiltersConverter

AdhocFiltersConverter accepts a `AdhocFilters` object and generates the Go code to build this object using builders.

```go
func AdhocFiltersConverter(input AdhocFilters) string
```

### <span class="badge function"></span> ResultAssertionsConverter

ResultAssertionsConverter accepts a `ResultAssertions` object and generates the Go code to build this object using builders.

```go
func ResultAssertionsConverter(input ResultAssertions) string
```

### <span class="badge function"></span> ScopesFiltersConverter

ScopesFiltersConverter accepts a `ScopesFilters` object and generates the Go code to build this object using builders.

```go
func ScopesFiltersConverter(input ScopesFilters) string
```

### <span class="badge function"></span> ScopesConverter

ScopesConverter accepts a `Scopes` object and generates the Go code to build this object using builders.

```go
func ScopesConverter(input Scopes) string
```

### <span class="badge function"></span> TimeRangeConverter

TimeRangeConverter accepts a `TimeRange` object and generates the Go code to build this object using builders.

```go
func TimeRangeConverter(input TimeRange) string
```

### <span class="badge function"></span> DataqueryConverter

DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.

```go
func DataqueryConverter(input Dataquery) string
```

### <span class="badge function"></span> QueryConverter

QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.

```go
func QueryConverter(input dashboardv2beta1.DataQueryKind) string
```

