---
title: <span class="badge object-type-class"></span> RuleGroup
---
# <span class="badge object-type-class"></span> RuleGroup

## Definition

```php
class RuleGroup implements \JsonSerializable
{
    public ?string $folderUid;

    /**
     * The interval, in seconds, at which all rules in the group are evaluated.
     * If a group contains many rules, the rules are evaluated sequentially.
     */
    public int $interval;

    /**
     * @var array<\Grafana\Foundation\Alerting\Rule>|null
     */
    public ?array $rules;

    public ?string $title;

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

 * <span class="badge builder"></span> [RuleGroupBuilder](./builder-RuleGroupBuilder.md)
