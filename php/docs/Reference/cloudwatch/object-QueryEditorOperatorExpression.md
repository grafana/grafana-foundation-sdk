---
title: <span class="badge object-type-class"></span> QueryEditorOperatorExpression
---
# <span class="badge object-type-class"></span> QueryEditorOperatorExpression

## Definition

```php
class QueryEditorOperatorExpression implements \JsonSerializable
{
    public string $type;

    public \Grafana\Foundation\Cloudwatch\QueryEditorProperty $property;

    /**
     * TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
     */
    public \Grafana\Foundation\Cloudwatch\QueryEditorOperator $operator;

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

 * <span class="badge builder"></span> [QueryEditorOperatorExpressionBuilder](./builder-QueryEditorOperatorExpressionBuilder.md)
