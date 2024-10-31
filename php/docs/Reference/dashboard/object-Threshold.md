---
title: <span class="badge object-type-class"></span> Threshold
---
# <span class="badge object-type-class"></span> Threshold

User-defined value for a metric that triggers visual changes in a panel when this value is met or exceeded

They are used to conditionally style and color visualizations based on query results , and can be applied to most visualizations.

## Definition

```php
class Threshold implements \JsonSerializable
{
    /**
     * Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
     * Nulls currently appear here when serializing -Infinity to JSON.
     */
    public ?float $value;

    /**
     * Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
     */
    public string $color;

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

