---
title: <span class="badge object-type-interface"></span> Threshold
---
# <span class="badge object-type-interface"></span> Threshold

User-defined value for a metric that triggers visual changes in a panel when this value is met or exceeded

They are used to conditionally style and color visualizations based on query results , and can be applied to most visualizations.

## Definition

```typescript
export interface Threshold {
	// Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
	// Nulls currently appear here when serializing -Infinity to JSON.
	value: number | null;
	// Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
	color: string;
}

```
## Methods

No methods.
