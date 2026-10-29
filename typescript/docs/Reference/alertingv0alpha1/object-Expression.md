---
title: <span class="badge object-type-interface"></span> Expression
---
# <span class="badge object-type-interface"></span> Expression

## Definition

```typescript
export interface Expression {
	// The type of query if this is a query expression
	queryType?: string;
	relativeTimeRange?: alertingv0alpha1.RelativeTimeRange;
	// The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
	datasourceUID?: alertingv0alpha1.DatasourceUID;
	model: any;
	// Used to mark the expression to be used as the final source for the rule evaluation
	// Only one expression in a rule can be marked as the source
	// For AlertRules, this is the expression that will be evaluated against the alerting condition
	// For RecordingRules, this is the expression that will be recorded
	source?: boolean;
}

```
## See also

 * <span class="badge builder"></span> [ExpressionBuilder](./builder-ExpressionBuilder.md)
