---
title: <span class="badge builder"></span> ContactPointBuilder
---
# <span class="badge builder"></span> ContactPointBuilder

## Constructor

```go
func NewContactPointBuilder() *ContactPointBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ContactPointBuilder) Build() (ContactPoint, error)
```

### <span class="badge object-method"></span> DisableResolveMessage

```go
func (builder *ContactPointBuilder) DisableResolveMessage(disableResolveMessage bool) *ContactPointBuilder
```

### <span class="badge object-method"></span> Name

Name is used as grouping key in the UI. Contact points with the

same name will be grouped in the UI.

```go
func (builder *ContactPointBuilder) Name(name string) *ContactPointBuilder
```

### <span class="badge object-method"></span> Provenance

```go
func (builder *ContactPointBuilder) Provenance(provenance string) *ContactPointBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *ContactPointBuilder) Settings(settings alerting.Json) *ContactPointBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *ContactPointBuilder) Type(typeArg alerting.ContactPointType) *ContactPointBuilder
```

### <span class="badge object-method"></span> Uid

UID is the unique identifier of the contact point. The UID can be

set by the user.

```go
func (builder *ContactPointBuilder) Uid(uid string) *ContactPointBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ContactPoint](./object-ContactPoint.md)
