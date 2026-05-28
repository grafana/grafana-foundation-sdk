---
title: <span class="badge object-type-class"></span> AnnotationPanelFilter
---
# <span class="badge object-type-class"></span> AnnotationPanelFilter

## Definition

```php
class AnnotationPanelFilter implements \JsonSerializable
{
    /**
     * Should the specified panels be included or excluded
     */
    public ?bool $exclude;

    /**
     * Panel IDs that should be included or excluded
     * @var array<int>
     */
    public array $ids;

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

 * <span class="badge builder"></span> [AnnotationPanelFilterBuilder](./builder-AnnotationPanelFilterBuilder.md)
