---
title: <span class="badge object-type-class"></span> DashboardMeta
---
# <span class="badge object-type-class"></span> DashboardMeta

## Definition

```python
class DashboardMeta:
    annotations_permissions: typing.Optional[dashboard.AnnotationPermission]
    can_admin: typing.Optional[bool]
    can_delete: typing.Optional[bool]
    can_edit: typing.Optional[bool]
    can_save: typing.Optional[bool]
    can_star: typing.Optional[bool]
    created: typing.Optional[str]
    created_by: typing.Optional[str]
    expires: typing.Optional[str]
    folder_id: typing.Optional[int]
    folder_title: typing.Optional[str]
    folder_uid: typing.Optional[str]
    folder_url: typing.Optional[str]
    has_acl: typing.Optional[bool]
    is_folder: typing.Optional[bool]
    is_snapshot: typing.Optional[bool]
    is_starred: typing.Optional[bool]
    provisioned: typing.Optional[bool]
    provisioned_external_id: typing.Optional[str]
    public_dashboard_access_token: typing.Optional[str]
    public_dashboard_enabled: typing.Optional[bool]
    public_dashboard_uid: typing.Optional[str]
    slug: typing.Optional[str]
    type_val: typing.Optional[str]
    updated: typing.Optional[str]
    updated_by: typing.Optional[str]
    url: typing.Optional[str]
    version: typing.Optional[int]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [DashboardMeta](./builder-DashboardMeta.md)
