---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```php
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> infinitePan

Enable infinite pan

```php
infinitePan(bool $infinitePan)
```

### <span class="badge object-method"></span> inlineEditing

Enable inline editing

```php
inlineEditing(bool $inlineEditing)
```

### <span class="badge object-method"></span> panZoom

Enable pan and zoom

```php
panZoom(bool $panZoom)
```

### <span class="badge object-method"></span> root

The root element of canvas (frame), where all canvas elements are nested

TODO: Figure out how to define a default value for this

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\CanvasOptionsRoot> $root

```php
root(\Grafana\Foundation\Cog\Builder $root)
```

### <span class="badge object-method"></span> showAdvancedTypes

Show all available element types

```php
showAdvancedTypes(bool $showAdvancedTypes)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
