---
title: <span class="badge object-type-class"></span> QueryEditorArrayExpression
---
# <span class="badge object-type-class"></span> QueryEditorArrayExpression

## Definition

```php
class QueryEditorArrayExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionType $type;

    /**
     * @var array<\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression|\Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression|\Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression|\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression|\Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression|\Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression>
     */
    public array $expressions;

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

 * <span class="badge builder"></span> [QueryEditorArrayExpressionBuilder](./builder-QueryEditorArrayExpressionBuilder.md)
