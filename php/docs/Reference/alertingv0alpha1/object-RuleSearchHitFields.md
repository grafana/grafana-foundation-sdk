---
title: <span class="badge object-type-class"></span> RuleSearchHitFields
---
# <span class="badge object-type-class"></span> RuleSearchHitFields

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Definition

```php
class RuleSearchHitFields implements \JsonSerializable
{
    public ?string $title;

    public ?string $folder;

    public ?string $type;

    public ?string $interval;

    public ?bool $paused;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    /**
     * @var array<string>|null
     */
    public ?array $datasourceUIDs;

    /**
     * Alert-rule fields.
     * @var array<string, string>|null
     */
    public ?array $annotations;

    public ?string $for;

    public ?string $keepFiringFor;

    public ?string $dashboardUID;

    public ?int $panelID;

    public ?string $receiver;

    public ?string $notificationType;

    public ?string $routingTree;

    /**
     * Recording-rule fields.
     */
    public ?string $metric;

    public ?string $targetDatasourceUID;

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

 * <span class="badge builder"></span> [RuleSearchHitFieldsBuilder](./builder-RuleSearchHitFieldsBuilder.md)
