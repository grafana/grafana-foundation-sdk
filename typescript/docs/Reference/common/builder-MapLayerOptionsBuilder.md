---
title: <span class="badge builder"></span> MapLayerOptionsBuilder
---
# <span class="badge builder"></span> MapLayerOptionsBuilder

## Constructor

```typescript
new MapLayerOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> config

Custom options depending on the type

```typescript
config(config: any)
```

### <span class="badge object-method"></span> filterData

Defines a frame MatcherConfig that may filter data for the given layer

```typescript
filterData(filterData: any)
```

### <span class="badge object-method"></span> location

Common method to define geometry fields

```typescript
location(location: cog.Builder<common.FrameGeometrySource>)
```

### <span class="badge object-method"></span> name

configured unique display name

```typescript
name(name: string)
```

### <span class="badge object-method"></span> opacity

Common properties:

https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html

Layer opacity (0-1)

```typescript
opacity(opacity: number)
```

### <span class="badge object-method"></span> tooltip

Check tooltip (defaults to true)

```typescript
tooltip(tooltip: boolean)
```

### <span class="badge object-method"></span> type

```typescript
type(type: string)
```

## See also

 * <span class="badge object-type-interface"></span> [MapLayerOptions](./object-MapLayerOptions.md)
