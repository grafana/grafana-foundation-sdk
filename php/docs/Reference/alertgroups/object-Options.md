---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    /**
     * Comma-separated list of values used to filter alert results
     */
    public string $labels;

    /**
     * Name of the alertmanager used as a source for alerts
     */
    public string $alertmanager;

    /**
     * Expand all alert groups by default
     */
    public bool $expandAll;

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

