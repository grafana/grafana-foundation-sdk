---
title: <span class="badge object-type-class"></span> DataSourceRef
---
# <span class="badge object-type-class"></span> DataSourceRef

Ref to a DataSource instance

## Definition

```php
class DataSourceRef implements \JsonSerializable
{
    /**
     * The plugin type-id
     */
    public ?string $type;

    /**
     * Specific datasource instance
     */
    public ?string $uid;

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

