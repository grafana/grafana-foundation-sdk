---
title: <span class="badge object-type-struct"></span> SubscriptionsQuery
---
# <span class="badge object-type-struct"></span> SubscriptionsQuery

## Definition

```go
type SubscriptionsQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
    Kind string `json:"kind"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SubscriptionsQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (subscriptionsQuery *SubscriptionsQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SubscriptionsQuery` objects.

```go
func (subscriptionsQuery *SubscriptionsQuery) Equals(other SubscriptionsQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SubscriptionsQuery` fields for violations and returns them.

```go
func (subscriptionsQuery *SubscriptionsQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [SubscriptionsQueryBuilder](./builder-SubscriptionsQueryBuilder.md)
