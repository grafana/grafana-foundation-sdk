---
title: <span class="badge builder"></span> AdhocVariableBuilder
---
# <span class="badge builder"></span> AdhocVariableBuilder

## Constructor

```php
new AdhocVariableBuilder(string $name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> allowCustomValue

```php
allowCustomValue(bool $allowCustomValue)
```

### <span class="badge object-method"></span> baseFilters

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels>> $baseFilters

```php
baseFilters(array $baseFilters)
```

### <span class="badge object-method"></span> datasource

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1AdhocVariableKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> defaultKeys

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\MetricFindValue>> $defaultKeys

```php
defaultKeys(array $defaultKeys)
```

### <span class="badge object-method"></span> description

```php
description(string $description)
```

### <span class="badge object-method"></span> filters

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels>> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> group

```php
group(string $group)
```

### <span class="badge object-method"></span> hide

```php
hide(\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide)
```

### <span class="badge object-method"></span> label

```php
label(string $label)
```

### <span class="badge object-method"></span> name

```php
name(string $name)
```

### <span class="badge object-method"></span> skipUrlSync

```php
skipUrlSync(bool $skipUrlSync)
```

### <span class="badge object-method"></span> spec

```php
spec(\Grafana\Foundation\Dashboardv2beta1\AdhocVariableSpec $spec)
```

## See also

 * <span class="badge object-type-class"></span> [AdhocVariableKind](./object-AdhocVariableKind.md)
