---
title: <span class="badge object-type-class"></span> SQLExpression
---
# <span class="badge object-type-class"></span> SQLExpression

## Definition

```php
class SQLExpression implements \JsonSerializable
{
    /**
     * SELECT part of the SQL expression
     */
    public ?\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression $select;

    /**
     * FROM part of the SQL expression
     * @var \Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression|\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression|null
     */
    public $from;

    /**
     * WHERE part of the SQL expression
     */
    public ?\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression $where;

    /**
     * GROUP BY part of the SQL expression
     */
    public ?\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression $groupBy;

    /**
     * ORDER BY part of the SQL expression
     */
    public ?\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression $orderBy;

    /**
     * The sort order of the SQL expression, `ASC` or `DESC`
     */
    public ?string $orderByDirection;

    /**
     * LIMIT part of the SQL expression
     */
    public ?int $limit;

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
