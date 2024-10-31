---
title: <span class="badge builder"></span> SnapshotBuilder
---
# <span class="badge builder"></span> SnapshotBuilder

## Constructor

```go
func NewSnapshotBuilder() *SnapshotBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SnapshotBuilder) Build() (Snapshot, error)
```

### <span class="badge object-method"></span> Dashboard

```go
func (builder *SnapshotBuilder) Dashboard(dashboard cog.Builder[dashboard.Dashboard]) *SnapshotBuilder
```

### <span class="badge object-method"></span> Expires

Time when the snapshot expires, default is never to expire

```go
func (builder *SnapshotBuilder) Expires(expires string) *SnapshotBuilder
```

### <span class="badge object-method"></span> External

Is the snapshot saved in an external grafana instance

```go
func (builder *SnapshotBuilder) External(external bool) *SnapshotBuilder
```

### <span class="badge object-method"></span> ExternalUrl

external url, if snapshot was shared in external grafana instance

```go
func (builder *SnapshotBuilder) ExternalUrl(externalUrl string) *SnapshotBuilder
```

### <span class="badge object-method"></span> Id

Unique identifier of the snapshot

```go
func (builder *SnapshotBuilder) Id(id uint32) *SnapshotBuilder
```

### <span class="badge object-method"></span> Key

Optional, defined the unique key of the snapshot, required if external is true

```go
func (builder *SnapshotBuilder) Key(key string) *SnapshotBuilder
```

### <span class="badge object-method"></span> Name

Optional, name of the snapshot

```go
func (builder *SnapshotBuilder) Name(name string) *SnapshotBuilder
```

### <span class="badge object-method"></span> OrgId

org id of the snapshot

```go
func (builder *SnapshotBuilder) OrgId(orgId uint32) *SnapshotBuilder
```

### <span class="badge object-method"></span> OriginalUrl

original url, url of the dashboard that was snapshotted

```go
func (builder *SnapshotBuilder) OriginalUrl(originalUrl string) *SnapshotBuilder
```

### <span class="badge object-method"></span> Url

url of the snapshot, if snapshot was shared internally

```go
func (builder *SnapshotBuilder) Url(url string) *SnapshotBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Snapshot](./object-Snapshot.md)
