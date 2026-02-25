---
title: <span class="badge object-type-interface"></span> LibraryPanelKindSpec
---
# <span class="badge object-type-interface"></span> LibraryPanelKindSpec

## Definition

```typescript
export interface LibraryPanelKindSpec {
	// Panel ID for the library panel in the dashboard
	id: number;
	// Title for the library panel in the dashboard
	title: string;
	libraryPanel: dashboardv2beta1.LibraryPanelRef;
}

```
