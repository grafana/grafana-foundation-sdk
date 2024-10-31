---
title: <span class="badge object-type-struct"></span> NotificationPolicy
---
# <span class="badge object-type-struct"></span> NotificationPolicy

A Route is a node that contains definitions of how to handle alerts. This is modified

from the upstream alertmanager in that it adds the ObjectMatchers property.

## Definition

```go
type NotificationPolicy struct {
    Continue *bool `json:"continue,omitempty"`
    GroupBy []string `json:"group_by,omitempty"`
    GroupInterval *string `json:"group_interval,omitempty"`
    GroupWait *string `json:"group_wait,omitempty"`
    // Deprecated. Remove before v1.0 release.
    Match map[string]string `json:"match,omitempty"`
    MatchRe *alerting.MatchRegexps `json:"match_re,omitempty"`
    // Matchers is a slice of Matchers that is sortable, implements Stringer, and
    // provides a Matches method to match a LabelSet against all Matchers in the
    // slice. Note that some users of Matchers might require it to be sorted.
    Matchers *alerting.Matchers `json:"matchers,omitempty"`
    MuteTimeIntervals []string `json:"mute_time_intervals,omitempty"`
    ObjectMatchers *alerting.ObjectMatchers `json:"object_matchers,omitempty"`
    Provenance *alerting.Provenance `json:"provenance,omitempty"`
    Receiver *string `json:"receiver,omitempty"`
    RepeatInterval *string `json:"repeat_interval,omitempty"`
    Routes []alerting.NotificationPolicy `json:"routes,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotificationPolicy` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (notificationPolicy *NotificationPolicy) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NotificationPolicy` objects.

```go
func (notificationPolicy *NotificationPolicy) Equals(other NotificationPolicy) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NotificationPolicy` fields for violations and returns them.

```go
func (notificationPolicy *NotificationPolicy) Validate() error
```

## See also

 * <span class="badge builder"></span> [NotificationPolicyBuilder](./builder-NotificationPolicyBuilder.md)
