---
title: <span class="badge object-type-interface"></span> SimplifiedRouting
---
# <span class="badge object-type-interface"></span> SimplifiedRouting

## Definition

```typescript
export interface SimplifiedRouting {
	type: alertingv1beta1.NotificationSettingsType.SimplifiedRouting;
	receiver: string;
	groupBy?: string[];
	groupWait?: alertingv1beta1.PromDuration;
	groupInterval?: alertingv1beta1.PromDuration;
	repeatInterval?: alertingv1beta1.PromDuration;
	muteTimeIntervals?: alertingv1beta1.TimeIntervalRef[];
	activeTimeIntervals?: alertingv1beta1.TimeIntervalRef[];
}

```
## See also

 * <span class="badge builder"></span> [SimplifiedRoutingBuilder](./builder-SimplifiedRoutingBuilder.md)
