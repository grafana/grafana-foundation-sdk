---
title: <span class="badge object-type-struct"></span> Playlist
---
# <span class="badge object-type-struct"></span> Playlist

## Definition

```go
type Playlist struct {
    Title string `json:"title"`
    Interval string `json:"interval"`
    Items []playlistv0alpha1.Item `json:"items"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Playlist` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (playlist *Playlist) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Playlist` objects.

```go
func (playlist *Playlist) Equals(other Playlist) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Playlist` fields for violations and returns them.

```go
func (playlist *Playlist) Validate() error
```

## See also

 * <span class="badge builder"></span> [PlaylistBuilder](./builder-PlaylistBuilder.md)
