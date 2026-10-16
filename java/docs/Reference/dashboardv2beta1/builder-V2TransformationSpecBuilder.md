---
title: <span class="badge builder"></span> V2TransformationSpecBuilder
---
# <span class="badge builder"></span> V2TransformationSpecBuilder

Dashboard v2 shape: no `id`, it is carried by the parent's `group`.

## Constructor

```java
new V2TransformationSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public V2TransformationSpec build()
```

### <span class="badge object-method"></span> disabled

Disabled transformations are skipped

```java
public V2TransformationSpecBuilder disabled(Boolean disabled)
```

### <span class="badge object-method"></span> filter

Optional frame matcher. When missing it will be applied to all results

```java
public V2TransformationSpecBuilder filter(MatcherConfig filter)
```

### <span class="badge object-method"></span> options

Options to be passed to the transformer

Valid options depend on the transformer id

```java
public V2TransformationSpecBuilder options(Object options)
```

### <span class="badge object-method"></span> topic

Where to pull DataFrames from as input to transformation

```java
public V2TransformationSpecBuilder topic(DataTopic topic)
```

## See also

 * <span class="badge object-type-class"></span> [V2TransformationSpec](./object-V2TransformationSpec.md)
