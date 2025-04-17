---
title: <span class="badge object-type-struct"></span> Metadata
---
# <span class="badge object-type-struct"></span> Metadata

## Definition

```go
type Metadata struct {
    Name string `json:"name"`
    Namespace *string `json:"namespace,omitempty"`
    Labels map[string]string `json:"labels,omitempty"`
    Annotations map[string]string `json:"annotations,omitempty"`
    Uid *string `json:"uid,omitempty"`
    ResourceVersion *string `json:"resourceVersion,omitempty"`
    Generation *int64 `json:"generation,omitempty"`
    CreationTimestamp *string `json:"creationTimestamp,omitempty"`
    UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
    DeletionTimestamp *string `json:"deletionTimestamp,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Metadata` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metadata *Metadata) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Metadata` objects.

```go
func (metadata *Metadata) Equals(other Metadata) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Metadata` fields for violations and returns them.

```go
func (metadata *Metadata) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetadataBuilder](./builder-MetadataBuilder.md)
