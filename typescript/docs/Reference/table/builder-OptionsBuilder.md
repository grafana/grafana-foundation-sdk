---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```typescript
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> cellHeight

Controls the height of the rows

```typescript
cellHeight(cellHeight: common.TableCellHeight)
```

### <span class="badge object-method"></span> footer

Controls footer options

```typescript
footer(footer: cog.Builder<common.TableFooterOptions>)
```

### <span class="badge object-method"></span> frameIndex

Represents the index of the selected frame

```typescript
frameIndex(frameIndex: number)
```

### <span class="badge object-method"></span> showHeader

Controls whether the panel should show the header

```typescript
showHeader(showHeader: boolean)
```

### <span class="badge object-method"></span> showTypeIcons

Controls whether the header should show icons for the column types

```typescript
showTypeIcons(showTypeIcons: boolean)
```

### <span class="badge object-method"></span> sortBy

Used to control row sorting

```typescript
sortBy(sortBy: cog.Builder<common.TableSortByFieldState>[])
```

## See also

 * <span class="badge object-type-interface"></span> [Options](./object-Options.md)
