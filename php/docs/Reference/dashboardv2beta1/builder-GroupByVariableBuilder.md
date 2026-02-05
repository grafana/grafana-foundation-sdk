---
title: <span class="badge builder"></span> GroupByVariableBuilder
---
# <span class="badge builder"></span> GroupByVariableBuilder

## Constructor

```php
new GroupByVariableBuilder(string $name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> current

```php
current(\Grafana\Foundation\Dashboardv2beta1\VariableOption $current)
```

### <span class="badge object-method"></span> datasource

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1GroupByVariableKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> defaultValue

```php
defaultValue(\Grafana\Foundation\Dashboardv2beta1\VariableOption $defaultValue)
```

### <span class="badge object-method"></span> description

```php
description(string $description)
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

### <span class="badge object-method"></span> multi

```php
multi(bool $multi)
```

### <span class="badge object-method"></span> name

```php
name(string $name)
```

### <span class="badge object-method"></span> options

@param array<\Grafana\Foundation\Dashboardv2beta1\VariableOption> $options

```php
options(array $options)
```

### <span class="badge object-method"></span> skipUrlSync

```php
skipUrlSync(bool $skipUrlSync)
```

### <span class="badge object-method"></span> spec

```php
spec(\Grafana\Foundation\Dashboardv2beta1\GroupByVariableSpec $spec)
```

## See also

 * <span class="badge object-type-class"></span> [GroupByVariableKind](./object-GroupByVariableKind.md)
