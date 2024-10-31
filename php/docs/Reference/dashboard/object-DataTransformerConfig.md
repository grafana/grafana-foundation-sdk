---
title: <span class="badge object-type-class"></span> DataTransformerConfig
---
# <span class="badge object-type-class"></span> DataTransformerConfig

Transformations allow to manipulate data returned by a query before the system applies a visualization.

Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,

use the output of one transformation as the input to another transformation, etc.

## Definition

```php
class DataTransformerConfig implements \JsonSerializable
{
    /**
     * Unique identifier of transformer
     */
    public string $id;

    /**
     * Disabled transformations are skipped
     */
    public ?bool $disabled;

    /**
     * Optional frame matcher. When missing it will be applied to all results
     */
    public ?\Grafana\Foundation\Dashboard\MatcherConfig $filter;

    /**
     * Options to be passed to the transformer
     * Valid options depend on the transformer id
     * @var mixed
     */
    public $options;

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

