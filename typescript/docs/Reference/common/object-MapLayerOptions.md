---
title: <span class="badge object-type-interface"></span> MapLayerOptions
---
# <span class="badge object-type-interface"></span> MapLayerOptions

## Definition

```typescript
export interface MapLayerOptions {
	type: string;
	// configured unique display name
	name: string;
	// Custom options depending on the type
	config?: any;
	// Common method to define geometry fields
	location?: common.FrameGeometrySource;
	// Defines a frame MatcherConfig that may filter data for the given layer
	filterData?: any;
	// Common properties:
	// https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
	// Layer opacity (0-1)
	opacity?: number;
	// Check tooltip (defaults to true)
	tooltip?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MapLayerOptionsBuilder](./builder-MapLayerOptionsBuilder.md)
