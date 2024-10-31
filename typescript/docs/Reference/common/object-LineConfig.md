---
title: <span class="badge object-type-interface"></span> LineConfig
---
# <span class="badge object-type-interface"></span> LineConfig

TODO docs

## Definition

```typescript
export interface LineConfig {
	lineColor?: string;
	lineWidth?: number;
	lineInterpolation?: common.LineInterpolation;
	lineStyle?: common.LineStyle;
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	spanNulls?: boolean | number;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [LineConfigBuilder](./builder-LineConfigBuilder.md)
