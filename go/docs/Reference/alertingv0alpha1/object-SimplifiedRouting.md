---
title: <span class="badge object-type-struct"></span> SimplifiedRouting
---
# <span class="badge object-type-struct"></span> SimplifiedRouting

## Definition

```go
type SimplifiedRouting struct {
    Type alertingv0alpha1.NotificationSettingsType `json:"type"`
    Receiver string `json:"receiver"`
    GroupBy []string `json:"groupBy,omitempty"`
    GroupWait *alertingv0alpha1.PromDuration `json:"groupWait,omitempty"`
    GroupInterval *alertingv0alpha1.PromDuration `json:"groupInterval,omitempty"`
    RepeatInterval *alertingv0alpha1.PromDuration `json:"repeatInterval,omitempty"`
    MuteTimeIntervals []alertingv0alpha1.TimeIntervalRef `json:"muteTimeIntervals,omitempty"`
    ActiveTimeIntervals []alertingv0alpha1.TimeIntervalRef `json:"activeTimeIntervals,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SimplifiedRouting` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (simplifiedRouting *SimplifiedRouting) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SimplifiedRouting` objects.

```go
func (simplifiedRouting *SimplifiedRouting) Equals(other SimplifiedRouting) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SimplifiedRouting` fields for violations and returns them.

```go
func (simplifiedRouting *SimplifiedRouting) Validate() error
```

## See also

 * <span class="badge builder"></span> [SimplifiedRoutingBuilder](./builder-SimplifiedRoutingBuilder.md)
