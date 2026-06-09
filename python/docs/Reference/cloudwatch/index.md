# <span class="badge package-variant-dataquery"></span> cloudwatch

## Objects

 * <span class="badge object-type-class"></span> [AnnotationQuery](./object-AnnotationQuery.md)
 * <span class="badge object-type-map"></span> [Dimensions](./object-Dimensions.md)
 * <span class="badge object-type-class"></span> [LogGroup](./object-LogGroup.md)
 * <span class="badge object-type-class"></span> [LogsQuery](./object-LogsQuery.md)
 * <span class="badge object-type-enum"></span> [LogsQueryLanguage](./object-LogsQueryLanguage.md)
 * <span class="badge object-type-enum"></span> [MetricEditorMode](./object-MetricEditorMode.md)
 * <span class="badge object-type-enum"></span> [MetricQueryType](./object-MetricQueryType.md)
 * <span class="badge object-type-class"></span> [MetricStat](./object-MetricStat.md)
 * <span class="badge object-type-class"></span> [MetricsQuery](./object-MetricsQuery.md)
 * <span class="badge object-type-class"></span> [QueryEditorArrayExpression](./object-QueryEditorArrayExpression.md)
 * <span class="badge object-type-disjunction"></span> [QueryEditorExpression](./object-QueryEditorExpression.md)
 * <span class="badge object-type-enum"></span> [QueryEditorExpressionType](./object-QueryEditorExpressionType.md)
 * <span class="badge object-type-class"></span> [QueryEditorFunctionExpression](./object-QueryEditorFunctionExpression.md)
 * <span class="badge object-type-class"></span> [QueryEditorFunctionParameterExpression](./object-QueryEditorFunctionParameterExpression.md)
 * <span class="badge object-type-class"></span> [QueryEditorGroupByExpression](./object-QueryEditorGroupByExpression.md)
 * <span class="badge object-type-class"></span> [QueryEditorOperator](./object-QueryEditorOperator.md)
 * <span class="badge object-type-class"></span> [QueryEditorOperatorExpression](./object-QueryEditorOperatorExpression.md)
 * <span class="badge object-type-disjunction"></span> [QueryEditorOperatorType](./object-QueryEditorOperatorType.md)
 * <span class="badge object-type-disjunction"></span> [QueryEditorOperatorValueType](./object-QueryEditorOperatorValueType.md)
 * <span class="badge object-type-class"></span> [QueryEditorProperty](./object-QueryEditorProperty.md)
 * <span class="badge object-type-class"></span> [QueryEditorPropertyExpression](./object-QueryEditorPropertyExpression.md)
 * <span class="badge object-type-enum"></span> [QueryEditorPropertyType](./object-QueryEditorPropertyType.md)
 * <span class="badge object-type-enum"></span> [QueryMode](./object-QueryMode.md)
 * <span class="badge object-type-disjunction"></span> [Request](./object-Request.md)
 * <span class="badge object-type-class"></span> [SQLExpression](./object-SQLExpression.md)
## Builders

 * <span class="badge builder"></span> [AnnotationQuery](./builder-AnnotationQuery.md)
 * <span class="badge builder"></span> [LogGroup](./builder-LogGroup.md)
 * <span class="badge builder"></span> [LogsQuery](./builder-LogsQuery.md)
 * <span class="badge builder"></span> [MetricStat](./builder-MetricStat.md)
 * <span class="badge builder"></span> [MetricsQuery](./builder-MetricsQuery.md)
 * <span class="badge builder"></span> [Query](./builder-Query.md)
 * <span class="badge builder"></span> [QueryEditorArrayExpression](./builder-QueryEditorArrayExpression.md)
 * <span class="badge builder"></span> [QueryEditorFunctionExpression](./builder-QueryEditorFunctionExpression.md)
 * <span class="badge builder"></span> [QueryEditorFunctionParameterExpression](./builder-QueryEditorFunctionParameterExpression.md)
 * <span class="badge builder"></span> [QueryEditorGroupByExpression](./builder-QueryEditorGroupByExpression.md)
 * <span class="badge builder"></span> [QueryEditorOperator](./builder-QueryEditorOperator.md)
 * <span class="badge builder"></span> [QueryEditorOperatorExpression](./builder-QueryEditorOperatorExpression.md)
 * <span class="badge builder"></span> [QueryEditorProperty](./builder-QueryEditorProperty.md)
 * <span class="badge builder"></span> [QueryEditorPropertyExpression](./builder-QueryEditorPropertyExpression.md)
 * <span class="badge builder"></span> [QueryV2](./builder-QueryV2.md)
 * <span class="badge builder"></span> [SQLExpression](./builder-SQLExpression.md)
## Functions

### <span class="badge function"></span> variant_config

variant_config returns the configuration related to cloudwatch data queries.

This configuration describes how to unmarshal it, convert it to code, …

```python
def variant_config() -> variants.DataqueryConfig
```

