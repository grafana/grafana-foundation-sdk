# folderv1beta1

## Objects

 * <span class="badge object-type-struct"></span> [Folder](./object-Folder.md)
 * <span class="badge object-type-scalar"></span> [FolderKind](./object-FolderKind.md)
 * <span class="badge object-type-scalar"></span> [FolderV1Beta1](./object-FolderV1Beta1.md)
## Builders

 * <span class="badge builder"></span> [FolderBuilder](./builder-FolderBuilder.md)
## Functions

### <span class="badge function"></span> NewFolder

NewFolder creates a new Folder object.

```go
func NewFolder() *Folder
```

### <span class="badge function"></span> Manifest

Creates a resource manifest from a Folder.

```go
func Manifest(name string, folder cog.Builder[Folder]) *FolderBuilder
```

### <span class="badge function"></span> FolderConverter

FolderConverter accepts a `Folder` object and generates the Go code to build this object using builders.

```go
func FolderConverter(input Folder) string
```

