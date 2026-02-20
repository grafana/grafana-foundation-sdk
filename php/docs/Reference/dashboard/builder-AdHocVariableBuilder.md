---
title: <span class="badge builder"></span> AdHocVariableBuilder
---
# <span class="badge builder"></span> AdHocVariableBuilder

## Constructor

```php
new AdHocVariableBuilder(string $name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> allowCustomValue

Allow custom values to be entered in the variable

```php
allowCustomValue(bool $allowCustomValue)
```

### <span class="badge object-method"></span> datasource

Data source used to fetch values for a variable. It can be defined but `null`.

```php
datasource(\Grafana\Foundation\Common\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> definition

```php
definition(string $definition)
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

### <span class="badge object-method"></span> label

Optional display name

```php
label(string $label)
```

### <span class="badge object-method"></span> name

Name of variable

```php
name(string $name)
```

### <span class="badge object-method"></span> staticOptions

Additional static options for query variable

@param array<\Grafana\Foundation\Dashboard\VariableOption> $staticOptions

```php
staticOptions(array $staticOptions)
```

### <span class="badge object-method"></span> staticOptionsOrder

Ordering of static options in relation to options returned from data source for query variable

```php
staticOptionsOrder(\Grafana\Foundation\Dashboard\VariableModelStaticOptionsOrder $staticOptionsOrder)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)
