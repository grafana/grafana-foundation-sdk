---
title: <span class="badge builder"></span> ActionBuilder
---
# <span class="badge builder"></span> ActionBuilder

## Constructor

```go
func NewActionBuilder() *ActionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ActionBuilder) Build() (Action, error)
```

### <span class="badge object-method"></span> Confirmation

```go
func (builder *ActionBuilder) Confirmation(confirmation string) *ActionBuilder
```

### <span class="badge object-method"></span> Fetch

```go
func (builder *ActionBuilder) Fetch(fetch cog.Builder[dashboard.FetchOptions]) *ActionBuilder
```

### <span class="badge object-method"></span> Infinity

```go
func (builder *ActionBuilder) Infinity(infinity cog.Builder[dashboard.InfinityOptions]) *ActionBuilder
```

### <span class="badge object-method"></span> OneClick

```go
func (builder *ActionBuilder) OneClick(oneClick bool) *ActionBuilder
```

### <span class="badge object-method"></span> Style

```go
func (builder *ActionBuilder) Style(style cog.Builder[dashboard.DashboardActionStyle]) *ActionBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *ActionBuilder) Title(title string) *ActionBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *ActionBuilder) Type(typeArg dashboard.ActionType) *ActionBuilder
```

### <span class="badge object-method"></span> Variables

```go
func (builder *ActionBuilder) Variables(variables []cog.Builder[dashboard.ActionVariable]) *ActionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Action](./object-Action.md)
