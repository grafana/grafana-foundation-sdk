---
title: <span class="badge builder"></span> TransformationBuilder
---
# <span class="badge builder"></span> TransformationBuilder

## Constructor

```php
new TransformationBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> disabled

Disabled transformations are skipped

```php
disabled(bool $disabled)
```

### <span class="badge object-method"></span> filter

Optional frame matcher. When missing it will be applied to all results

```php
filter(\Grafana\Foundation\Dashboardv2\MatcherConfig $filter)
```

### <span class="badge object-method"></span> group

The group is the transformation ID

```php
group(string $group)
```

### <span class="badge object-method"></span> options

Options to be passed to the transformer

Valid options depend on the transformer id

@param mixed $options

```php
options($options)
```

### <span class="badge object-method"></span> topic

Where to pull DataFrames from as input to transformation

```php
topic(\Grafana\Foundation\Dashboardv2\DataTopic $topic)
```

## See also

 * <span class="badge object-type-class"></span> [TransformationKind](./object-TransformationKind.md)
