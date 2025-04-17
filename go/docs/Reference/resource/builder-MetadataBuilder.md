---
title: <span class="badge builder"></span> MetadataBuilder
---
# <span class="badge builder"></span> MetadataBuilder

## Constructor

```go
func NewMetadataBuilder() *MetadataBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetadataBuilder) Build() (Metadata, error)
```

### <span class="badge object-method"></span> Annotations

```go
func (builder *MetadataBuilder) Annotations(annotations map[string]string) *MetadataBuilder
```

### <span class="badge object-method"></span> CreationTimestamp

```go
func (builder *MetadataBuilder) CreationTimestamp(creationTimestamp string) *MetadataBuilder
```

### <span class="badge object-method"></span> DeletionTimestamp

```go
func (builder *MetadataBuilder) DeletionTimestamp(deletionTimestamp string) *MetadataBuilder
```

### <span class="badge object-method"></span> Generation

```go
func (builder *MetadataBuilder) Generation(generation int64) *MetadataBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *MetadataBuilder) Labels(labels map[string]string) *MetadataBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *MetadataBuilder) Name(name string) *MetadataBuilder
```

### <span class="badge object-method"></span> Namespace

```go
func (builder *MetadataBuilder) Namespace(namespace string) *MetadataBuilder
```

### <span class="badge object-method"></span> ResourceVersion

```go
func (builder *MetadataBuilder) ResourceVersion(resourceVersion string) *MetadataBuilder
```

### <span class="badge object-method"></span> Uid

```go
func (builder *MetadataBuilder) Uid(uid string) *MetadataBuilder
```

### <span class="badge object-method"></span> UpdateTimestamp

```go
func (builder *MetadataBuilder) UpdateTimestamp(updateTimestamp string) *MetadataBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Metadata](./object-Metadata.md)
