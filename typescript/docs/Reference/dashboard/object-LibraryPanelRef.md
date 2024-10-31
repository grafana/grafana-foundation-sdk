---
title: <span class="badge object-type-interface"></span> LibraryPanelRef
---
# <span class="badge object-type-interface"></span> LibraryPanelRef

A library panel is a reusable panel that you can use in any dashboard.

When you make a change to a library panel, that change propagates to all instances of where the panel is used.

Library panels streamline reuse of panels across multiple dashboards.

## Definition

```typescript
export interface LibraryPanelRef {
	// Library panel name
	name: string;
	// Library panel uid
	uid: string;
}

```
## Methods

No methods.
