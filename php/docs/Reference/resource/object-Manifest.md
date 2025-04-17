---
title: <span class="badge object-type-class"></span> Manifest
---
# <span class="badge object-type-class"></span> Manifest

## Definition

```php
class Manifest implements \JsonSerializable
{
    public string $apiVersion;

    public string $kind;

    public \Grafana\Foundation\Resource\Metadata $metadata;

    /**
     * @var mixed
     */
    public $spec;

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

 * <span class="badge builder"></span> [ManifestBuilder](./builder-ManifestBuilder.md)
