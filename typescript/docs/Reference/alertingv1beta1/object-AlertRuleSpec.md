---
title: <span class="badge object-type-interface"></span> AlertRuleSpec
---
# <span class="badge object-type-interface"></span> AlertRuleSpec

## Definition

```typescript
export interface AlertRuleSpec {
	title: string;
	paused?: boolean;
	trigger: alertingv1beta1.IntervalTrigger;
	labels?: Record<string, alertingv1beta1.TemplateString>;
	annotations?: Record<string, alertingv1beta1.TemplateString>;
	for?: string;
	keepFiringFor?: string;
	missingSeriesEvalsToResolve?: number;
	noDataState: alertingv1beta1.NoDataState;
	execErrState: alertingv1beta1.ExecErrState;
	notificationSettings?: alertingv1beta1.NotificationSettings;
	expressions: alertingv1beta1.ExpressionMap;
	panelRef?: alertingv1beta1.PanelRef;
}

```
## See also

 * <span class="badge builder"></span> [AlertRuleBuilder](./builder-AlertRuleBuilder.md)
