---
title: <span class="badge builder"></span> NotebookLayoutItemSpecBuilder
---
# <span class="badge builder"></span> NotebookLayoutItemSpecBuilder

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Constructor

```php
new NotebookLayoutItemSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> collapsed

```php
collapsed(bool $collapsed)
```

### <span class="badge object-method"></span> element

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ElementReference> $element

```php
element(\Grafana\Foundation\Cog\Builder $element)
```

### <span class="badge object-method"></span> source

```php
source(\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecSource $source)
```

## See also

 * <span class="badge object-type-class"></span> [NotebookLayoutItemSpec](./object-NotebookLayoutItemSpec.md)
