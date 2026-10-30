---
title: <span class="badge object-type-interface"></span> RecordingRuleStatus
---
# <span class="badge object-type-interface"></span> RecordingRuleStatus

## Definition

```typescript
export interface RecordingRuleStatus {
	health?: alertingv1beta1.RecordingRuleHealth;
	lastEvaluationTime?: string;
	// duration of the last evaluation in seconds
	evaluationDuration?: number;
	lastError?: string;
}

```
## See also

 * <span class="badge builder"></span> [RecordingRuleStatusBuilder](./builder-RecordingRuleStatusBuilder.md)
