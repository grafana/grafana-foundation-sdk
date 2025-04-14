---
title: <span class="badge object-type-class"></span> BuilderQueryExpression
---
# <span class="badge object-type-class"></span> BuilderQueryExpression

## Definition

```php
class BuilderQueryExpression implements \JsonSerializable
{
    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyExpression $from;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression $columns;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $where;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpressionArray $reduce;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray $groupBy;

    public ?int $limit;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByExpressionArray $orderBy;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $fuzzySearch;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $timeFilter;

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

 * <span class="badge builder"></span> [BuilderQueryExpressionBuilder](./builder-BuilderQueryExpressionBuilder.md)
