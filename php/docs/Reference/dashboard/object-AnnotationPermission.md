---
title: <span class="badge object-type-class"></span> AnnotationPermission
---
# <span class="badge object-type-class"></span> AnnotationPermission

+k8s:deepcopy-gen=true

## Definition

```php
class AnnotationPermission implements \JsonSerializable
{
    /**
     * +k8s:deepcopy-gen=true
     */
    public ?\Grafana\Foundation\Dashboard\AnnotationActions $dashboard;

    /**
     * +k8s:deepcopy-gen=true
     */
    public ?\Grafana\Foundation\Dashboard\AnnotationActions $organization;

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

 * <span class="badge builder"></span> [AnnotationPermissionBuilder](./builder-AnnotationPermissionBuilder.md)
