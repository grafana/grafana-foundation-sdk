---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    /**
     * Enable inline editing
     */
    public bool $inlineEditing;

    /**
     * Show all available element types
     */
    public bool $showAdvancedTypes;

    /**
     * The root element of canvas (frame), where all canvas elements are nested
     * TODO: Figure out how to define a default value for this
     */
    public \Grafana\Foundation\Canvas\CanvasOptionsRoot $root;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

