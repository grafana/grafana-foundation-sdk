---
title: <span class="badge object-type-interface"></span> RecordingRuleSpec
---
# <span class="badge object-type-interface"></span> RecordingRuleSpec

## Definition

```typescript
export interface RecordingRuleSpec {
	title: string;
	paused?: boolean;
	trigger: alertingv1beta1.IntervalTrigger;
	labels?: Record<string, alertingv1beta1.TemplateString>;
	metric: alertingv1beta1.MetricName;
	expressions: alertingv1beta1.ExpressionMap;
	targetDatasourceUID: alertingv1beta1.DatasourceUID;
}

```
## See also

 * <span class="badge builder"></span> [RecordingRuleBuilder](./builder-RecordingRuleBuilder.md)
