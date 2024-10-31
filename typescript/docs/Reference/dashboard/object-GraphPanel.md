---
title: <span class="badge object-type-interface"></span> GraphPanel
---
# <span class="badge object-type-interface"></span> GraphPanel

Support for legacy graph panel.

@deprecated this a deprecated panel type

## Definition

```typescript
export interface GraphPanel {
	type: "graph";
	// @deprecated this is part of deprecated graph panel
	legend?: {
		show: boolean;
		sort?: string;
		sortDesc?: boolean;
	};
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [GraphPanelBuilder](./builder-GraphPanelBuilder.md)
