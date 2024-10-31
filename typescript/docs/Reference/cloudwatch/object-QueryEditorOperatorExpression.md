---
title: <span class="badge object-type-interface"></span> QueryEditorOperatorExpression
---
# <span class="badge object-type-interface"></span> QueryEditorOperatorExpression

## Definition

```typescript
export interface QueryEditorOperatorExpression {
	type: "operator";
	property: cloudwatch.QueryEditorProperty;
	// TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
	operator: cloudwatch.QueryEditorOperator;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [QueryEditorOperatorExpressionBuilder](./builder-QueryEditorOperatorExpressionBuilder.md)
