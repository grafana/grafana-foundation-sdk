---
title: <span class="badge builder"></span> RowBuilder
---
# <span class="badge builder"></span> RowBuilder

## Constructor

```go
func NewRowBuilder(title string) *RowBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RowBuilder) Build() (RowPanel, error)
```

### <span class="badge object-method"></span> Collapsed

Whether this row should be collapsed or not.

```go
func (builder *RowBuilder) Collapsed(collapsed bool) *RowBuilder
```

### <span class="badge object-method"></span> Datasource

Name of default datasource for the row

```go
func (builder *RowBuilder) Datasource(datasource dashboard.DataSourceRef) *RowBuilder
```

### <span class="badge object-method"></span> GridPos

Row grid position

```go
func (builder *RowBuilder) GridPos(gridPos dashboard.GridPos) *RowBuilder
```

### <span class="badge object-method"></span> Id

Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.

```go
func (builder *RowBuilder) Id(id uint32) *RowBuilder
```

### <span class="badge object-method"></span> Repeat

Name of template variable to repeat for.

```go
func (builder *RowBuilder) Repeat(repeat string) *RowBuilder
```

### <span class="badge object-method"></span> Title

Row title

```go
func (builder *RowBuilder) Title(title string) *RowBuilder
```

### <span class="badge object-method"></span> WithPanel

List of panels in the row

```go
func (builder *RowBuilder) WithPanel(panel cog.Builder[dashboard.Panel]) *RowBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RowPanel](./object-RowPanel.md)