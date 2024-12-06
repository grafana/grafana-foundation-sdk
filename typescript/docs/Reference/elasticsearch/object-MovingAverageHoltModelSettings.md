---
title: <span class="badge object-type-interface"></span> MovingAverageHoltModelSettings
---
# <span class="badge object-type-interface"></span> MovingAverageHoltModelSettings

## Definition

```typescript
export interface MovingAverageHoltModelSettings {
	model: "holt";
	settings: {
		alpha?: string;
		beta?: string;
	};
	window: string;
	minimize: boolean;
	predict: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MovingAverageHoltModelSettingsBuilder](./builder-MovingAverageHoltModelSettingsBuilder.md)
