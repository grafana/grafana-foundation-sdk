---
title: <span class="badge object-type-class"></span> Rule
---
# <span class="badge object-type-class"></span> Rule

## Definition

```php
class Rule implements \JsonSerializable
{
    /**
     * @var array<string, string>|null
     */
    public ?array $annotations;

    public string $condition;

    /**
     * @var array<\Grafana\Foundation\Alerting\Query>
     */
    public array $data;

    public \Grafana\Foundation\Alerting\RuleExecErrState $execErrState;

    public string $folderUID;

    /**
     * The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
     * Before this time has elapsed, the rule is only considered to be Pending.
     */
    public string $for;

    public ?int $id;

    public ?bool $isPaused;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    public \Grafana\Foundation\Alerting\RuleNoDataState $noDataState;

    public int $orgID;

    public string $provenance;

    public string $ruleGroup;

    public string $title;

    public ?string $uid;

    public ?string $updated;

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

 * <span class="badge builder"></span> [RuleBuilder](./builder-RuleBuilder.md)
