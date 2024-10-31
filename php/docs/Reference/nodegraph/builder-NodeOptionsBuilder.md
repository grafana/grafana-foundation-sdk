---
title: <span class="badge builder"></span> NodeOptionsBuilder
---
# <span class="badge builder"></span> NodeOptionsBuilder

## Constructor

```php
new NodeOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> arcs

Define which fields are shown as part of the node arc (colored circle around the node).

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Nodegraph\ArcOption>> $arcs

```php
arcs(array $arcs)
```

### <span class="badge object-method"></span> mainStatUnit

Unit for the main stat to override what ever is set in the data frame.

```php
mainStatUnit(string $mainStatUnit)
```

### <span class="badge object-method"></span> secondaryStatUnit

Unit for the secondary stat to override what ever is set in the data frame.

```php
secondaryStatUnit(string $secondaryStatUnit)
```

## See also

 * <span class="badge object-type-class"></span> [NodeOptions](./object-NodeOptions.md)
