---
title: <span class="badge object-type-class"></span> DashboardMeta
---
# <span class="badge object-type-class"></span> DashboardMeta

## Definition

```php
class DashboardMeta implements \JsonSerializable
{
    public ?\Grafana\Foundation\Dashboard\AnnotationPermission $annotationsPermissions;

    public ?bool $canAdmin;

    public ?bool $canDelete;

    public ?bool $canEdit;

    public ?bool $canSave;

    public ?bool $canStar;

    public ?string $created;

    public ?string $createdBy;

    public ?string $expires;

    public ?int $folderId;

    public ?string $folderTitle;

    public ?string $folderUid;

    public ?string $folderUrl;

    public ?bool $hasAcl;

    public ?bool $isFolder;

    public ?bool $isSnapshot;

    public ?bool $isStarred;

    public ?bool $provisioned;

    public ?string $provisionedExternalId;

    public ?string $publicDashboardAccessToken;

    public ?bool $publicDashboardEnabled;

    public ?string $publicDashboardUid;

    public ?string $slug;

    public ?string $type;

    public ?string $updated;

    public ?string $updatedBy;

    public ?string $url;

    public ?int $version;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [DashboardMetaBuilder](./builder-DashboardMetaBuilder.md)
