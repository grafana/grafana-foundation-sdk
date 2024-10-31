---
title: <span class="badge builder"></span> KeyBuilder
---
# <span class="badge builder"></span> KeyBuilder

## Constructor

```go
func NewKeyBuilder() *KeyBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *KeyBuilder) Build() (Key, error)
```

### <span class="badge object-method"></span> Tick

```go
func (builder *KeyBuilder) Tick(tick float64) *KeyBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *KeyBuilder) Type(typeArg string) *KeyBuilder
```

### <span class="badge object-method"></span> Uid

```go
func (builder *KeyBuilder) Uid(uid string) *KeyBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Key](./object-Key.md)
