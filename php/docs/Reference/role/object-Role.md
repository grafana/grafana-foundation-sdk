---
title: <span class="badge object-type-class"></span> Role
---
# <span class="badge object-type-class"></span> Role

## Definition

```php
class Role implements \JsonSerializable
{
    /**
     * The role identifier `managed:builtins:editor:permissions`
     */
    public string $name;

    /**
     * Optional display
     */
    public ?string $displayName;

    /**
     * Name of the team.
     */
    public ?string $groupName;

    /**
     * Role description
     */
    public ?string $description;

    /**
     * Do not show this role
     */
    public bool $hidden;

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

 * <span class="badge builder"></span> [RoleBuilder](./builder-RoleBuilder.md)
