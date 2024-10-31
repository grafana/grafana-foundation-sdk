---
title: <span class="badge object-type-interface"></span> RuleGroup
---
# <span class="badge object-type-interface"></span> RuleGroup

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
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [RuleGroupBuilder](./builder-RuleGroupBuilder.md)
