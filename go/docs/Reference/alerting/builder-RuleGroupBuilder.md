---
title: <span class="badge builder"></span> RuleGroupBuilder
---
# <span class="badge builder"></span> RuleGroupBuilder

## Constructor

```go
func NewRuleGroupBuilder(title string) *RuleGroupBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RuleGroupBuilder) Build() (RuleGroup, error)
```

### <span class="badge object-method"></span> FolderUid

```go
func (builder *RuleGroupBuilder) FolderUid(folderUid string) *RuleGroupBuilder
```

### <span class="badge object-method"></span> Interval

The interval, in seconds, at which all rules in the group are evaluated.

If a group contains many rules, the rules are evaluated sequentially.

```go
func (builder *RuleGroupBuilder) Interval(interval alerting.Duration) *RuleGroupBuilder
```

### <span class="badge object-method"></span> Rules

```go
func (builder *RuleGroupBuilder) Rules(rules []cog.Builder[alerting.Rule]) *RuleGroupBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *RuleGroupBuilder) Title(title string) *RuleGroupBuilder
```

### <span class="badge object-method"></span> WithRule

```go
func (builder *RuleGroupBuilder) WithRule(rule cog.Builder[alerting.Rule]) *RuleGroupBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RuleGroup](./object-RuleGroup.md)
