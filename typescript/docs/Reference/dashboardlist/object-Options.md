---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	keepTime: boolean;
	includeVars: boolean;
	showStarred: boolean;
	showRecentlyViewed: boolean;
	showSearch: boolean;
	showHeadings: boolean;
	showFolderNames: boolean;
	maxItems: number;
	query: string;
	tags: string[];
	// folderId is deprecated, and migrated to folderUid on panel init
	folderId?: number;
	folderUID?: string;
}

```
## Methods

No methods.
