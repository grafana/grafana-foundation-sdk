---
title: <span class="badge builder"></span> NodeOptionsBuilder
---
# <span class="badge builder"></span> NodeOptionsBuilder

## Constructor

```java
new NodeOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public NodeOptions build()
```

### <span class="badge object-method"></span> arcs

Define which fields are shown as part of the node arc (colored circle around the node).

```java
public NodeOptionsBuilder arcs(List<com.grafana.foundation.cog.Builder<ArcOption>> arcs)
```

### <span class="badge object-method"></span> mainStatUnit

Unit for the main stat to override what ever is set in the data frame.

```java
public NodeOptionsBuilder mainStatUnit(String mainStatUnit)
```

### <span class="badge object-method"></span> secondaryStatUnit

Unit for the secondary stat to override what ever is set in the data frame.

```java
public NodeOptionsBuilder secondaryStatUnit(String secondaryStatUnit)
```

## See also

 * <span class="badge object-type-class"></span> [NodeOptions](./object-NodeOptions.md)
