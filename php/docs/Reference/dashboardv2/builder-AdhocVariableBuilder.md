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

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\AdHocFilterWithLabels>> $baseFilters

```php
baseFilters(array $baseFilters)
```

### <span class="badge object-method"></span> datasource

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2AdhocVariableKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> defaultKeys

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\MetricFindValue>> $defaultKeys

```php
defaultKeys(array $defaultKeys)
```

### <span class="badge object-method"></span> description

```php
description(string $description)
```

### <span class="badge object-method"></span> enableGroupBy

Whether the group-by operator is enabled in the ad hoc filter combobox.

```php
enableGroupBy(bool $enableGroupBy)
```

### <span class="badge object-method"></span> filters

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\AdHocFilterWithLabels>> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> group

```php
group(string $group)
```

### <span class="badge object-method"></span> hide

```php
hide(\Grafana\Foundation\Dashboardv2\VariableHide $hide)
```

### <span class="badge object-method"></span> label

```php
label(string $label)
```

### <span class="badge object-method"></span> labels

@param array<string, string> $labels

```php
labels(array $labels)
```

### <span class="badge object-method"></span> name

```php
name(string $name)
```

### <span class="badge object-method"></span> origin

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ControlSourceRef> $origin

```php
origin(\Grafana\Foundation\Cog\Builder $origin)
```

### <span class="badge object-method"></span> skipUrlSync

```php
skipUrlSync(bool $skipUrlSync)
```

## See also

 * <span class="badge object-type-class"></span> [AdhocVariableKind](./object-AdhocVariableKind.md)
