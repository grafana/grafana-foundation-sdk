---
title: <span class="badge object-type-class"></span> SearchResultResource
---
# <span class="badge object-type-class"></span> SearchResultResource

#SearchResultResource is the full identity of a hit.

## Definition

```php
class SearchResultResource implements \JsonSerializable
{
    public string $group;

    public string $resource;

    public string $kind;

    public string $name;

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

 * <span class="badge builder"></span> [SearchResultResourceBuilder](./builder-SearchResultResourceBuilder.md)
