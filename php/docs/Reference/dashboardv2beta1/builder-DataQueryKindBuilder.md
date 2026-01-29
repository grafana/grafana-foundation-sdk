---
title: <span class="badge builder"></span> DataQueryKindBuilder
---
# <span class="badge builder"></span> DataQueryKindBuilder

## Constructor

```php
new DataQueryKindBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> group

```php
group(string $group)
```

### <span class="badge object-method"></span> spec

@param mixed $spec

```php
spec($spec)
```

### <span class="badge object-method"></span> version

```php
version(string $version)
```

## See also

 * <span class="badge object-type-class"></span> [DataQueryKind](./object-DataQueryKind.md)
