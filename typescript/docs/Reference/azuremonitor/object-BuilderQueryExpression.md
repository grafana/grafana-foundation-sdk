---
title: <span class="badge object-type-interface"></span> BuilderQueryExpression
---
# <span class="badge object-type-interface"></span> BuilderQueryExpression

## Definition

```typescript
export interface BuilderQueryExpression {
	from?: azuremonitor.BuilderQueryEditorPropertyExpression;
	columns?: azuremonitor.BuilderQueryEditorColumnsExpression;
	where?: azuremonitor.BuilderQueryEditorWhereExpressionArray;
	reduce?: azuremonitor.BuilderQueryEditorReduceExpressionArray;
	groupBy?: azuremonitor.BuilderQueryEditorGroupByExpressionArray;
	limit?: number;
	orderBy?: azuremonitor.BuilderQueryEditorOrderByExpressionArray;
	fuzzySearch?: azuremonitor.BuilderQueryEditorWhereExpressionArray;
	timeFilter?: azuremonitor.BuilderQueryEditorWhereExpressionArray;
}

```
## See also

 * <span class="badge builder"></span> [BuilderQueryExpressionBuilder](./builder-BuilderQueryExpressionBuilder.md)
