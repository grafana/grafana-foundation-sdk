---
title: <span class="badge object-type-interface"></span> Rule
---
# <span class="badge object-type-interface"></span> Rule

## Definition

```typescript
export interface Rule {
	annotations?: Record<string, string>;
	condition: string;
	data: alerting.Query[];
	execErrState: "OK" | "Alerting" | "Error";
	folderUID: string;
	// The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
	// Before this time has elapsed, the rule is only considered to be Pending.
	for: string;
	id?: number;
	isPaused?: boolean;
	labels?: Record<string, string>;
	noDataState: "Alerting" | "NoData" | "OK";
	notification_settings?: alerting.NotificationSettings;
	orgID: number;
	provenance?: alerting.Provenance;
	ruleGroup: string;
	title: string;
	uid?: string;
	updated?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [RuleBuilder](./builder-RuleBuilder.md)
