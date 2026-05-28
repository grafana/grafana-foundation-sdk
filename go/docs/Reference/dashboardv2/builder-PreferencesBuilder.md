---
title: <span class="badge builder"></span> PreferencesBuilder
---
# <span class="badge builder"></span> PreferencesBuilder

## Constructor

```go
func NewPreferencesBuilder() *PreferencesBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PreferencesBuilder) Build() (Preferences, error)
```

### <span class="badge object-method"></span> Layout

default layout template to be used when new containers are created

```go
func (builder *PreferencesBuilder) Layout(layout dashboardv2.AutoGridLayoutKindOrGridLayoutKind) *PreferencesBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Preferences](./object-Preferences.md)
