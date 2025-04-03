---
title: <span class="badge builder"></span> AnnotationQueryBuilder
---
# <span class="badge builder"></span> AnnotationQueryBuilder

## Constructor

```php
new AnnotationQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> builtIn

Set to 1 for the standard annotation query all dashboards have by default.

```php
builtIn(float $builtIn)
```

### <span class="badge object-method"></span> datasource

Datasource where the annotations data is

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> enable

When enabled the annotation query is issued with every dashboard refresh

```php
enable(bool $enable)
```

### <span class="badge object-method"></span> expr

```php
expr(string $expr)
```

### <span class="badge object-method"></span> filter

Filters to apply when fetching annotations

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationPanelFilter> $filter

```php
filter(\Grafana\Foundation\Cog\Builder $filter)
```

### <span class="badge object-method"></span> hide

Annotation queries can be toggled on or off at the top of the dashboard.

When hide is true, the toggle is not shown in the dashboard.

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> iconColor

Color to use for the annotation event markers

```php
iconColor(string $iconColor)
```

### <span class="badge object-method"></span> name

Name of annotation.

```php
name(string $name)
```

### <span class="badge object-method"></span> target

TODO.. this should just be a normal query target

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationTarget> $target

```php
target(\Grafana\Foundation\Cog\Builder $target)
```

### <span class="badge object-method"></span> type

TODO -- this should not exist here, it is based on the --grafana-- datasource

```php
type(string $type)
```

## See also

 * <span class="badge object-type-class"></span> [AnnotationQuery](./object-AnnotationQuery.md)
