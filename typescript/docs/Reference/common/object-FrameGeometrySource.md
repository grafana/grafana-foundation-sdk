---
title: <span class="badge object-type-interface"></span> FrameGeometrySource
---
# <span class="badge object-type-interface"></span> FrameGeometrySource

## Definition

```typescript
export interface FrameGeometrySource {
	mode: common.FrameGeometrySourceMode;
	// Field mappings
	geohash?: string;
	latitude?: string;
	longitude?: string;
	wkt?: string;
	lookup?: string;
	// Path to Gazetteer
	gazetteer?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [FrameGeometrySourceBuilder](./builder-FrameGeometrySourceBuilder.md)
