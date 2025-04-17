---
title: <span class="badge builder"></span> ManifestBuilder
---
# <span class="badge builder"></span> ManifestBuilder

## Constructor

```go
func NewManifestBuilder() *ManifestBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ManifestBuilder) Build() (Manifest, error)
```

### <span class="badge object-method"></span> ApiVersion

```go
func (builder *ManifestBuilder) ApiVersion(apiVersion string) *ManifestBuilder
```

### <span class="badge object-method"></span> Kind

```go
func (builder *ManifestBuilder) Kind(kind string) *ManifestBuilder
```

### <span class="badge object-method"></span> Metadata

```go
func (builder *ManifestBuilder) Metadata(metadata cog.Builder[resource.Metadata]) *ManifestBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *ManifestBuilder) Spec(spec any) *ManifestBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Manifest](./object-Manifest.md)
