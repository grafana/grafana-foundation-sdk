---
title: <span class="badge object-type-interface"></span> ScaleDimensionConfig
---
# <span class="badge object-type-interface"></span> ScaleDimensionConfig

## Definition

```typescript
export interface ScaleDimensionConfig {
	min: number;
	max: number;
	fixed?: number;
	// fixed: T -- will be added by each element
	field?: string;
	// | *"linear"
	mode?: common.ScaleDimensionMode;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ScaleDimensionConfigBuilder](./builder-ScaleDimensionConfigBuilder.md)
