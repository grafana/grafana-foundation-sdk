---
title: <span class="badge object-type-interface"></span> SQLExpression
---
# <span class="badge object-type-interface"></span> SQLExpression

## Definition

```typescript
export interface SQLExpression {
	// SELECT part of the SQL expression
	select?: cloudwatch.QueryEditorFunctionExpression;
	// FROM part of the SQL expression
	from?: cloudwatch.QueryEditorPropertyExpression | cloudwatch.QueryEditorFunctionExpression;
	// WHERE part of the SQL expression
	where?: cloudwatch.QueryEditorArrayExpression;
	// GROUP BY part of the SQL expression
	groupBy?: cloudwatch.QueryEditorArrayExpression;
	// ORDER BY part of the SQL expression
	orderBy?: cloudwatch.QueryEditorFunctionExpression;
	// The sort order of the SQL expression, `ASC` or `DESC`
	orderByDirection?: string;
	// LIMIT part of the SQL expression
	limit?: number;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [SQLExpressionBuilder](./builder-SQLExpressionBuilder.md)
