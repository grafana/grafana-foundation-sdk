# <span class="badge package-variant-dataquery"></span> prometheus

## Objects

 * <span class="badge object-type-class"></span> [AdhocFilters](./object-AdhocFilters.md)
 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
 * <span class="badge object-type-enum"></span> [PromQueryFormat](./object-PromQueryFormat.md)
 * <span class="badge object-type-enum"></span> [QueryEditorMode](./object-QueryEditorMode.md)
 * <span class="badge object-type-class"></span> [ResultAssertions](./object-ResultAssertions.md)
 * <span class="badge object-type-enum"></span> [ResultAssertionsType](./object-ResultAssertionsType.md)
 * <span class="badge object-type-class"></span> [Scopes](./object-Scopes.md)
 * <span class="badge object-type-class"></span> [ScopesFilters](./object-ScopesFilters.md)
 * <span class="badge object-type-class"></span> [TimeRange](./object-TimeRange.md)
## Builders

 * <span class="badge builder"></span> [AdhocFilters](./builder-AdhocFilters.md)
 * <span class="badge builder"></span> [Dataquery](./builder-Dataquery.md)
 * <span class="badge builder"></span> [Query](./builder-Query.md)
 * <span class="badge builder"></span> [QueryV2](./builder-QueryV2.md)
 * <span class="badge builder"></span> [ResultAssertions](./builder-ResultAssertions.md)
 * <span class="badge builder"></span> [Scopes](./builder-Scopes.md)
 * <span class="badge builder"></span> [ScopesFilters](./builder-ScopesFilters.md)
 * <span class="badge builder"></span> [TimeRange](./builder-TimeRange.md)
## Functions

### <span class="badge function"></span> variant_config

variant_config returns the configuration related to prometheus data queries.

This configuration describes how to unmarshal it, convert it to code, …

```python
def variant_config() -> variants.DataqueryConfig
```

