---
title: <span class="badge object-type-class"></span> Metadata
---
# <span class="badge object-type-class"></span> Metadata

## Definition

```php
class Metadata implements \JsonSerializable
{
    public string $name;

    public ?string $namespace;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    /**
     * @var array<string, string>|null
     */
    public ?array $annotations;

    public ?string $uid;

    public ?string $resourceVersion;

    public ?int $generation;

    public ?string $creationTimestamp;

    public ?string $updateTimestamp;

    public ?string $deletionTimestamp;

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

 * <span class="badge builder"></span> [MetadataBuilder](./builder-MetadataBuilder.md)
