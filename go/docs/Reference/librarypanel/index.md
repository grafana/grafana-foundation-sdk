# <span class="badge package-core"></span> librarypanel

## Objects

 * <span class="badge object-type-struct"></span> [LibraryElementDTOMeta](./object-LibraryElementDTOMeta.md)
 * <span class="badge object-type-struct"></span> [LibraryElementDTOMetaUser](./object-LibraryElementDTOMetaUser.md)
 * <span class="badge object-type-struct"></span> [LibraryPanel](./object-LibraryPanel.md)
 * <span class="badge object-type-struct"></span> [PanelModel](./object-PanelModel.md)
 * <span class="badge object-type-enum"></span> [PanelModelRepeatDirection](./object-PanelModelRepeatDirection.md)
## Builders

 * <span class="badge builder"></span> [LibraryElementDTOMetaBuilder](./builder-LibraryElementDTOMetaBuilder.md)
 * <span class="badge builder"></span> [LibraryElementDTOMetaUserBuilder](./builder-LibraryElementDTOMetaUserBuilder.md)
 * <span class="badge builder"></span> [LibraryPanelBuilder](./builder-LibraryPanelBuilder.md)
 * <span class="badge builder"></span> [PanelModelBuilder](./builder-PanelModelBuilder.md)
## Functions

### <span class="badge function"></span> NewLibraryPanel

NewLibraryPanel creates a new LibraryPanel object.

```go
func NewLibraryPanel() *LibraryPanel
```

### <span class="badge function"></span> NewLibraryElementDTOMeta

NewLibraryElementDTOMeta creates a new LibraryElementDTOMeta object.

```go
func NewLibraryElementDTOMeta() *LibraryElementDTOMeta
```

### <span class="badge function"></span> NewLibraryElementDTOMetaUser

NewLibraryElementDTOMetaUser creates a new LibraryElementDTOMetaUser object.

```go
func NewLibraryElementDTOMetaUser() *LibraryElementDTOMetaUser
```

### <span class="badge function"></span> NewPanelModel

NewPanelModel creates a new PanelModel object.

```go
func NewPanelModel() *PanelModel
```

### <span class="badge function"></span> LibraryPanelConverter

LibraryPanelConverter accepts a `LibraryPanel` object and generates the Go code to build this object using builders.

```go
func LibraryPanelConverter(input LibraryPanel) string
```

### <span class="badge function"></span> LibraryElementDTOMetaConverter

LibraryElementDTOMetaConverter accepts a `LibraryElementDTOMeta` object and generates the Go code to build this object using builders.

```go
func LibraryElementDTOMetaConverter(input LibraryElementDTOMeta) string
```

### <span class="badge function"></span> LibraryElementDTOMetaUserConverter

LibraryElementDTOMetaUserConverter accepts a `LibraryElementDTOMetaUser` object and generates the Go code to build this object using builders.

```go
func LibraryElementDTOMetaUserConverter(input LibraryElementDTOMetaUser) string
```

### <span class="badge function"></span> PanelModelConverter

PanelModelConverter accepts a `PanelModel` object and generates the Go code to build this object using builders.

```go
func PanelModelConverter(input PanelModel) string
```

