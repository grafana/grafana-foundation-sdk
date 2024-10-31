---
title: <span class="badge object-type-interface"></span> ScalarDimensionConfig
---
# <span class="badge object-type-interface"></span> ScalarDimensionConfig

## Definition

```typescript
export interface ScalarDimensionConfig {
	min: number;
	max: number;
	fixed?: number;
	// fixed: T -- will be added by each element
	field?: string;
	mode?: common.ScalarDimensionMode;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ScalarDimensionConfigBuilder](./builder-ScalarDimensionConfigBuilder.md)
