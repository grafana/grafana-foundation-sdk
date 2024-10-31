---
title: <span class="badge builder"></span> TeamBuilder
---
# <span class="badge builder"></span> TeamBuilder

## Constructor

```go
func NewTeamBuilder(name string) *TeamBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TeamBuilder) Build() (Team, error)
```

### <span class="badge object-method"></span> Email

Email of the team.

```go
func (builder *TeamBuilder) Email(email string) *TeamBuilder
```

### <span class="badge object-method"></span> Name

Name of the team.

```go
func (builder *TeamBuilder) Name(name string) *TeamBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Team](./object-Team.md)
