---
title: <span class="badge builder"></span> V2PanelKindBuilder
---
# <span class="badge builder"></span> V2PanelKindBuilder

The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,

which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has

to be forked rather than shared because PanelKind reaches TransformationKind through

QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not

change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.

## Constructor

```java
new V2PanelKindBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public V2PanelKind build()
```

### <span class="badge object-method"></span> spec

```java
public V2PanelKindBuilder spec(com.grafana.foundation.cog.Builder<V2PanelSpec> spec)
```

## See also

 * <span class="badge object-type-class"></span> [V2PanelKind](./object-V2PanelKind.md)
