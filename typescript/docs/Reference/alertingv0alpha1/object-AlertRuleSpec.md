---
title: <span class="badge object-type-interface"></span> AlertRuleSpec
---
# <span class="badge object-type-interface"></span> AlertRuleSpec

## Definition

```typescript
export interface AlertRuleSpec {
	title: string;
	paused?: boolean;
	trigger: alertingv0alpha1.IntervalTrigger;
	labels?: Record<string, alertingv0alpha1.TemplateString>;
	annotations?: Record<string, alertingv0alpha1.TemplateString>;
	for?: string;
	keepFiringFor?: string;
	missingSeriesEvalsToResolve?: number;
	noDataState: alertingv0alpha1.NoDataState;
	execErrState: alertingv0alpha1.ExecErrState;
	notificationSettings?: alertingv0alpha1.NotificationSettings;
	expressions: alertingv0alpha1.ExpressionMap;
	panelRef?: alertingv0alpha1.PanelRef;
}

```
## See also

 * <span class="badge builder"></span> [AlertRuleBuilder](./builder-AlertRuleBuilder.md)
