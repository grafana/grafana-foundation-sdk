---
title: <span class="badge builder"></span> LibraryPanelRef
---
# <span class="badge builder"></span> LibraryPanelRef

A library panel is a reusable panel that you can use in any dashboard.

When you make a change to a library panel, that change propagates to all instances of where the panel is used.

Library panels streamline reuse of panels across multiple dashboards.

## Constructor

```python
LibraryPanelRef()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.LibraryPanelRef
```

### <span class="badge object-method"></span> name

Library panel name

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> uid

Library panel uid

```python
def uid(uid: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [LibraryPanelRef](./object-LibraryPanelRef.md)
