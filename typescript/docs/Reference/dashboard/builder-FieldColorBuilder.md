---
title: <span class="badge builder"></span> FieldColorBuilder
---
# <span class="badge builder"></span> FieldColorBuilder

## Constructor

```typescript
new FieldColorBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> fixedColor

The fixed color value for fixed or shades color modes.

```typescript
fixedColor(fixedColor: string)
```

### <span class="badge object-method"></span> mode

The main color scheme mode.

```typescript
mode(mode: dashboard.FieldColorModeId)
```

### <span class="badge object-method"></span> seriesBy

Some visualizations need to know how to assign a series color from by value color schemes.

```typescript
seriesBy(seriesBy: dashboard.FieldColorSeriesByMode)
```

## See also

 * <span class="badge object-type-interface"></span> [FieldColor](./object-FieldColor.md)
