---
title: <span class="badge builder"></span> IntervalVariableBuilder
---
# <span class="badge builder"></span> IntervalVariableBuilder

## Constructor

```php
new IntervalVariableBuilder(string $name)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> auto

Dynamically calculates interval by dividing time range by the count specified.

```php
auto(bool $auto)
```

### <span class="badge object-method"></span> current

Shows current selected variable text/value on the dashboard

```php
current(\Grafana\Foundation\Dashboard\VariableOption $current)
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

### <span class="badge object-method"></span> minInterval

The minimum threshold below which the step count intervals will not divide the time.

```php
minInterval(string $autoMin)
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

### <span class="badge object-method"></span> stepCount

How many times the current time range should be divided to calculate the value, similar to the Max data points query option.

For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.

```php
stepCount(int $autoCount)
```

### <span class="badge object-method"></span> values

Query used to fetch values for a variable

@param string|array<string, mixed> $query

```php
values($query)
```

## See also

 * <span class="badge object-type-class"></span> [VariableModel](./object-VariableModel.md)