---
title: <span class="badge builder"></span> MapLayerOptionsBuilder
---
# <span class="badge builder"></span> MapLayerOptionsBuilder

## Constructor

```java
new MapLayerOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public MapLayerOptions build()
```

### <span class="badge object-method"></span> config

Custom options depending on the type

```java
public MapLayerOptionsBuilder config(Object config)
```

### <span class="badge object-method"></span> filterData

Defines a frame MatcherConfig that may filter data for the given layer

```java
public MapLayerOptionsBuilder filterData(Object filterData)
```

### <span class="badge object-method"></span> location

Common method to define geometry fields

```java
public MapLayerOptionsBuilder location(com.grafana.foundation.cog.Builder<FrameGeometrySource> location)
```

### <span class="badge object-method"></span> name

configured unique display name

```java
public MapLayerOptionsBuilder name(String name)
```

### <span class="badge object-method"></span> opacity

Common properties:

https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html

Layer opacity (0-1)

```java
public MapLayerOptionsBuilder opacity(Long opacity)
```

### <span class="badge object-method"></span> tooltip

Check tooltip (defaults to true)

```java
public MapLayerOptionsBuilder tooltip(Boolean tooltip)
```

### <span class="badge object-method"></span> type

```java
public MapLayerOptionsBuilder type(String type)
```

## See also

 * <span class="badge object-type-class"></span> [MapLayerOptions](./object-MapLayerOptions.md)
