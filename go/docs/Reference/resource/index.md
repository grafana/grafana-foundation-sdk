# resource

## Objects

 * <span class="badge object-type-struct"></span> [Manifest](./object-Manifest.md)
 * <span class="badge object-type-struct"></span> [Metadata](./object-Metadata.md)
## Builders

 * <span class="badge builder"></span> [ManifestBuilder](./builder-ManifestBuilder.md)
 * <span class="badge builder"></span> [MetadataBuilder](./builder-MetadataBuilder.md)
## Functions

### <span class="badge function"></span> NewManifest

NewManifest creates a new Manifest object.

```go
func NewManifest() *Manifest
```

### <span class="badge function"></span> NewMetadata

NewMetadata creates a new Metadata object.

```go
func NewMetadata() *Metadata
```

### <span class="badge function"></span> ManifestConverter

ManifestConverter accepts a `Manifest` object and generates the Go code to build this object using builders.

```go
func ManifestConverter(input Manifest) string
```

### <span class="badge function"></span> MetadataConverter

MetadataConverter accepts a `Metadata` object and generates the Go code to build this object using builders.

```go
func MetadataConverter(input Metadata) string
```

