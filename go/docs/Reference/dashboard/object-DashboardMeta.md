---
title: <span class="badge object-type-struct"></span> DashboardMeta
---
# <span class="badge object-type-struct"></span> DashboardMeta

## Definition

```go
type DashboardMeta struct {
    AnnotationsPermissions *dashboard.AnnotationPermission `json:"annotationsPermissions,omitempty"`
    CanAdmin *bool `json:"canAdmin,omitempty"`
    CanDelete *bool `json:"canDelete,omitempty"`
    CanEdit *bool `json:"canEdit,omitempty"`
    CanSave *bool `json:"canSave,omitempty"`
    CanStar *bool `json:"canStar,omitempty"`
    Created *time.Time `json:"created,omitempty"`
    CreatedBy *string `json:"createdBy,omitempty"`
    Expires *time.Time `json:"expires,omitempty"`
    FolderId *int64 `json:"folderId,omitempty"`
    FolderTitle *string `json:"folderTitle,omitempty"`
    FolderUid *string `json:"folderUid,omitempty"`
    FolderUrl *string `json:"folderUrl,omitempty"`
    HasAcl *bool `json:"hasAcl,omitempty"`
    IsFolder *bool `json:"isFolder,omitempty"`
    IsSnapshot *bool `json:"isSnapshot,omitempty"`
    IsStarred *bool `json:"isStarred,omitempty"`
    Provisioned *bool `json:"provisioned,omitempty"`
    ProvisionedExternalId *string `json:"provisionedExternalId,omitempty"`
    PublicDashboardAccessToken *string `json:"publicDashboardAccessToken,omitempty"`
    PublicDashboardEnabled *bool `json:"publicDashboardEnabled,omitempty"`
    PublicDashboardUid *string `json:"publicDashboardUid,omitempty"`
    Slug *string `json:"slug,omitempty"`
    Type *string `json:"type,omitempty"`
    Updated *time.Time `json:"updated,omitempty"`
    UpdatedBy *string `json:"updatedBy,omitempty"`
    Url *string `json:"url,omitempty"`
    Version *int64 `json:"version,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardMeta` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardMeta *DashboardMeta) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardMeta` objects.

```go
func (dashboardMeta *DashboardMeta) Equals(other DashboardMeta) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardMeta` fields for violations and returns them.

```go
func (dashboardMeta *DashboardMeta) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardMetaBuilder](./builder-DashboardMetaBuilder.md)
