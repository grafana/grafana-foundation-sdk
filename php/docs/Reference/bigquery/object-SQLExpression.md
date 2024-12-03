---
title: <span class="badge object-type-class"></span> SQLExpression
---
# <span class="badge object-type-class"></span> SQLExpression

## Definition

```php
class SQLExpression implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Bigquery\QueryEditorFunctionExpression>|null
     */
    public ?array $columns;

    public ?string $from;

    /**
     * whereJsonTree?: _
     */
    public ?string $whereString;

    /**
     * @var array<\Grafana\Foundation\Bigquery\QueryEditorGroupByExpression>|null
     */
    public ?array $groupBy;

    public ?\Grafana\Foundation\Bigquery\QueryEditorPropertyExpression $orderBy;

    public ?\Grafana\Foundation\Bigquery\OrderByDirection $orderByDirection;

    public ?int $limit;

    public ?int $offset;

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

 * <span class="badge builder"></span> [SQLExpressionBuilder](./builder-SQLExpressionBuilder.md)
