---
title: <span class="badge object-type-interface"></span> MovingAverageHoltWintersModelSettings
---
# <span class="badge object-type-interface"></span> MovingAverageHoltWintersModelSettings

## Definition

```typescript
export interface MovingAverageHoltWintersModelSettings {
	model: "holt_winters";
	settings: {
		alpha?: string;
		beta?: string;
		gamma?: string;
		period?: string;
		pad?: boolean;
	};
	window: string;
	minimize: boolean;
	predict: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MovingAverageHoltWintersModelSettingsBuilder](./builder-MovingAverageHoltWintersModelSettingsBuilder.md)
