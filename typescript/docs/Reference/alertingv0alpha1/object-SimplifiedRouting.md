---
title: <span class="badge object-type-interface"></span> SimplifiedRouting
---
# <span class="badge object-type-interface"></span> SimplifiedRouting

## Definition

```typescript
export interface SimplifiedRouting {
	type: alertingv0alpha1.NotificationSettingsType.SimplifiedRouting;
	receiver: string;
	groupBy?: string[];
	groupWait?: alertingv0alpha1.PromDuration;
	groupInterval?: alertingv0alpha1.PromDuration;
	repeatInterval?: alertingv0alpha1.PromDuration;
	muteTimeIntervals?: alertingv0alpha1.TimeIntervalRef[];
	activeTimeIntervals?: alertingv0alpha1.TimeIntervalRef[];
}

```
## See also

 * <span class="badge builder"></span> [SimplifiedRoutingBuilder](./builder-SimplifiedRoutingBuilder.md)
