---
title: <span class="badge object-type-class"></span> Scopes
---
# <span class="badge object-type-class"></span> Scopes

## Definition

```php
class Scopes implements \JsonSerializable
{
    /**
     * @var array<string>|null
     */
    public ?array $defaultPath;

    /**
     * @var array<\Grafana\Foundation\Prometheus\ScopesFilters>|null
     */
    public ?array $filters;

    public string $title;

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

 * <span class="badge builder"></span> [ScopesBuilder](./builder-ScopesBuilder.md)
