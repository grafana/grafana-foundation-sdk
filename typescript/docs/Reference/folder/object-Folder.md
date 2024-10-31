---
title: <span class="badge object-type-interface"></span> Folder
---
# <span class="badge object-type-interface"></span> Folder

TODO:

common metadata will soon support setting the parent folder in the metadata

## Definition

```typescript
export interface Folder {
	// Unique folder id. (will be k8s name)
	uid: string;
	// Folder title
	title: string;
	// Description of the folder.
	description?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [FolderBuilder](./builder-FolderBuilder.md)
