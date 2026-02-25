---
title: <span class="badge object-type-class"></span> Dashboardv2beta1
---
# <span class="badge object-type-class"></span> Dashboardv2beta1

## Definition

```php
class Dashboardv2beta1 implements \JsonSerializable
{
}
```
## Methods

### <span class="badge object-method"></span> manifest

Creates a resource manifest from a Dashboard.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboard> $dashboard

```php
static manifest(string $name, \Grafana\Foundation\Cog\Builder $dashboard)
```

### <span class="badge object-method"></span> rows

```php
static rows()
```

### <span class="badge object-method"></span> row

```php
static row(string $title)
```

### <span class="badge object-method"></span> autoGrid

```php
static autoGrid()
```

### <span class="badge object-method"></span> autoGridItem

```php
static autoGridItem(string $name)
```

### <span class="badge object-method"></span> tabs

```php
static tabs()
```

### <span class="badge object-method"></span> tab

```php
static tab(string $title)
```

### <span class="badge object-method"></span> grid

```php
static grid()
```

### <span class="badge object-method"></span> gridItem

```php
static gridItem(string $name)
```

