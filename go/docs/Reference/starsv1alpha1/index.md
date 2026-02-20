# starsv1alpha1

## Objects

 * <span class="badge object-type-struct"></span> [Resource](./object-Resource.md)
 * <span class="badge object-type-struct"></span> [Stars](./object-Stars.md)
 * <span class="badge object-type-scalar"></span> [StarsKind](./object-StarsKind.md)
## Builders

 * <span class="badge builder"></span> [ResourceBuilder](./builder-ResourceBuilder.md)
 * <span class="badge builder"></span> [StarsBuilder](./builder-StarsBuilder.md)
## Functions

### <span class="badge function"></span> NewResource

NewResource creates a new Resource object.

```go
func NewResource() *Resource
```

### <span class="badge function"></span> NewStars

NewStars creates a new Stars object.

```go
func NewStars() *Stars
```

### <span class="badge function"></span> Manifest

Creates a resource manifest from Stars.

```go
func Manifest(name string, stars cog.Builder[Stars]) *StarsBuilder
```

### <span class="badge function"></span> ResourceConverter

ResourceConverter accepts a `Resource` object and generates the Go code to build this object using builders.

```go
func ResourceConverter(input Resource) string
```

### <span class="badge function"></span> StarsConverter

StarsConverter accepts a `Stars` object and generates the Go code to build this object using builders.

```go
func StarsConverter(input Stars) string
```

