---
title: <span class="badge object-type-struct"></span> Item
---
# <span class="badge object-type-struct"></span> Item

## Definition

```go
type Item struct {
    // type of the item.
    Type playlistv0alpha1.ItemType `json:"type"`
    // Value depends on type and describes the playlist item.
    //  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
    //  is not portable as the numerical identifier is non-deterministic between different instances.
    //  Will be replaced by dashboard_by_uid in the future. (deprecated)
    //  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
    //  dashboards behind the tag will be added to the playlist.
    //  - dashboard_by_uid: The value is the dashboard UID
    Value string `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Item` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (item *Item) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Item` objects.

```go
func (item *Item) Equals(other Item) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Item` fields for violations and returns them.

```go
func (item *Item) Validate() error
```

## See also

 * <span class="badge builder"></span> [ItemBuilder](./builder-ItemBuilder.md)
