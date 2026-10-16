---
title: <span class="badge object-type-interface"></span> Threshold
---
# <span class="badge object-type-interface"></span> Threshold

## Definition

```typescript
export interface Threshold {
	// Value null means -Infinity
	value: number | null;
	// Optional dashboard-variable expression (e.g. `$myVar`) resolved at render time; `value` is the numeric fallback when the expression cannot be resolved to a single finite number.
	valueExpr?: string;
	color: string;
}

```
