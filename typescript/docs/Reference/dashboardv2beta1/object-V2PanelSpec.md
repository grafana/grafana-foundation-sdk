---
title: <span class="badge object-type-interface"></span> V2PanelSpec
---
# <span class="badge object-type-interface"></span> V2PanelSpec

## Definition

```typescript
export interface V2PanelSpec {
	id: number;
	title: string;
	// Shown in a info icon tooltip next to panel title
	description?: string;
	// Shown in a sub header below the title.
	subtitle?: string;
	links: dashboardv2beta1.DataLink[];
	data: dashboardv2beta1.V2QueryGroupKind;
	vizConfig: dashboardv2beta1.VizConfigKind;
	transparent?: boolean;
}

```
## See also

 * <span class="badge builder"></span> [V2PanelSpecBuilder](./builder-V2PanelSpecBuilder.md)
