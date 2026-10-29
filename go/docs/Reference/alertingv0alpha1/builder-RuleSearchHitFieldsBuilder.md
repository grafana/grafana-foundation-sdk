---
title: <span class="badge builder"></span> RuleSearchHitFieldsBuilder
---
# <span class="badge builder"></span> RuleSearchHitFieldsBuilder

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Constructor

```go
func NewRuleSearchHitFieldsBuilder() *RuleSearchHitFieldsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RuleSearchHitFieldsBuilder) Build() (RuleSearchHitFields, error)
```

### <span class="badge object-method"></span> Annotations

Alert-rule fields.

```go
func (builder *RuleSearchHitFieldsBuilder) Annotations(annotations map[string]string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> DashboardUID

```go
func (builder *RuleSearchHitFieldsBuilder) DashboardUID(dashboardUID string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> DatasourceUIDs

```go
func (builder *RuleSearchHitFieldsBuilder) DatasourceUIDs(datasourceUIDs []string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> Folder

```go
func (builder *RuleSearchHitFieldsBuilder) Folder(folder string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> For

```go
func (builder *RuleSearchHitFieldsBuilder) For(forArg string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> Interval

```go
func (builder *RuleSearchHitFieldsBuilder) Interval(interval string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> KeepFiringFor

```go
func (builder *RuleSearchHitFieldsBuilder) KeepFiringFor(keepFiringFor string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *RuleSearchHitFieldsBuilder) Labels(labels map[string]string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> Metric

Recording-rule fields.

```go
func (builder *RuleSearchHitFieldsBuilder) Metric(metric string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> NotificationType

```go
func (builder *RuleSearchHitFieldsBuilder) NotificationType(notificationType string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> PanelID

```go
func (builder *RuleSearchHitFieldsBuilder) PanelID(panelID int64) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> Paused

```go
func (builder *RuleSearchHitFieldsBuilder) Paused(paused bool) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> Receiver

```go
func (builder *RuleSearchHitFieldsBuilder) Receiver(receiver string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> RoutingTree

```go
func (builder *RuleSearchHitFieldsBuilder) RoutingTree(routingTree string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> TargetDatasourceUID

```go
func (builder *RuleSearchHitFieldsBuilder) TargetDatasourceUID(targetDatasourceUID string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *RuleSearchHitFieldsBuilder) Title(title string) *RuleSearchHitFieldsBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *RuleSearchHitFieldsBuilder) Type(typeArg string) *RuleSearchHitFieldsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RuleSearchHitFields](./object-RuleSearchHitFields.md)
