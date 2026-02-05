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

```php
builtIn(bool $builtIn)
```

### <span class="badge object-method"></span> enable

```php
enable(bool $enable)
```

### <span class="badge object-method"></span> filter

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter> $filter

```php
filter(\Grafana\Foundation\Cog\Builder $filter)
```

### <span class="badge object-method"></span> hide

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> iconColor

```php
iconColor(string $iconColor)
```

### <span class="badge object-method"></span> legacyOptions

Catch-all field for datasource-specific properties. Should not be available in as code tooling.

@param array<string, mixed> $legacyOptions

```php
legacyOptions(array $legacyOptions)
```

### <span class="badge object-method"></span> name

```php
name(string $name)
```

### <span class="badge object-method"></span> placement

Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.

```php
placement(string $placement)
```

### <span class="badge object-method"></span> query

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind> $query

```php
query(\Grafana\Foundation\Cog\Builder $query)
```

### <span class="badge object-method"></span> spec

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec> $spec

```php
spec(\Grafana\Foundation\Cog\Builder $spec)
```

## See also

 * <span class="badge object-type-class"></span> [AnnotationQueryKind](./object-AnnotationQueryKind.md)
