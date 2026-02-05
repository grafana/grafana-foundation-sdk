---
title: <span class="badge object-type-class"></span> FetchOptions
---
# <span class="badge object-type-class"></span> FetchOptions

## Definition

```php
class FetchOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\HttpRequestMethod $method;

    public string $url;

    public ?string $body;

    /**
     * These are 2D arrays of strings, each representing a key-value pair
     * We are defining them this way because we can't generate a go struct that
     * that would have exactly two strings in each sub-array
     * @var array<array<string>>|null
     */
    public ?array $queryParams;

    /**
     * @var array<array<string>>|null
     */
    public ?array $headers;

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

 * <span class="badge builder"></span> [FetchOptionsBuilder](./builder-FetchOptionsBuilder.md)
