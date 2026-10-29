---
title: <span class="badge object-type-interface"></span> RuleSearchHitFields
---
# <span class="badge object-type-interface"></span> RuleSearchHitFields

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Definition

```typescript
export interface RuleSearchHitFields {
	title?: string;
	folder?: string;
	type?: string;
	interval?: string;
	paused?: boolean;
	labels?: Record<string, string>;
	datasourceUIDs?: string[];
	// Alert-rule fields.
	annotations?: Record<string, string>;
	for?: string;
	keepFiringFor?: string;
	dashboardUID?: string;
	panelID?: number;
	receiver?: string;
	notificationType?: string;
	routingTree?: string;
	// Recording-rule fields.
	metric?: string;
	targetDatasourceUID?: string;
}

```
## See also

 * <span class="badge builder"></span> [RuleSearchHitFieldsBuilder](./builder-RuleSearchHitFieldsBuilder.md)
