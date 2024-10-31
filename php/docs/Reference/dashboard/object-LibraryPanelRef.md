---
title: <span class="badge object-type-class"></span> LibraryPanelRef
---
# <span class="badge object-type-class"></span> LibraryPanelRef

A library panel is a reusable panel that you can use in any dashboard.

When you make a change to a library panel, that change propagates to all instances of where the panel is used.

Library panels streamline reuse of panels across multiple dashboards.

## Definition

```php
class LibraryPanelRef implements \JsonSerializable
{
    /**
     * Library panel name
     */
    public string $name;

    /**
     * Library panel uid
     */
    public string $uid;

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

