---
title: <span class="badge object-type-interface"></span> <span class="badge deprecated"></span> RuleGroup
---
# <span class="badge object-type-interface"></span> <span class="badge deprecated"></span> RuleGroup

<span class="badge deprecated"></span>Prefer using the alertingv1beta1 package instead

## Definition

```typescript
export interface RuleGroup {
	folderUid?: string;
	// The interval, in seconds, at which all rules in the group are evaluated.
	// If a group contains many rules, the rules are evaluated sequentially.
	interval?: alerting.Duration;
	rules?: alerting.Rule[];
	title?: string;
}

```
## See also

 * <span class="badge builder"></span> <span class="badge deprecated"></span> [RuleGroupBuilder](./builder-RuleGroupBuilder.md)
