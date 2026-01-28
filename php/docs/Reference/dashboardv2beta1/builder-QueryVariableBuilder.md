---
title: <span class="badge builder"></span> QueryVariableBuilder
---
# <span class="badge builder"></span> QueryVariableBuilder

## Constructor

```php
new QueryVariableBuilder(string $name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> allValue

```php
allValue(string $allValue)
```

### <span class="badge object-method"></span> allowCustomValue

```php
allowCustomValue(bool $allowCustomValue)
```

### <span class="badge object-method"></span> current

```php
current(\Grafana\Foundation\Dashboardv2beta1\VariableOption $current)
```

### <span class="badge object-method"></span> definition

```php
definition(string $definition)
```

### <span class="badge object-method"></span> description

```php
description(string $description)
```

### <span class="badge object-method"></span> hide

```php
hide(\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide)
```

### <span class="badge object-method"></span> includeAll

```php
includeAll(bool $includeAll)
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

### <span class="badge object-method"></span> placeholder

```php
placeholder(string $placeholder)
```

### <span class="badge object-method"></span> query

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind> $query

```php
query(\Grafana\Foundation\Cog\Builder $query)
```

### <span class="badge object-method"></span> refresh

```php
refresh(\Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh)
```

### <span class="badge object-method"></span> regex

```php
regex(string $regex)
```

### <span class="badge object-method"></span> regexApplyTo

```php
regexApplyTo(\Grafana\Foundation\Dashboardv2beta1\VariableRegexApplyTo $regexApplyTo)
```

### <span class="badge object-method"></span> skipUrlSync

```php
skipUrlSync(bool $skipUrlSync)
```

### <span class="badge object-method"></span> sort

```php
sort(\Grafana\Foundation\Dashboardv2beta1\VariableSort $sort)
```

### <span class="badge object-method"></span> spec

```php
spec(\Grafana\Foundation\Dashboardv2beta1\QueryVariableSpec $spec)
```

### <span class="badge object-method"></span> staticOptions

@param array<\Grafana\Foundation\Dashboardv2beta1\VariableOption> $staticOptions

```php
staticOptions(array $staticOptions)
```

### <span class="badge object-method"></span> staticOptionsOrder

```php
staticOptionsOrder(\Grafana\Foundation\Dashboardv2beta1\QueryVariableSpecStaticOptionsOrder $staticOptionsOrder)
```

## See also

 * <span class="badge object-type-class"></span> [QueryVariableKind](./object-QueryVariableKind.md)
