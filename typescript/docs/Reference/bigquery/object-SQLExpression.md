---
title: <span class="badge object-type-interface"></span> SQLExpression
---
# <span class="badge object-type-interface"></span> SQLExpression

## Definition

```typescript
export interface SQLExpression {
	columns?: bigquery.QueryEditorFunctionExpression[];
	from?: string;
	// whereJsonTree?: _
	whereString?: string;
	groupBy?: bigquery.QueryEditorGroupByExpression[];
	orderBy?: bigquery.QueryEditorPropertyExpression;
	orderByDirection?: bigquery.OrderByDirection;
	limit?: number;
	offset?: number;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [SQLExpressionBuilder](./builder-SQLExpressionBuilder.md)
