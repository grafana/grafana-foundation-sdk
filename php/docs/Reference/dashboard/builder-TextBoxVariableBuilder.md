---
title: <span class="badge builder"></span> TextBoxVariableBuilder
---
# <span class="badge builder"></span> TextBoxVariableBuilder

## Constructor

```php
new TextBoxVariableBuilder(string $name)
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

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```php
current(\Grafana\Foundation\Dashboard\VariableOption $current)
```

### <span class="badge object-method"></span> defaultValue

Query used to fetch values for a variable

@param string|array<string, mixed> $query

```php
defaultValue($query)
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

### <span class="badge object-method"></span> options

Options that can be selected for a variable.

@param array<\Grafana\Foundation\Dashboard\VariableOption> $options

```php
options(array $options)
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
