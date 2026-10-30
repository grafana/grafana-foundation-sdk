---
title: <span class="badge object-type-class"></span> SimplifiedRouting
---
# <span class="badge object-type-class"></span> SimplifiedRouting

## Definition

```php
class SimplifiedRouting implements \JsonSerializable
{
    public \Grafana\Foundation\Alertingv1beta1\NotificationSettingsType $type;

    public string $receiver;

    /**
     * @var array<string>|null
     */
    public ?array $groupBy;

    public string $groupWait;

    public string $groupInterval;

    public string $repeatInterval;

    /**
     * @var array<string>|null
     */
    public ?array $muteTimeIntervals;

    /**
     * @var array<string>|null
     */
    public ?array $activeTimeIntervals;

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

 * <span class="badge builder"></span> [SimplifiedRoutingBuilder](./builder-SimplifiedRoutingBuilder.md)
