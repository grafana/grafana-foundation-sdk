---
title: <span class="badge object-type-class"></span> RecordingRuleSpec
---
# <span class="badge object-type-class"></span> RecordingRuleSpec

## Definition

```php
class RecordingRuleSpec implements \JsonSerializable
{
    public string $title;

    public ?bool $paused;

    public \Grafana\Foundation\Alertingv1beta1\IntervalTrigger $trigger;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    public string $metric;

    /**
     * @var array<string, \Grafana\Foundation\Alertingv1beta1\Expression>
     */
    public array $expressions;

    public string $targetDatasourceUID;

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

 * <span class="badge builder"></span> [RecordingRuleBuilder](./builder-RecordingRuleBuilder.md)
