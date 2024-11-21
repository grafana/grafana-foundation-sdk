---
title: <span class="badge object-type-interface"></span> LibraryPanel
---
# <span class="badge object-type-interface"></span> LibraryPanel

## Definition

```typescript
export interface LibraryPanel {
	// Folder UID
	folderUid?: string;
	// Library element UID
	uid: string;
	// Panel name (also saved in the model)
	name: string;
	// Panel description
	description?: string;
	// The panel type (from inside the model)
	type: string;
	// Dashboard version when this was saved (zero if unknown)
	schemaVersion?: number;
	// panel version, incremented each time the dashboard is updated.
	version: number;
	// TODO: should be the same panel schema defined in dashboard
	// Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
	model: librarypanel.PanelModel;
	// Object storage metadata
	meta?: librarypanel.LibraryElementDTOMeta;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [LibraryPanelBuilder](./builder-LibraryPanelBuilder.md)
