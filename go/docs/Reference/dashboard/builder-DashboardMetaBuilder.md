---
title: <span class="badge builder"></span> DashboardMetaBuilder
---
# <span class="badge builder"></span> DashboardMetaBuilder

## Constructor

```go
func NewDashboardMetaBuilder() *DashboardMetaBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DashboardMetaBuilder) Build() (DashboardMeta, error)
```

### <span class="badge object-method"></span> AnnotationsPermissions

```go
func (builder *DashboardMetaBuilder) AnnotationsPermissions(annotationsPermissions cog.Builder[dashboard.AnnotationPermission]) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> CanAdmin

```go
func (builder *DashboardMetaBuilder) CanAdmin(canAdmin bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> CanDelete

```go
func (builder *DashboardMetaBuilder) CanDelete(canDelete bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> CanEdit

```go
func (builder *DashboardMetaBuilder) CanEdit(canEdit bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> CanSave

```go
func (builder *DashboardMetaBuilder) CanSave(canSave bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> CanStar

```go
func (builder *DashboardMetaBuilder) CanStar(canStar bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> Created

```go
func (builder *DashboardMetaBuilder) Created(created time.Time) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> CreatedBy

```go
func (builder *DashboardMetaBuilder) CreatedBy(createdBy string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> Expires

```go
func (builder *DashboardMetaBuilder) Expires(expires time.Time) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> FolderId

```go
func (builder *DashboardMetaBuilder) FolderId(folderId int64) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> FolderTitle

```go
func (builder *DashboardMetaBuilder) FolderTitle(folderTitle string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> FolderUid

```go
func (builder *DashboardMetaBuilder) FolderUid(folderUid string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> FolderUrl

```go
func (builder *DashboardMetaBuilder) FolderUrl(folderUrl string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> HasAcl

```go
func (builder *DashboardMetaBuilder) HasAcl(hasAcl bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> IsFolder

```go
func (builder *DashboardMetaBuilder) IsFolder(isFolder bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> IsSnapshot

```go
func (builder *DashboardMetaBuilder) IsSnapshot(isSnapshot bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> IsStarred

```go
func (builder *DashboardMetaBuilder) IsStarred(isStarred bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> Provisioned

```go
func (builder *DashboardMetaBuilder) Provisioned(provisioned bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> ProvisionedExternalId

```go
func (builder *DashboardMetaBuilder) ProvisionedExternalId(provisionedExternalId string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> PublicDashboardAccessToken

```go
func (builder *DashboardMetaBuilder) PublicDashboardAccessToken(publicDashboardAccessToken string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> PublicDashboardEnabled

```go
func (builder *DashboardMetaBuilder) PublicDashboardEnabled(publicDashboardEnabled bool) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> PublicDashboardUid

```go
func (builder *DashboardMetaBuilder) PublicDashboardUid(publicDashboardUid string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> Slug

```go
func (builder *DashboardMetaBuilder) Slug(slug string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *DashboardMetaBuilder) Type(typeArg string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> Updated

```go
func (builder *DashboardMetaBuilder) Updated(updated time.Time) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> UpdatedBy

```go
func (builder *DashboardMetaBuilder) UpdatedBy(updatedBy string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> Url

```go
func (builder *DashboardMetaBuilder) Url(url string) *DashboardMetaBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *DashboardMetaBuilder) Version(version int64) *DashboardMetaBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DashboardMeta](./object-DashboardMeta.md)
