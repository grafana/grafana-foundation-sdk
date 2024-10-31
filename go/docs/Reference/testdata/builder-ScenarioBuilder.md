---
title: <span class="badge builder"></span> ScenarioBuilder
---
# <span class="badge builder"></span> ScenarioBuilder

## Constructor

```go
func NewScenarioBuilder() *ScenarioBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ScenarioBuilder) Build() (Scenario, error)
```

### <span class="badge object-method"></span> Description

```go
func (builder *ScenarioBuilder) Description(description string) *ScenarioBuilder
```

### <span class="badge object-method"></span> HideAliasField

```go
func (builder *ScenarioBuilder) HideAliasField(hideAliasField bool) *ScenarioBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *ScenarioBuilder) Id(id string) *ScenarioBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *ScenarioBuilder) Name(name string) *ScenarioBuilder
```

### <span class="badge object-method"></span> StringInput

```go
func (builder *ScenarioBuilder) StringInput(stringInput string) *ScenarioBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Scenario](./object-Scenario.md)
