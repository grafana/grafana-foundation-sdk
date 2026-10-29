---
title: <span class="badge object-type-interface"></span> AlertRuleStatus
---
# <span class="badge object-type-interface"></span> AlertRuleStatus

## Definition

```typescript
export interface AlertRuleStatus {
	health?: alertingv1beta1.AlertRuleHealth;
	state?: alertingv1beta1.AlertRuleState;
	stateReason?: alertingv1beta1.AlertRuleStateReason;
	lastEvaluationTime?: string;
	// duration of the last evaluation in seconds
	evaluationDuration?: number;
	lastError?: string;
	totals?: alertingv1beta1.AlertRuleInstanceTotals;
}

```
## See also

 * <span class="badge builder"></span> [AlertRuleStatusBuilder](./builder-AlertRuleStatusBuilder.md)
