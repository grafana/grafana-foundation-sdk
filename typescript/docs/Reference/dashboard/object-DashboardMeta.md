---
title: <span class="badge object-type-interface"></span> DashboardMeta
---
# <span class="badge object-type-interface"></span> DashboardMeta

## Definition

```typescript
export interface DashboardMeta {
	annotationsPermissions?: dashboard.AnnotationPermission;
	canAdmin?: boolean;
	canDelete?: boolean;
	canEdit?: boolean;
	canSave?: boolean;
	canStar?: boolean;
	created?: string;
	createdBy?: string;
	expires?: string;
	// Deprecated: use FolderUID instead
	folderId?: number;
	folderTitle?: string;
	folderUid?: string;
	folderUrl?: string;
	hasAcl?: boolean;
	isFolder?: boolean;
	isSnapshot?: boolean;
	isStarred?: boolean;
	provisioned?: boolean;
	provisionedExternalId?: string;
	publicDashboardEnabled?: boolean;
	slug?: string;
	type?: string;
	updated?: string;
	updatedBy?: string;
	url?: string;
	version?: number;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [DashboardMetaBuilder](./builder-DashboardMetaBuilder.md)