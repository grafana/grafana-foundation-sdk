---
title: <span class="badge builder"></span> RuleBuilder
---
# <span class="badge builder"></span> RuleBuilder

## Constructor

```go
func NewRuleBuilder(title string) *RuleBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RuleBuilder) Build() (Rule, error)
```

### <span class="badge object-method"></span> Annotations

```go
func (builder *RuleBuilder) Annotations(annotations map[string]string) *RuleBuilder
```

### <span class="badge object-method"></span> Condition

```go
func (builder *RuleBuilder) Condition(condition string) *RuleBuilder
```

### <span class="badge object-method"></span> ExecErrState

```go
func (builder *RuleBuilder) ExecErrState(execErrState alerting.RuleExecErrState) *RuleBuilder
```

### <span class="badge object-method"></span> FolderUID

```go
func (builder *RuleBuilder) FolderUID(folderUID string) *RuleBuilder
```

### <span class="badge object-method"></span> For

The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.

Before this time has elapsed, the rule is only considered to be Pending.

```go
func (builder *RuleBuilder) For(forArg string) *RuleBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *RuleBuilder) Id(id int64) *RuleBuilder
```

### <span class="badge object-method"></span> IsPaused

```go
func (builder *RuleBuilder) IsPaused(isPaused bool) *RuleBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *RuleBuilder) Labels(labels map[string]string) *RuleBuilder
```

### <span class="badge object-method"></span> NoDataState

```go
func (builder *RuleBuilder) NoDataState(noDataState alerting.RuleNoDataState) *RuleBuilder
```

### <span class="badge object-method"></span> OrgID

```go
func (builder *RuleBuilder) OrgID(orgID int64) *RuleBuilder
```

### <span class="badge object-method"></span> Provenance

```go
func (builder *RuleBuilder) Provenance(provenance alerting.Provenance) *RuleBuilder
```

### <span class="badge object-method"></span> Queries

```go
func (builder *RuleBuilder) Queries(data []cog.Builder[alerting.Query]) *RuleBuilder
```

### <span class="badge object-method"></span> RuleGroup

```go
func (builder *RuleBuilder) RuleGroup(ruleGroup string) *RuleBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *RuleBuilder) Title(title string) *RuleBuilder
```

### <span class="badge object-method"></span> Uid

```go
func (builder *RuleBuilder) Uid(uid string) *RuleBuilder
```

### <span class="badge object-method"></span> Updated

```go
func (builder *RuleBuilder) Updated(updated time.Time) *RuleBuilder
```

### <span class="badge object-method"></span> WithQuery

```go
func (builder *RuleBuilder) WithQuery(data cog.Builder[alerting.Query]) *RuleBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Rule](./object-Rule.md)
