---
title: <span class="badge object-type-interface"></span> V2TransformationSpec
---
# <span class="badge object-type-interface"></span> V2TransformationSpec

Dashboard v2 shape: no `id`, it is carried by the parent's `group`.

## Definition

```typescript
export interface V2TransformationSpec {
	// Disabled transformations are skipped
	disabled?: boolean;
	// Optional frame matcher. When missing it will be applied to all results
	filter?: dashboardv2beta1.MatcherConfig;
	// Where to pull DataFrames from as input to transformation
	topic?: dashboardv2beta1.DataTopic;
	// Options to be passed to the transformer
	// Valid options depend on the transformer id
	options: any;
}

```
## See also

 * <span class="badge builder"></span> [V2TransformationSpecBuilder](./builder-V2TransformationSpecBuilder.md)
