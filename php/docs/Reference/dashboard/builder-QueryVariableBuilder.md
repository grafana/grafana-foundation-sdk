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

Custom all value

```php
allValue(string $allValue)
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```php
current(\Grafana\Foundation\Dashboard\VariableOption $current)
```

### <span class="badge object-method"></span> datasource

Data source used to fetch values for a variable. It can be defined but `null`.

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> description

Description of variable. It can be defined but `null`.

```php
description(string $description)
```

### <span class="badge object-method"></span> hide

Visibility configuration for the variable

```php
hide(\Grafana\Foundation\Dashboard\VariableHide $hide)
```

### <span class="badge object-method"></span> includeAll

Whether all value option is available or not

```php
includeAll(bool $includeAll)
```

### <span class="badge object-method"></span> label

Optional display name

```php
label(string $label)
```

### <span class="badge object-method"></span> multi

Whether multiple values can be selected or not from variable value list

```php
multi(bool $multi)
```

### <span class="badge object-method"></span> name

Name of variable

```php
name(string $name)
```

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

@param array<\Grafana\Foundation\Dashboard\VariableOption> $options

```php
options(array $options)
```

### <span class="badge object-method"></span> query

Query used to fetch values for a variable

@param string|array<string, mixed> $query

```php
query($query)
```

### <span class="badge object-method"></span> refresh

Options to config when to refresh a variable

```php
refresh(\Grafana\Foundation\Dashboard\VariableRefresh $refresh)
```

### <span class="badge object-method"></span> regex

Optional field, if you want to extract part of a series name or metric node segment.

Named capture groups can be used to separate the display text and value.

```php
regex(string $regex)
```

### <span class="badge object-method"></span> sort

Options sort order

```php
sort(\Grafana\Foundation\Dashboard\VariableSort $sort)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
