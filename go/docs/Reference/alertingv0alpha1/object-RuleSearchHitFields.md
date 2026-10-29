---
title: <span class="badge object-type-struct"></span> RuleSearchHitFields
---
# <span class="badge object-type-struct"></span> RuleSearchHitFields

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Definition

```go
type RuleSearchHitFields struct {
    Title *string `json:"title,omitempty"`
    Folder *string `json:"folder,omitempty"`
    Type *string `json:"type,omitempty"`
    Interval *string `json:"interval,omitempty"`
    Paused *bool `json:"paused,omitempty"`
    Labels map[string]string `json:"labels,omitempty"`
    DatasourceUIDs []string `json:"datasourceUIDs,omitempty"`
    // Alert-rule fields.
    Annotations map[string]string `json:"annotations,omitempty"`
    For *string `json:"for,omitempty"`
    KeepFiringFor *string `json:"keepFiringFor,omitempty"`
    DashboardUID *string `json:"dashboardUID,omitempty"`
    PanelID *int64 `json:"panelID,omitempty"`
    Receiver *string `json:"receiver,omitempty"`
    NotificationType *string `json:"notificationType,omitempty"`
    RoutingTree *string `json:"routingTree,omitempty"`
    // Recording-rule fields.
    Metric *string `json:"metric,omitempty"`
    TargetDatasourceUID *string `json:"targetDatasourceUID,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RuleSearchHitFields` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (ruleSearchHitFields *RuleSearchHitFields) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RuleSearchHitFields` objects.

```go
func (ruleSearchHitFields *RuleSearchHitFields) Equals(other RuleSearchHitFields) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RuleSearchHitFields` fields for violations and returns them.

```go
func (ruleSearchHitFields *RuleSearchHitFields) Validate() error
```

## See also

 * <span class="badge builder"></span> [RuleSearchHitFieldsBuilder](./builder-RuleSearchHitFieldsBuilder.md)
