# <span class="badge package-variant-dataquery"></span> athena

## Objects

 * <span class="badge object-type-struct"></span> [ConnectionArgs](./object-ConnectionArgs.md)
 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
 * <span class="badge object-type-enum"></span> [FormatOptions](./object-FormatOptions.md)
 * <span class="badge object-type-scalar"></span> [DefaultKey](./object-DefaultKey.md)
## Builders

 * <span class="badge builder"></span> [ConnectionArgsBuilder](./builder-ConnectionArgsBuilder.md)
 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
 * <span class="badge builder"></span> [QueryBuilder](./builder-QueryBuilder.md)
## Functions

### <span class="badge function"></span> NewDataquery

NewDataquery creates a new Dataquery object.

```go
func NewDataquery() *Dataquery
```

### <span class="badge function"></span> NewConnectionArgs

NewConnectionArgs creates a new ConnectionArgs object.

```go
func NewConnectionArgs() *ConnectionArgs
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to grafana-athena-datasource dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> DataqueryConverter

DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.

```go
func DataqueryConverter(input Dataquery) string
```

### <span class="badge function"></span> ConnectionArgsConverter

ConnectionArgsConverter accepts a `ConnectionArgs` object and generates the Go code to build this object using builders.

```go
func ConnectionArgsConverter(input ConnectionArgs) string
```

### <span class="badge function"></span> QueryConverter

QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.

```go
func QueryConverter(input dashboardv2beta1.DataQueryKind) string
```

