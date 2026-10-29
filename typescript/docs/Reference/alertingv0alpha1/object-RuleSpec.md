---
title: <span class="badge object-type-interface"></span> RuleSpec
---
# <span class="badge object-type-interface"></span> RuleSpec

## Definition

```typescript
export interface RuleSpec {
	title: string;
	paused?: boolean;
	trigger: alertingv0alpha1.IntervalTrigger;
	labels?: Record<string, alertingv0alpha1.TemplateString>;
	expressions: alertingv0alpha1.ExpressionMap;
}

```
## See also

 * <span class="badge builder"></span> [RuleSpecBuilder](./builder-RuleSpecBuilder.md)
