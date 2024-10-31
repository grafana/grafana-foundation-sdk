---
title: <span class="badge object-type-struct"></span> NotificationTemplate
---
# <span class="badge object-type-struct"></span> NotificationTemplate

## Definition

```go
type NotificationTemplate struct {
    Name *string `json:"name,omitempty"`
    Provenance *alerting.Provenance `json:"provenance,omitempty"`
    Template *string `json:"template,omitempty"`
    Version *string `json:"version,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotificationTemplate` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (notificationTemplate *NotificationTemplate) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NotificationTemplate` objects.

```go
func (notificationTemplate *NotificationTemplate) Equals(other NotificationTemplate) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NotificationTemplate` fields for violations and returns them.

```go
func (notificationTemplate *NotificationTemplate) Validate() error
```

## See also

 * <span class="badge builder"></span> [NotificationTemplateBuilder](./builder-NotificationTemplateBuilder.md)
