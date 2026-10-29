---
title: <span class="badge object-type-interface"></span> RecordingRuleSpec
---
# <span class="badge object-type-interface"></span> RecordingRuleSpec

## Definition

```typescript
export interface RecordingRuleSpec {
	title: string;
	paused?: boolean;
	trigger: alertingv0alpha1.IntervalTrigger;
	labels?: Record<string, alertingv0alpha1.TemplateString>;
	metric: alertingv0alpha1.MetricName;
	expressions: alertingv0alpha1.ExpressionMap;
	targetDatasourceUID: alertingv0alpha1.DatasourceUID;
}

```
## See also

 * <span class="badge builder"></span> [RecordingRuleBuilder](./builder-RecordingRuleBuilder.md)
