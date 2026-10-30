---
title: <span class="badge object-type-interface"></span> AlertRuleInstanceTotals
---
# <span class="badge object-type-interface"></span> AlertRuleInstanceTotals

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Definition

```typescript
export interface AlertRuleInstanceTotals {
	healthy?: number;
	firing?: number;
	pending?: number;
	recovering?: number;
	nodata?: number;
	error?: number;
}

```
## See also

 * <span class="badge builder"></span> [AlertRuleInstanceTotalsBuilder](./builder-AlertRuleInstanceTotalsBuilder.md)
