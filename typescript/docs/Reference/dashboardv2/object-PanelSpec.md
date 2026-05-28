---
title: <span class="badge object-type-interface"></span> PanelSpec
---
# <span class="badge object-type-interface"></span> PanelSpec

## Definition

```typescript
export interface PanelSpec {
	id: number;
	title: string;
	description: string;
	links: dashboardv2.DataLink[];
	data: dashboardv2.QueryGroupKind;
	vizConfig: dashboardv2.VizConfigKind;
	transparent?: boolean;
}

```
