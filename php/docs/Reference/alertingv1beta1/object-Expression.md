---
title: <span class="badge object-type-class"></span> Expression
---
# <span class="badge object-type-class"></span> Expression

## Definition

```php
class Expression implements \JsonSerializable
{
    /**
     * The type of query if this is a query expression
     */
    public ?string $queryType;

    public ?\Grafana\Foundation\Alertingv1beta1\RelativeTimeRange $relativeTimeRange;

    /**
     * The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
     */
    public string $datasourceUID;

    /**
     * @var mixed
     */
    public $model;

    /**
     * Used to mark the expression to be used as the final source for the rule evaluation
     * Only one expression in a rule can be marked as the source
     * For AlertRules, this is the expression that will be evaluated against the alerting condition
     * For RecordingRules, this is the expression that will be recorded
     */
    public ?bool $source;

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

 * <span class="badge builder"></span> [ExpressionBuilder](./builder-ExpressionBuilder.md)
