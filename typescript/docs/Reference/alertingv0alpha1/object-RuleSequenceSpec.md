---
title: <span class="badge object-type-interface"></span> RuleSequenceSpec
---
# <span class="badge object-type-interface"></span> RuleSequenceSpec

## Definition

```typescript
export interface RuleSequenceSpec {
	trigger: alertingv0alpha1.IntervalTrigger;
	recordingRules: alertingv0alpha1.RuleRef[];
	alertingRules?: alertingv0alpha1.RuleRef[];
}

```
## See also

 * <span class="badge builder"></span> [RuleSequenceBuilder](./builder-RuleSequenceBuilder.md)
