---
title: <span class="badge object-type-class"></span> Action
---
# <span class="badge object-type-class"></span> Action

## Definition

```php
class Action implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\ActionType $type;

    public string $title;

    public ?\Grafana\Foundation\Dashboardv2beta1\FetchOptions $fetch;

    public ?\Grafana\Foundation\Dashboardv2beta1\InfinityOptions $infinity;

    public ?string $confirmation;

    public ?bool $oneClick;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\ActionVariable>|null
     */
    public ?array $variables;

    public ?\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle $style;

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

 * <span class="badge builder"></span> [ActionBuilder](./builder-ActionBuilder.md)
