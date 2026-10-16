---
title: <span class="badge builder"></span> V2TransformationSpecBuilder
---
# <span class="badge builder"></span> V2TransformationSpecBuilder

Dashboard v2 shape: no `id`, it is carried by the parent's `group`.

## Constructor

```php
new V2TransformationSpecBuilder()
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
filter(\Grafana\Foundation\Dashboardv2beta1\MatcherConfig $filter)
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
topic(\Grafana\Foundation\Dashboardv2beta1\DataTopic $topic)
```

## See also

 * <span class="badge object-type-class"></span> [V2TransformationSpec](./object-V2TransformationSpec.md)
