---
title: <span class="badge object-type-interface"></span> ThresholdsConfig
---
# <span class="badge object-type-interface"></span> ThresholdsConfig

Thresholds configuration for the panel

## Definition

```typescript
export interface ThresholdsConfig {
	// Thresholds mode.
	mode: dashboard.ThresholdsMode;
	// Must be sorted by 'value', first value is always -Infinity
	steps: dashboard.Threshold[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ThresholdsConfigBuilder](./builder-ThresholdsConfigBuilder.md)
