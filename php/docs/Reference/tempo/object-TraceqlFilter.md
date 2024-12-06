---
title: <span class="badge object-type-class"></span> TraceqlFilter
---
# <span class="badge object-type-class"></span> TraceqlFilter

## Definition

```php
class TraceqlFilter implements \JsonSerializable
{
    /**
     * Uniquely identify the filter, will not be used in the query generation
     */
    public string $id;

    /**
     * The tag for the search filter, for example: .http.status_code, .service.name, status
     */
    public ?string $tag;

    /**
     * The operator that connects the tag to the value, for example: =, >, !=, =~
     */
    public ?string $operator;

    /**
     * The value for the search filter
     * @var string|array<string>|null
     */
    public $value;

    /**
     * The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
     */
    public ?string $valueType;

    /**
     * The scope of the filter, can either be unscoped/all scopes, resource or span
     */
    public ?\Grafana\Foundation\Tempo\TraceqlSearchScope $scope;

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

 * <span class="badge builder"></span> [TraceqlFilterBuilder](./builder-TraceqlFilterBuilder.md)
