---
title: <span class="badge object-type-struct"></span> Snapshot
---
# <span class="badge object-type-struct"></span> Snapshot

A dashboard snapshot shares an interactive dashboard publicly.

It is a read-only version of a dashboard, and is not editable.

It is possible to create a snapshot of a snapshot.

Grafana strips away all sensitive information from the dashboard.

Sensitive information stripped: queries (metric, template,annotation) and panel links.

## Definition

```go
type Snapshot struct {
    // Time when the snapshot was created
    Created time.Time `json:"created"`
    // Time when the snapshot expires, default is never to expire
    Expires string `json:"expires"`
    // Is the snapshot saved in an external grafana instance
    External bool `json:"external"`
    // external url, if snapshot was shared in external grafana instance
    ExternalUrl string `json:"externalUrl"`
    // Unique identifier of the snapshot
    Id uint32 `json:"id"`
    // Optional, defined the unique key of the snapshot, required if external is true
    Key string `json:"key"`
    // Optional, name of the snapshot
    Name string `json:"name"`
    // org id of the snapshot
    OrgId uint32 `json:"orgId"`
    // last time when the snapshot was updated
    Updated time.Time `json:"updated"`
    // url of the snapshot, if snapshot was shared internally
    Url *string `json:"url,omitempty"`
    // user id of the snapshot creator
    UserId uint32 `json:"userId"`
    Dashboard *dashboard.Dashboard `json:"dashboard,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Snapshot` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (snapshot *Snapshot) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Snapshot` objects.

```go
func (snapshot *Snapshot) Equals(other Snapshot) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Snapshot` fields for violations and returns them.

```go
func (snapshot *Snapshot) Validate() error
```

## See also

 * <span class="badge builder"></span> [SnapshotBuilder](./builder-SnapshotBuilder.md)
