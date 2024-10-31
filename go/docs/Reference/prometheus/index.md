# prometheus

## Objects

 * <span class="badge object-type-enum"></span> [PromQueryFormat](./object-PromQueryFormat.md)
 * <span class="badge object-type-enum"></span> [QueryEditorMode](./object-QueryEditorMode.md)
 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
## Builders

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to prometheus dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> DataqueryConverter

DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.

```go
func DataqueryConverter(input Dataquery) string
```

