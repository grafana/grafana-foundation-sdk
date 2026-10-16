---
title: <span class="badge object-type-interface"></span> PanelSpec
---
# <span class="badge object-type-interface"></span> PanelSpec

## Definition

```typescript
export interface PanelSpec {
	id: number;
	title: string;
	// Shown in a info icon tooltip next to panel title
	description?: string;
	// Shown in a sub header below the title.
	subtitle?: string;
	links: dashboardv2beta1.DataLink[];
	data: dashboardv2beta1.QueryGroupKind;
	vizConfig: dashboardv2beta1.VizConfigKind;
	transparent?: boolean;
}

```
