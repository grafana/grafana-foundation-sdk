---
title: <span class="badge object-type-interface"></span> QueryEditorFunctionExpression
---
# <span class="badge object-type-interface"></span> QueryEditorFunctionExpression

## Definition

```typescript
export interface QueryEditorFunctionExpression {
	type: bigquery.QueryEditorExpressionType.Function;
	name?: string;
	parameters?: bigquery.QueryEditorFunctionParameterExpression[];
}

```
## See also

 * <span class="badge builder"></span> [QueryEditorFunctionExpressionBuilder](./builder-QueryEditorFunctionExpressionBuilder.md)
