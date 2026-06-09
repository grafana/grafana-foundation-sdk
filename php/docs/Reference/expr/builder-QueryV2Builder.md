---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```php
new QueryV2Builder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> classicConditions

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeClassicConditions> $typeClassicConditions

```php
classicConditions(\Grafana\Foundation\Cog\Builder $typeClassicConditions)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> labels

@param array<string, string> $labels

```php
labels(array $labels)
```

### <span class="badge object-method"></span> math

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeMath> $typeMath

```php
math(\Grafana\Foundation\Cog\Builder $typeMath)
```

### <span class="badge object-method"></span> reduce

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeReduce> $typeReduce

```php
reduce(\Grafana\Foundation\Cog\Builder $typeReduce)
```

### <span class="badge object-method"></span> resample

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeResample> $typeResample

```php
resample(\Grafana\Foundation\Cog\Builder $typeResample)
```

### <span class="badge object-method"></span> sql

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeSql> $typeSql

```php
sql(\Grafana\Foundation\Cog\Builder $typeSql)
```

### <span class="badge object-method"></span> threshold

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeThreshold> $typeThreshold

```php
threshold(\Grafana\Foundation\Cog\Builder $typeThreshold)
```

### <span class="badge object-method"></span> version

```php
version(string $version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
