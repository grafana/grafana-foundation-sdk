---
title: <span class="badge object-type-class"></span> ControlsOptions
---
# <span class="badge object-type-class"></span> ControlsOptions

## Definition

```php
class ControlsOptions implements \JsonSerializable
{
    /**
     * Zoom (upper left)
     */
    public ?bool $showZoom;

    /**
     * let the mouse wheel zoom
     */
    public ?bool $mouseWheelZoom;

    /**
     * Lower right
     */
    public ?bool $showAttribution;

    /**
     * Scale options
     */
    public ?bool $showScale;

    /**
     * Show debug
     */
    public ?bool $showDebug;

    /**
     * Show measure
     */
    public ?bool $showMeasure;

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

## See also

 * <span class="badge builder"></span> [ControlsOptionsBuilder](./builder-ControlsOptionsBuilder.md)
