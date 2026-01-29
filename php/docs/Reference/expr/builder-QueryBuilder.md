---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```php
new QueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> typeClassicConditions

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeClassicConditions> $typeClassicConditions

```php
typeClassicConditions(\Grafana\Foundation\Cog\Builder $typeClassicConditions)
```

### <span class="badge object-method"></span> typeMath

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeMath> $typeMath

```php
typeMath(\Grafana\Foundation\Cog\Builder $typeMath)
```

### <span class="badge object-method"></span> typeReduce

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeReduce> $typeReduce

```php
typeReduce(\Grafana\Foundation\Cog\Builder $typeReduce)
```

### <span class="badge object-method"></span> typeResample

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeResample> $typeResample

```php
typeResample(\Grafana\Foundation\Cog\Builder $typeResample)
```

### <span class="badge object-method"></span> typeSql

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeSql> $typeSql

```php
typeSql(\Grafana\Foundation\Cog\Builder $typeSql)
```

### <span class="badge object-method"></span> typeThreshold

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeThreshold> $typeThreshold

```php
typeThreshold(\Grafana\Foundation\Cog\Builder $typeThreshold)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> version

```php
version(string $version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
