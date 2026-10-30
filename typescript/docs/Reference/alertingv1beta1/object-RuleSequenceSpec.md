---
title: <span class="badge object-type-interface"></span> RuleSequenceSpec
---
# <span class="badge object-type-interface"></span> RuleSequenceSpec

## Definition

```typescript
export interface RuleSequenceSpec {
	trigger: alertingv1beta1.IntervalTrigger;
	recordingRules: alertingv1beta1.RuleRef[];
	alertRules?: alertingv1beta1.RuleRef[];
}

```
## See also

 * <span class="badge builder"></span> [RuleSequenceBuilder](./builder-RuleSequenceBuilder.md)
