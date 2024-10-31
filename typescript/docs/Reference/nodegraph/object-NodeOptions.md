---
title: <span class="badge object-type-interface"></span> NodeOptions
---
# <span class="badge object-type-interface"></span> NodeOptions

## Definition

```typescript
export interface NodeOptions {
	// Unit for the main stat to override what ever is set in the data frame.
	mainStatUnit?: string;
	// Unit for the secondary stat to override what ever is set in the data frame.
	secondaryStatUnit?: string;
	// Define which fields are shown as part of the node arc (colored circle around the node).
	arcs?: nodegraph.ArcOption[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [NodeOptionsBuilder](./builder-NodeOptionsBuilder.md)
