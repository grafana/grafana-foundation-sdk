---
title: <span class="badge object-type-interface"></span> MovingAverageEWMAModelSettings
---
# <span class="badge object-type-interface"></span> MovingAverageEWMAModelSettings

## Definition

```typescript
export interface MovingAverageEWMAModelSettings {
	model: "ewma";
	settings?: {
		alpha?: string;
	};
	window: string;
	minimize: boolean;
	predict: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MovingAverageEWMAModelSettingsBuilder](./builder-MovingAverageEWMAModelSettingsBuilder.md)
